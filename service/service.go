package service

import (
	"fmt"
	"github.com/labstack/echo"
	"gitlab.wallstcn.com/operation/nsqmonitor/service/middleware"
	"gitlab.wallstcn.com/operation/nsqmonitor/common"
	"gitlab.wallstcn.com/operation/nsqmonitor/business"
)

var app = echo.New()

func init() {
	app.Use(middleware.RequestCORS())

	MountAPIModule(app)
}

// RunServer starts a server
func RunServer() {
	if common.GlobalConf.CertPem != "" && common.GlobalConf.KeyPem != "" {
		app.StartTLS(common.GlobalConf.Bind, common.GlobalConf.CertPem, common.GlobalConf.KeyPem)
	} else {
		fmt.Println(app.Start(common.GlobalConf.Bind))
	}

}

func MountAPIModule(e *echo.Echo) {
	apiv1 := e.Group("/v1")
	MountAPI(apiv1)
}

func MountAPI(group *echo.Group) {
	group.GET("/overview", business.HTTPGetOverview)
	group.GET("/consumer", business.HTTPGetConsumer)
}
