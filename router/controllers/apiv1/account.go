package apiv1

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/open-fightcoder/oj-web/managers"
	"github.com/open-fightcoder/oj-web/router/controllers/baseController"
)

func RegisterAccount(router *gin.RouterGroup) {
	router.POST("account/login", httpHandlerLogin)
	router.POST("account/register", httpHandlerRegister)
}

func httpHandlerLogin(c *gin.Context) {
	email := c.Query("email")
	password := c.Query("password")
	if flag, token, mess := managers.AccountLogin(email, password); flag == false {
		c.JSON(http.StatusOK, (&baseController.Base{}).Fail(mess))
	} else {
		cookie := &http.Cookie{
			Name:     "token",
			Value:    base64.StdEncoding.EncodeToString([]byte(token)),
			Path:     "/",
			HttpOnly: true,
		}
		http.SetCookie(c.Writer, cookie)
		c.JSON(http.StatusOK, (&baseController.Base{}).Success())
	}
}

func httpHandlerRegister(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	fmt.Println("===============>, ", email)
	fmt.Println("===============>, ", password)
	if flag, userId, mess := managers.AccountRegister(email, password); flag == false {
		c.JSON(http.StatusOK, (&baseController.Base{}).Fail(mess))
	} else {
		c.JSON(http.StatusOK, (&baseController.Base{}).Success(userId))
	}
}