package base_app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_model"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

var serviceMap = map[string]base_model.ServiceRouter{
	"Auth": base_model.ServiceRouter{
		Endpoints: []string{"self"},
		QueryMap: map[string]base_model.QueryModel{
			"Login": base_model.QueryModel{},
		},
	},
}

func (app *BaseApp) Start(r *http.Request) (*logger.TraceContext, *base_model.ServiceRouter, *base_model.Request, *base_model.Response, time.Time, error) {
	var err error
	tctx := logger.NewTraceContext(app.host, app.name)
	startTime := logger.StartTrace(tctx)
	res := base_model.Response{TraceId: tctx.GetTraceId(), Data: map[string]interface{}{}}

	var req base_model.Request
	bufbody := new(bytes.Buffer)
	bufbody.ReadFrom(r.Body)

	if err = json.Unmarshal(bufbody.Bytes(), &req); err != nil {
		res.Code = base_const.CodeServerInternalError
		res.Error = err.Error()
		return tctx, nil, nil, &res, startTime, err
	}

	service, ok := serviceMap[req.Service]
	if !ok {
		res.Code = base_const.CodeClientBadRequest
		err = fmt.Errorf("InvalidService")
		res.Error = err.Error()
		return tctx, nil, nil, &res, startTime, err
	}

	for _, query := range req.Queries {
		queryModel, ok := service.QueryMap[query.Name]
		if !ok {
			res.Code = base_const.CodeClientBadRequest
			err = fmt.Errorf("InvalidQuery")
			res.Error = err.Error()
			return tctx, nil, nil, &res, startTime, err
		}

		if queryModel.RequiredAuth {
			fmt.Println("Valid Auth")
		}
	}

	return tctx, &service, &req, &res, startTime, err
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
