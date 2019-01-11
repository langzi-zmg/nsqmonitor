package api

//import (
//	"github.com/labstack/echo"
//	"gitlab.wallstcn.com/operation/nsqmonitor/business"
//	"gitlab.wallstcn.com/operation/nsqmonitor/helper"
//)
//
//
//
//type Pagination struct {
//	Page  int64 `json:"page" query:"page"`
//	Limit int64 `json:"limit" query:"limit"`
//}
//
//
//// @Title overview  list
//// @Description 获取overview list
//// @Accept  json
//// @Param page query int false "页数|默认1"
//// @Param limit query int false "每页条目数|默认10"
//// @Resource overview
//// @Router /v1/overview [get]
//
//
//func HTTPGetOverview(ctx echo.Context) error {
//
//	return nil
//}
//
//// @Title consumer list by page and limit
//// @Description 获取consumer list by page and limit
//// @Accept  json
//// @Param page query int false "页数|默认1"
//// @Param limit query int false "每页条目数|默认10"
//// @Resource consumer
//// @Router /v1/consumer  [get]
//
//func HTTPGetConsumer(ctx echo.Context) error {
//
//	return helper.SuccessResponse(ctx, business.ConsumerList)
//}
