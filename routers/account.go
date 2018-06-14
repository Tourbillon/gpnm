// Copyright (c) 2018-present Anbillon Team (anbillonteam@gmail.com).

package routers

import (
	"net/http"
	"time"

	"anbillon.com/gpnm/middleware"
	"anbillon.com/gpnm/models"
	"anbillon.com/gpnm/modules/build"
	"anbillon.com/gpnm/modules/setting"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// DoLogin login with name and password.
func DoLogin(c echo.Context) error {
	ctx := c.(*middleware.AppContext)

	name := c.FormValue("username")
	password := c.FormValue("password")
	var user models.User
	if err := ctx.User.SelectByName(&user, name); err != nil {
		return showLoginError(c)
	}

	if !user.ValidatePassword(password) {
		return showLoginError(c)
	}

	expires := time.Now().Add(1 * time.Hour)
	claims := jwt.MapClaims{
		"id":    user.Id,
		"exp":   expires.Unix(),
		"admin": user.Role == models.RoleAdmin,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(setting.Config.JwtSecret))
	if err != nil {
		return showLoginError(c)
	}

	cookie := new(http.Cookie)
	cookie.Name = setting.AccessToken
	cookie.Value = accessToken
	cookie.Path = "/gpnm"
	cookie.Expires = expires
	cookie.Secure = !build.Debug
	cookie.HttpOnly = true
	c.SetCookie(cookie)

	return c.Redirect(http.StatusFound, "/gpnm/package/home")
}

func DoLogout(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = setting.AccessToken
	cookie.Value = ""
	cookie.Path = "/gpnm"
	cookie.MaxAge = -1
	c.SetCookie(cookie)

	return c.Redirect(http.StatusFound, "/gpnm/account/login")
}

func AddUser(c echo.Context) error {
	ctx := c.(*middleware.AppContext)

	name := c.FormValue("username")
	password := c.FormValue("password")
	user := models.User{
		Name:     name,
		Password: password,
		Salt:     models.GenRandString(10),
		Role:     models.RoleNormal,
	}
	user.EncryptPassword()
	if _, err := ctx.User.InsertOne(&user); err != nil {
		return err
	}

	return c.Redirect(http.StatusFound, "/gpnm/package")
}

func DoModifyPassword(c echo.Context) error {
	ctx := c.(*middleware.AppContext)

	uid, ok := c.Get("uid").(float64)
	if !ok {
		return showSetting(c, true)
	}
	id := int32(uid)

	oldPasswd := c.FormValue("old_password")
	newPasswd := c.FormValue("new_password")
	if len(oldPasswd) == 0 || len(newPasswd) == 0 {
		return showSetting(c, true)
	}

	var storedUser models.User
	if err := ctx.User.SelectById(&storedUser, id); err != nil {
		return c.String(http.StatusNotFound, "Not found")
	}

	if !storedUser.ValidatePassword(oldPasswd) {
		return showSetting(c, true)
	}

	user := models.User{
		Id:       id,
		Password: newPasswd,
		Salt:     storedUser.Salt,
	}
	user.EncryptPassword()

	if _, err := ctx.User.UpdateById(&user); err != nil {
		return err
	}

	return c.Redirect(http.StatusFound, "/gpnm/package/home")
}

func showLoginError(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", map[string]interface{}{
		"ShowError": true,
	})
}

func showSetting(c echo.Context, showError bool) error {
	isAdmin, ok := c.Get("admin").(bool)
	data := map[string]interface{}{
		"ShowError": showError,
	}
	if ok {
		data["IsAdmin"] = isAdmin
	}

	return c.Render(http.StatusOK, "setting.html", data)
}

// ShowLogin will show login html page.
func ShowLogin(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", nil)
}

// ShowSetting will show account setting html page.
func ShowSetting(c echo.Context) error {
	return showSetting(c, false)
}
