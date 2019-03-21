package home

import (
	"github.com/gin-gonic/gin"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (srv *Home) Action(c *gin.Context) {
	tctx, err := logger.NewAuthproxyActionTraceContext(srv.host, srv.name, c)
	startTime := logger.StartTrace(&tctx.TraceContext)
	defer func() { logger.EndTrace(&tctx.TraceContext, startTime, err, 1) }()

	if err != nil {
		c.JSON(500, gin.H{
			"err": "InvalidRequest",
		})
		return
	}

	resp := map[string]interface{}{
		"Index": map[string]interface{}{
			"Kind": "Panels",
			"Panels": []interface{}{
				map[string]interface{}{
					"Name": "Hoge",
					"Kind": "Msg",
				},
				map[string]interface{}{
					"Name": "Piyo",
					"Kind": "Msg",
				},
			},
		},
	}

	c.JSON(200, resp)
}
