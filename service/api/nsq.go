package api

import (
	"github.com/labstack/echo"
	"gitlab.wallstcn.com/operation/nsqmonitor/helper"
	"gitlab.wallstcn.com/operation/nsqmonitor/business"
	"gopkg.in/square/go-jose.v1/json"
	"fmt"
)

type Pagination struct {
	Page  int64 `json:"page" query:"page"`
	Limit int64 `json:"limit" query:"limit"`
}


// @Title overview  list
// @Description 获取overview list
// @Accept  json
// @Param page query int false "页数|默认1"
// @Param limit query int false "每页条目数|默认10"
// @Resource overview
// @Router /v1/overview [get]


func HTTPGetOverview(ctx echo.Context) error {

	fmt.Println("====")
	jsonOverview,err := json.Marshal(business.OverviewList)
	if err != nil {
		fmt.Println(err.Error())
	}
	return helper.SuccessResponse(ctx, string(jsonOverview))

}

// @Title consumer list by page and limit
// @Description 获取consumer list by page and limit
// @Accept  json
// @Param page query int false "页数|默认1"
// @Param limit query int false "每页条目数|默认10"
// @Resource consumer
// @Router /v1/consumer  [get]

func HTTPGetConsumer(ctx echo.Context) error {
	jsonConsumer,err := json.Marshal(business.ConsumerList)
	if err != nil {
		fmt.Println(err.Error())
	}
	return helper.SuccessResponse(ctx, string(jsonConsumer))
}
