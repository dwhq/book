package routes

import (
	"github.com/gin-gonic/gin"
	"demo/controllers"
	"demo/controllers/admin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")

	v1 := r.Group("/")
	{
		v1.GET("login", controllers.Index)
	}
	v2 := r.Group("admin/")
	{
		v2.GET("login",admin.Login)
	}
	return r
}
