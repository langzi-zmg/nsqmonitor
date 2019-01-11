package api

import "github.com/labstack/echo"
// MountAPI registers API
func MountAPI(group *echo.Group) {
	MountUserInfoAPI(group)
}

func MountUserInfoAPI(group *echo.Group) {
	task := group.Group("/")
	task.GET("/overview",HTTPGetOverview)
	task.GET("/consumer",HTTPGetConsumer )


}
