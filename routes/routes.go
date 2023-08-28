package routes

import (
	"github.com/Damir564/HttpApiGo/controller"
	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {
	r.GET("/users", controller.GetUsers)
	r.POST("/user", controller.CreateUser)
	r.DELETE("/user/:id", controller.DeleteUser)
	r.PUT("/user/:id", controller.UpdateUser)

	r.GET("/segments", controller.GetSegments)
	r.POST("/segment", controller.CreateSegment)
	r.DELETE("/segment/", controller.DeleteSegment)
	r.PUT("/segment/", controller.UpdateSegment)

	r.POST("/bind", controller.Bind)
	r.GET("/binds", controller.GetBinds)

	r.GET("/history", controller.GetHistory)
}

// func SegmentRoute(r *gin.Engine) {
// 	r.GET("/segment", controller.GetUsers)
// 	r.POST("/segment", controller.CreateUser)
// 	r.DELETE("/segment/:id", controller.DeleteUser)
// 	r.PUT("/segment/:id", controller.UpdateUser)
// }

// func UsersSegmentsRoute(r *gin.Engine) {
// 	r.GET("/segment", controller.GetUsers)
// 	r.POST("/bind", controller.CreateUser)
// 	r.DELETE("/segment/:id", controller.DeleteUser)
// 	r.PUT("/segment/:id", controller.UpdateUser)
// }
