package routes

import (
	"net/http"

	"github.com/Damir564/HttpApiGo/controller"
	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func UserRoute(r *gin.Engine) {
	r.GET("/", controller.GetUsers)
	r.POST("/", controller.CreateUser)
	r.DELETE("/:id", controller.DeleteUser)
	r.PUT("/:id", controller.UpdateUser)
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

}
