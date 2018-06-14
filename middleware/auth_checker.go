// Copyright (c) 2018-present Anbillon Team (anbillonteam@gmail.com).

package middleware

import (
	"net/http"

	"anbillon.com/gpnm/modules/setting"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// AuthChecker will check if authorization available.
func AuthChecker() echo.MiddlewareFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return checkCookie(c, handlerFunc, false)
		}
	}
}

// AuthCheckerWithLogin will check if authorization available when login.
func AuthCheckerWithLogin() echo.MiddlewareFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return checkCookie(c, handlerFunc, true)
		}
	}
}

func checkCookie(c echo.Context, handlerFunc echo.HandlerFunc, isLogin bool) error {
	cookie, err := c.Cookie(setting.AccessToken)
	if err != nil {
		if isLogin {
			return handlerFunc(c)
		}
		return c.Redirect(http.StatusFound, "/gpnm/account/login")
	}

	accessToken := cookie.Value
	token, err := jwt.Parse(accessToken, func(jt *jwt.Token) (interface{}, error) {
		return []byte(setting.Config.JwtSecret), nil
	})
	if err != nil || !token.Valid {
		return c.Redirect(http.StatusFound, "/gpnm/account/login")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Redirect(http.StatusFound, "/gpnm/account/login")
	}

	// the user has grant access token
	if isLogin {
		return c.Redirect(http.StatusFound, "/gpnm/package")
	}

	c.Set("uid", claims["id"])
	c.Set("admin", claims["admin"])

	return handlerFunc(c)
}
