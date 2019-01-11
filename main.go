package main

import (
	"gitlab.wallstcn.com/operation/nsqmonitor/common"
	"gitlab.wallstcn.com/operation/nsqmonitor/service"
)

func main() {
	common.LoadConfig("conf/nsqmonitor.yaml")
	//business.GetMine()

	service.RunServer()
}
