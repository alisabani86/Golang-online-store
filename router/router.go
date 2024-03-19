package router

import (
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(Handler *Handler) {
	r = gin.Default()
	r.GET("/ping", func(req *gin.Context) {
		req.JSON(200, gin.H{"message": "pong"})
	})

	r.POST("/signup", Handler.CreateUser)
	r.POST("/login", Handler.Login)
	r.GET("/logout", Handler.Logout)
	r.GET("/getproductbycategory", Handler.GetProduct)
	r.GET("/addcart", Handler.AddShopingCart)
	r.GET("/getcart", Handler.GetListCart)
	r.GET("/delete", Handler.DeleteCart)
	r.GET("/checkout", Handler.checkout)
}

func Start(addr string) error {
	return r.Run(addr)
}
