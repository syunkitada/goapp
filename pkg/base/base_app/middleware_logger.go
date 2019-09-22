package base_app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_model"
	"github.com/syunkitada/goapp/pkg/base/base_model/spec_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (app *BaseApp) Start(tctx *logger.TraceContext, db *gorm.DB, httpReq *http.Request) (service *spec_model.ServiceRouter,
	userAuthority *base_spec.UserAuthority, rawReq []byte, req *base_model.Request, res *base_model.Response, err error) {
	res = &base_model.Response{TraceId: tctx.GetTraceId(), ResultMap: map[string]base_model.Result{}}

	token := httpReq.Header.Get("X-Auth-Token")
	if token == "" {
		var tokenCookie *http.Cookie
		if tokenCookie, err = httpReq.Cookie("X-Auth-Token"); err != nil {
			fmt.Println("cookie is err:", err) // FIXME debug print
		} else {
			token = tokenCookie.Value
		}
	}
	var tokenErr error
	if token != "" {
		userAuthority, tokenErr = app.dbApi.LoginWithToken(tctx, db, token)
		fmt.Println("debug tokenErr", tokenErr) // FIXME debug print
	}

	req = &base_model.Request{}
	bufbody := new(bytes.Buffer)
	bufbody.ReadFrom(httpReq.Body)
	rawReq = bufbody.Bytes()
	if err = json.Unmarshal(rawReq, &req); err != nil {
		res.Code = base_const.CodeServerInternalError
		res.Error = err.Error()
		return
	}

	tmpService, ok := app.serviceMap[req.Service]
	if !ok {
		res.Code = base_const.CodeClientBadRequest
		err = fmt.Errorf("InvalidService")
		res.Error = err.Error()
		return
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
			fmt.Println("Valid Auth")
		}
	}
	service = &tmpService
	return
}

func (app *BaseApp) End(tctx *logger.TraceContext, startTime time.Time, err error) {
	logger.EndTrace(tctx, startTime, err, 2)
}

func (app *BaseApp) Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		client := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path

		tctx := logger.NewTraceContext(app.host, app.name)

		c.Set("TraceId", tctx.TraceId)
		// TODO FIX
		fmt.Println(tctx, app.host, app.name, map[string]string{
			"Msg":    "Start",
			"Client": client,
			"Method": method,
			"Path":   path,
		})

		c.Next()
		end := time.Now()
		latency := end.Sub(start)

		statusCode := c.Writer.Status()

		// TODO FIX
		fmt.Println(tctx, app.host, app.name, map[string]string{
			"Msg":       "End",
			"Client":    client,
			"Method":    method,
			"Path":      path,
			"StausCode": strconv.Itoa(statusCode),
			"Latency":   strconv.FormatInt(latency.Nanoseconds()/1000000, 10),
		})
	}
}
