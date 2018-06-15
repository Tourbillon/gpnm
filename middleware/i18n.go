// Copyright (c) 2018-present Anbillon Team (anbillonteam@gmail.com).

package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var bundle *i18n.Bundle

func init() {
	bundle = &i18n.Bundle{
		DefaultLanguage: language.English,
	}
	if _, err := bundle.LoadMessageFile("public/locale/active.zh.json"); err != nil {
		log.Fatalf("%v", err)
	}
}

// I18n check the request to get accepted language, then set
// localizer for renderer to use.
func I18n() echo.MiddlewareFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			accept := c.Request().Header.Get("Accept-Language")
			localizer := i18n.NewLocalizer(bundle, accept)
			c.Set("localizer", localizer)
			return handlerFunc(c)
		}
	}
}
