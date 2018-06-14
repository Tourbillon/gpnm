// Copyright (c) 2018-present Anbillon Team (anbillonteam@gmail.com).

package middleware

import (
	"anbillon.com/gpnm/models"
	"github.com/labstack/echo"
)

type AppContext struct {
	echo.Context
	PkgInfo *models.PkgInfoBrick
	User    *models.UserBrick
}

// AppCtx initializes a classic context for a request.
func AppCtx() echo.MiddlewareFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &AppContext{
				Context: c,
				PkgInfo: models.GetSqlBrick().PkgInfo,
				User:    models.GetSqlBrick().User,
			}
			return handlerFunc(cc)
		}
	}
}
