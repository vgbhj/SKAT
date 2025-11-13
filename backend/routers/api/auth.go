package api

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"github.com/vgbhj/SKAT/pkg/app"
	"github.com/vgbhj/SKAT/pkg/e"
	"github.com/vgbhj/SKAT/pkg/util"
	"github.com/vgbhj/SKAT/service/auth_service"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// @Summary SignUp
// @Produce  json
// @Param username formData string true "username"
// @Param password formData string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /signup [post]
func CreateUser(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	username := c.PostForm("username")
	password := c.PostForm("password")
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	if !ok {
		// app.MarkErrors(valid.Errors) // logging dont work
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	authService := auth_service.Auth{Username: username, Password: password}
	isNotExist, err := authService.CreateUser()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_USER, nil)
		return
	}
	if !isNotExist {
		appG.Response(http.StatusConflict, e.ERROR_ADD_USER_EXIST, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"Message": "User created successfully",
	})
}

// @Summary Login
// @Produce  json
// @Param username formData string true "username"
// @Param password formData string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /login [post]
func Login(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	username := c.PostForm("username")
	password := c.PostForm("password")
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	if !ok {
		// app.MarkErrors(valid.Errors) // logging dont work
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	authService := auth_service.Auth{Username: username, Password: password}
	isExist, err := authService.Check()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}
	if !isExist {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}

	token, err := util.GenerateToken(username, password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	c.SetCookie(
		"token",
		token,
		3600*24,
		"/",
		"",
		false,
		true,
	)

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}
