package Router

import (
	"github.com/gin-gonic/gin"
	. "goTest/Services"
)

func InitRouter() *gin.Engine {
	routerGroup := gin.Default()
	router := routerGroup.Group("/video/")
	{
		router.GET("/", Index)
		router.GET("/GetKindId", GetKindId)
		router.GET("/GetAllKind", GetAllKind)
		router.POST("/AddKind", AddKind)
		router.POST("/UpdateKind", UpdateKind)
		router.GET("/GetKindCount", GetKindCount)
		router.GET("/QueryData",QueryData)

		router.GET("/QueryDataString",QueryDataString)

		router.GET("/tk",Tk)

	}
	return routerGroup
}
