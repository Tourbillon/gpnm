// Copyright (c) 2018-present Anbillon Team (anbillonteam@gmail.com).

package cmd

import (
	mw "anbillon.com/gpnm/middleware"
	"anbillon.com/gpnm/modules/build"
	"anbillon.com/gpnm/modules/setting"
	"anbillon.com/gpnm/render"
	"anbillon.com/gpnm/routers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/cobra"
)

// newServerStartCmd will create a new command to start gpnm web server.
func newServerStartCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start gpnm web server",
		RunE:  runWebServer,
	}
}

// runWebServer is the run command to start gpnm web server.
func runWebServer(_ *cobra.Command, _ []string) error {
	e := echo.New()
	e.HideBanner = true

	e.Use(mw.AppCtx())
	e.Use(mw.I18n())
	if !build.Debug {
		e.Use(middleware.Recover())
	}
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())

	e.Renderer = render.NewRenderer()

	e.Static("/gpnm/js", "public/js")
	e.File("/favicon.ico", "public/favicon.ico")

	gpnm := e.Group("/gpnm")

	// a group which needs authorization checker
	pkg := gpnm.Group("/package", mw.AuthChecker())
	pkg.GET("", routers.ShowAllPackages)
	pkg.GET("/home", routers.ShowAllPackages)
	pkg.GET("/add", routers.ShowAddPackage)
	pkg.GET("/modify", routers.ShowModifyPackage)
	pkg.POST("/add", routers.DoAddPackage)
	pkg.POST("/delete", routers.DoDeletePackage)
	pkg.POST("/modify", routers.DoModifyPackage)

	account := gpnm.Group("/account")
	account.GET("/login", routers.ShowLogin, mw.AuthCheckerWithLogin())
	account.POST("/login", routers.DoLogin)
	account.GET("/logout", routers.DoLogout)
	accountToken := account.Group("", mw.AuthChecker())
	accountToken.POST("/add", routers.AddUser)
	accountToken.GET("/setting", routers.ShowSetting)
	accountToken.POST("/setting", routers.DoModifyPassword)

	e.GET("/*", routers.ShowPackage)

	return e.Start(setting.Config.Address)
}
