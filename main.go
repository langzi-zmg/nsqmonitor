package main

import (
	"gitlab.wallstcn.com/operation/nsqmonitor/common"
	"gitlab.wallstcn.com/operation/nsqmonitor/service"
	ivksvc "gitlab.wallstcn.com/wscnbackend/ivankastd/service"
	"gitlab.wallstcn.com/operation/nsqmonitor/rpcserver"
	"github.com/robfig/cron"
	"os"
	"gitlab.wallstcn.com/operation/nsqmonitor/business"
)

func main() {
	common.LoadConfig("conf/nsqmonitor.yaml")
	startService()

	//cron
	c := cron.New()
	spec := os.Getenv("CRONTIME")
	c.AddFunc(spec, func() {
		business.GetMine()
	})
	c.Start()

	service.RunServer()
}


func startService() {
	svc := ivksvc.NewService(common.GlobalConf.Micro)
	svc.Init()
	rpcserver.Init(svc)

}