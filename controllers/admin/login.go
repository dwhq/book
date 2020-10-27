package admin

import (
	"demo/config"
	"demo/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context)  {
	var deletedUsers = []model.List{}
	config.DB.Table("list").Find(&deletedUsers)
	fmt.Printf("user:%#v\n",deletedUsers)

	c.HTML(http.StatusOK, "admin/login.html",gin.H{
		"title": "后台登陆2223333",
	})
}