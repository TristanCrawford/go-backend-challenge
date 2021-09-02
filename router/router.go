package router

import (
	"backend.com/m/v2/controllers"
	"backend.com/m/v2/database"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()
	api := &controllers.API{
		DB: database.GetDB(),
	}

	// r.POST("/person/create", func(c *gin.Context) {})
	r.GET("/person/:id", api.GetPerson)
	// r.PATCH("/person/:id/update", func(c *gin.Context) {})
	// r.DELETE("/person/:id/delete", func(c *gin.Context) {})

	return r
}