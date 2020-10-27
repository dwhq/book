package main

import (
	_"demo/config"
	"demo/routes"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func setupRouter() *gin.Engine {
	//r := gin.Default()
	r := routes.NewRouter()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.String(200, "pong")
	//})
	//r.GET("/someJSON", func(c *gin.Context) {
	//	data := map[string]interface{}{
	//		"lang": "GO语言",
	//		"tag":  "<br>",
	//	}
	//
	//	// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
	//	c.AsciiJSON(http.StatusOK, data)
	//})
	//r.GET("/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.html", gin.H{
	//		"title": "Main website 😍",
	//	})
	//})
	//r.GET("/login", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "login.html", gin.H{})
	//})
	r.GET("/someGet", getting)
	//r.POST("/login", func(c *gin.Context) {
	//	// 你可以使用显式绑定声明绑定 multipart form：
	//	// c.ShouldBindWith(&form, binding.Form)
	//	// 或者简单地使用 ShouldBind 方法自动绑定：
	//	var form  LoginForm
	//	form.User = c.PostForm("user")
	//	form.Password = c.PostForm("password")
	//	//c.AsciiJSON(http.StatusOK, form)
	//	if form.User == "user" && form.Password == "password" {
	//		c.JSON(200, gin.H{"status": "you are logged in"})
	//	} else {
	//		c.JSON(401, gin.H{"status": "unauthorized"})
	//	}
	//})
	return r
}

func getting(c *gin.Context)  {
	data := map[string]interface{}{
		"tag":  "<br>",
		"test": "测试",
	}
	c.PureJSON(http.StatusOK,data)
}


func main() {

	r := routes.NewRouter()//路由

	r.Run(":"+os.Getenv("APP_PORT"))
}
