package controller

import (
	"GinDemo1/src/base"
	"GinDemo1/src/pojo/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserLogin(c *gin.Context) {
	var respJson string
	defer func() {

		c.JSON(http.StatusOK, respJson)
	}()
	name := c.PostForm("name")
	password := c.PostForm("password")
	nick := c.DefaultPostForm("password", "admin")
	fmt.Println("name:", name, "password:", password)
	fmt.Println("nick:", nick)
	login := user.Login(password, name)
	if login {
		respJson = base.RespJson(200, "success", "", 0)
	} else {
		respJson = base.RespJson(500, "fail", "", 0)
	}
}
