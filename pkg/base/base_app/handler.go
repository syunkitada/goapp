package base_app

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"

	"github.com/gin-gonic/gin"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_protocol"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/base/base_spec_model"
	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (app *BaseApp) ExecQuery(w http.ResponseWriter, r *http.Request, isProxy bool) {
	var err error
	tctx := logger.NewTraceContext(app.host, app.name)
	startTime := logger.StartTrace(tctx)
	defer func() {
		if p := recover(); p != nil {
			w.WriteHeader(http.StatusInternalServerError)
			err = error_utils.NewRecoveredError(p)
			logger.Errorf(tctx, err, "Panic occured")
			fmt.Println("panic occured", err)
		}
	}()

	w.Header().Set("Access-Control-Allow-Origin", "http://192.168.10.121:3000") // TODO FIXME
	w.Header().Set("Access-Control-Allow-Credentials", "true")                  // TODO FIXME

	bufbody := new(bytes.Buffer)
	if _, err = bufbody.ReadFrom(r.Body); err != nil {
		return
	}
	rawReq := bufbody.Bytes()
	service, userAuthority, req, rep, err := app.Start(tctx, r, rawReq, isProxy)

	defer func() { app.End(tctx, startTime, err) }()
	if err != nil {
		var bytes []byte
		bytes, err = json.Marshal(&rep)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Error(tctx, err, "Failed json.Marshal")
			return
		}
		if _, err = w.Write(bytes); err != nil {
			return
		}
		return
	}

	statusCode := 0
	var repBytes []byte
	for _, endpoint := range service.Endpoints {
		if endpoint == "" {
			if err = app.queryHandler.Exec(tctx, userAuthority, r, w, req, rep); err != nil {
				break
			}
			repBytes, err = json.Marshal(&rep)
			break
		}

		if repBytes, statusCode, err = app.Proxy(tctx, service, endpoint, rawReq); err != nil {
			fmt.Println("DEBUG proxy failed", err, req.Queries)
			continue
		} else {
			fmt.Println("DEBUG proxy success", endpoint, req.Queries)
			break
		}
	}

	if statusCode != 0 {
		w.WriteHeader(statusCode)
	} else {
		if err == nil {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
	if _, err = w.Write(repBytes); err != nil {
		return
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (app *BaseApp) Ws(w http.ResponseWriter, r *http.Request, isProxy bool) {
	var err error
	var tmpErr error
	conMutex := sync.Mutex{}
	tctx := logger.NewTraceContext(app.host, app.name)
	startTime := logger.StartTrace(tctx)
	defer func() {
		if p := recover(); p != nil {
			w.WriteHeader(http.StatusInternalServerError)
			err = error_utils.NewRecoveredError(p)
			logger.Errorf(tctx, err, "Panic occured")
		}
	}()

	w.Header().Set("Access-Control-Allow-Origin", "http://192.168.10.121:3000") // TODO FIXME
	w.Header().Set("Access-Control-Allow-Credentials", "true")                  // TODO FIXME

	defer func() { app.End(tctx, startTime, err) }()

	var wsConn *websocket.Conn
	wsConn, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	var proxyWsConn *websocket.Conn
	wsDone := false
	defer func() {
		conMutex.Lock()
		if wsConn != nil {
			if tmpErr = wsConn.Close(); tmpErr != nil {
				logger.Warningf(tctx, "Failed  proxyWsConn.Close: err=%s", tmpErr.Error())
			} else {
				logger.Info(tctx, "Success proxyWsConn.Close")
			}
		}
		if proxyWsConn != nil {
			if tmpErr = proxyWsConn.Close(); tmpErr != nil {
				logger.Warningf(tctx, "Failed  proxyWsConn.Close: err=%s", tmpErr.Error())
			} else {
				logger.Info(tctx, "Success proxyWsConn.Close")
			}
		}
		wsDone = true
		conMutex.Unlock()
	}()

	var service *base_spec_model.ServiceRouter
	var userAuthority *base_spec.UserAuthority
	var repBytes []byte

	// 初回のQueryにより認証を行う
	mt, message, err := wsConn.ReadMessage()
	if err != nil {
		return
	}
	rawReq := []byte(message)

	var req *base_protocol.Request
	var res *base_protocol.Response
	var bytes []byte
	service, userAuthority, req, res, err = app.Start(tctx, r, rawReq, isProxy)
	if err != nil {
		bytes, err = json.Marshal(&res)
		if err != nil {
			logger.Error(tctx, err, "Failed json.Marshal")
			return
		}
		err = wsConn.WriteMessage(mt, bytes)
		return
	}

	for _, endpoint := range service.Endpoints {
		if endpoint == "" {
			if err = app.queryHandler.ExecWs(tctx, userAuthority, r, w, req, res, nil); err != nil {
				bytes, err = json.Marshal(&res)
				if err != nil {
					logger.Error(tctx, err, "Failed json.Marshal")
					return
				}
				err = wsConn.WriteMessage(mt, bytes)
				return
			}
			repBytes, err = json.Marshal(&res)
			err = wsConn.WriteMessage(mt, repBytes)
			if err != nil {
				logger.Error(tctx, err, "Failed Write First Message")
				return
			}
			break
		}

		// プロキシ先とWsをつなげる
		endpoint := strings.Replace(endpoint, "http", "ws", -1) + "/wp"
		dialer := websocket.Dialer{
			Proxy:            http.ProxyFromEnvironment,
			HandshakeTimeout: 45 * time.Second,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
		header := http.Header{}
		header.Add("X-Service-Token", service.Token)
		proxyWsConn, _, err = dialer.Dial(endpoint, header)
		if err != nil {
			logger.Errorf(tctx, err, "Failed dial to %s", endpoint)
			continue
		}
		break
	}

	if proxyWsConn != nil { // Proxy
		isInitProxy := true
		for {
			if isInitProxy {
				if err = proxyWsConn.WriteMessage(mt, rawReq); err != nil {
					logger.Warningf(tctx, "Failed WriteMessage: %s", err.Error())
					return
				}
				go func() {
					for {
						mt, message, tmpErr := proxyWsConn.ReadMessage()
						if tmpErr != nil {
							conMutex.Lock()
							if !wsDone {
								logger.Warningf(tctx, "Failed proxyWsConn.ReadMessage: err=%s", tmpErr.Error())
							}
							conMutex.Unlock()
							return
						}
						if tmpErr := wsConn.WriteMessage(mt, message); tmpErr != nil {
							conMutex.Lock()
							if !wsDone {
								logger.Warningf(tctx, "Failed proxyWsConn.WriteMessage: err=%s", tmpErr.Error())
							}
							conMutex.Unlock()
							return
						}
					}
				}()
				isInitProxy = false
			}

			fmt.Println("Proxy: Waiting Message from client")
			mt, message, tmpErr := wsConn.ReadMessage()
			fmt.Println("Proxy: Recieved Message", string(message))
			if tmpErr != nil {
				logger.Warningf(tctx, "Failed ReadMessage: %s", tmpErr.Error())
				fmt.Println("DEBUG Failed ReadMessage: ", tmpErr.Error())
				return
			}
			rawReq := []byte(message)
			if err = proxyWsConn.WriteMessage(mt, rawReq); err != nil {
				logger.Warningf(tctx, "Failed WriteMessage: %s", err.Error())
			}
			continue
		}
	} else { // App handle ws
		if err = app.queryHandler.ExecWs(tctx, userAuthority, r, w, req, res, wsConn); err != nil {
			return
		}
	}
}

func (app *BaseApp) NewHandler() http.Handler {
	handler := http.NewServeMux()
	handler.HandleFunc("/q", func(w http.ResponseWriter, r *http.Request) {
		app.ExecQuery(w, r, false)
	})
	handler.HandleFunc("/p", func(w http.ResponseWriter, r *http.Request) {
		app.ExecQuery(w, r, true)
	})
	handler.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		app.Ws(w, r, false)
	})
	handler.HandleFunc("/wp", func(w http.ResponseWriter, r *http.Request) {
		app.Ws(w, r, true)
	})

	return handler
}

func (app *BaseApp) Start(tctx *logger.TraceContext, httpReq *http.Request, rawReq []byte, isProxy bool) (service *base_spec_model.ServiceRouter,
	userAuthority *base_spec.UserAuthority, req *base_protocol.Request, res *base_protocol.Response, err error) {
	res = &base_protocol.Response{TraceId: tctx.GetTraceId(), ResultMap: map[string]base_protocol.Result{}}

	req = &base_protocol.Request{}
	if err = json.Unmarshal(rawReq, &req); err != nil {
		res.Code = base_const.CodeServerInternalError
		res.Error = err.Error()
		return
	}

	tmpService, ok := app.serviceMap[req.Service]
	if !ok {
		res.Code = base_const.CodeClientBadRequest
		queries := []string{}
		for _, query := range req.Queries {
			queries = append(queries, query.Name)
		}
		err = fmt.Errorf("InvalidService: service=%s, queries=%v", req.Service, queries)
		res.Error = err.Error()
		return
	}

	var token string
	if !isProxy {
		token = httpReq.Header.Get("X-Auth-Token")
		if token == "" {
			if tokenCookie, tmpErr := httpReq.Cookie("X-Auth-Token"); tmpErr == nil {
				token = tokenCookie.Value
			}
		}
	} else {
		token = httpReq.Header.Get("X-Service-Token")
		if token == "" {
			res.Code = base_const.CodeClientInvalidAuth
			err = fmt.Errorf("InvalidServiceToken")
			res.Error = err.Error()
			return
		}
	}
	var tokenErr error
	if token != "" {
		userAuthority, tokenErr = app.dbApi.LoginWithToken(tctx, token)
		if tokenErr != nil {
			logger.Warningf(tctx, "Failed LoginWithToken: %v", tokenErr)
		}
	}

	for _, query := range req.Queries {
		queryModel, ok := tmpService.QueryMap[query.Name]
		if !ok {
			res.Code = base_const.CodeClientBadRequest
			err = fmt.Errorf("InvalidQuery: %s", query.Name)
			res.Error = err.Error()
			return
		}

		if queryModel.RequiredAuth {
			if userAuthority == nil {
				err = fmt.Errorf("InvalidAuth: InvalidToken")
				res.Code = base_const.CodeClientInvalidAuth
				res.Error = err.Error()
				return
			}

			if queryModel.RequiredProject {
				if req.Project == "" {
					err = fmt.Errorf("InvalidProject")
					res.Code = base_const.CodeClientInvalidAuth
					res.Error = err.Error()
					return
				}
				userAuthority.ProjectName = req.Project
			}

			if !isProxy {
				if queryModel.RequiredProject {
					if projectService, ok := userAuthority.ProjectServiceMap[req.Project]; !ok {
						err = fmt.Errorf("InvalidAuthProjectService: project=%s", req.Project)
						res.Code = base_const.CodeClientInvalidAuth
						res.Error = err.Error()
					} else {
						if _, ok := projectService.ServiceMap[req.Service]; !ok {
							err = fmt.Errorf("InvalidAuthService: service=%s", req.Service)
							res.Code = base_const.CodeClientInvalidAuth
							res.Error = err.Error()
							return
						}
					}
				} else {
					if _, ok := userAuthority.ServiceMap[req.Service]; !ok {
						err = fmt.Errorf("InvalidAuthService")
						res.Code = base_const.CodeClientInvalidAuth
						res.Error = err.Error()
						return
					}
				}
			}
		}
	}
	service = &tmpService
	return
}

func (app *BaseApp) End(tctx *logger.TraceContext, startTime time.Time, err error) {
	logger.EndTrace(tctx, startTime, err, 2)
}

// SecureHeaders adds secure headers to the API
// func (a *API) SecureHeaders(next http.Handler) http.Handler {
// return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
func (app *BaseApp) ValidateHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Add X-XSS-Protection header
		// Enables XSS filtering. Rather than sanitizing the page, the browser will prevent rendering of the page if an attack is detected.
		c.Writer.Header().Add("X-XSS-Protection", "1; mode=block")

		// Add Content-Type header
		// Content type tells the browser what type of content you are sending. If you do not include it, the browser will try to guess the type and may get it wrong.
		// w.Header().Add("Content-Type", "application/json")

		// Add X-Content-Type-Options header
		// Content Sniffing is the inspecting the content of a byte stream to attempt to deduce the file format of the data within it.
		// Browsers will do this to try to guess at the content type you are sending.
		// By setting this header to “nosniff”, it prevents IE and Chrome from content sniffing a response away from its actual content type. This reduces exposure to drive-by download attacks.
		c.Writer.Header().Add("X-Content-Type-Options", "nosniff")

		// Prevent page from being displayed in an iframe
		c.Writer.Header().Add("X-Frame-Options", "DENY")

		// Allow Origin
		c.Writer.Header().Add("Access-Control-Allow-Origin", app.appConf.AccessControlAllowOrigin)
		c.Writer.Header().Add("Access-Control-Allow-Credentials", "true")
	}
}

func (app *BaseApp) Proxy(tctx *logger.TraceContext, service *base_spec_model.ServiceRouter, endpoint string, rawReq []byte) (repBytes []byte, statusCode int, err error) {
	var httpResp *http.Response
	reqBuffer := bytes.NewBuffer(rawReq)
	var httpReq *http.Request
	if httpReq, err = http.NewRequest("POST", endpoint+"/p", reqBuffer); err != nil {
		return
	}
	httpReq.Header.Add("X-Service-Token", service.Token)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{
		Transport: tr,
	}

	httpResp, err = httpClient.Do(httpReq)
	if err != nil {
		return
	}

	defer func() {
		if tmpErr := httpResp.Body.Close(); tmpErr != nil {
			logger.Errorf(tctx, err, "Failed httpResp.Body.Close()")
		}
	}()
	statusCode = httpResp.StatusCode
	repBytes, err = ioutil.ReadAll(httpResp.Body)
	return
}
