package router

import (
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(Handler *Handler) {
	r = gin.Default()

	r.POST("/signup", Handler.CreateUser)
	r.POST("/login", Handler.Login)
	r.GET("/logout", Handler.Logout)
	r.GET("/tes", Handler.GetCookie)
	r.GET("/getproductbycategory", Handler.GetProduct)
}

func Start(addr string) error {
	return r.Run(addr)
}
