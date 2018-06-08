package routers

import (
	"github.com/gin-gonic/gin"
	. "gin_server/apis"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)

	router := gin.Default()

	router.GET("/", IndexApi)

	router.POST("/person", AddPersonApi)

	router.GET("/persons", GetPersonsApi)

	router.GET("/person/:id", GetPersonApi)

	router.PUT("/person/:id", ModifyPersonApi)

	router.DELETE("/person/:id", DeletePersonApi)

	return router
}
