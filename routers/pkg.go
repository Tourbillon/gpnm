// Copyright (c) 2018-present Anbillon Team (anbillonteam@gmail.com).

package routers

import (
	"net/http"
	"strconv"
	"strings"

	"anbillon.com/gpnm/middleware"
	"anbillon.com/gpnm/models"
	"anbillon.com/gpnm/modules/util"
	"github.com/labstack/echo"
)

func showAddPackageError(c echo.Context) error {
	return c.Render(http.StatusOK, "add.html", map[string]interface{}{
		"ShowError": true,
	})
}

func showModifyPackageError(c echo.Context) error {
	return c.Render(http.StatusOK, "modify.html", map[string]interface{}{
		"ShowError": true,
	})
}

// DoAddPackage add a new package into database.
func DoAddPackage(c echo.Context) error {
	ctx := c.(*middleware.AppContext)

	name := c.FormValue("package_name")
	rootRepoUrl := c.FormValue("root_repo_url")

	if len(name) == 0 || len(rootRepoUrl) == 0 {
		return showAddPackageError(c)
	}

	pkgInfo := models.PkgInfo{
		Name:        name,
		RootRepoUrl: rootRepoUrl,
	}
	if _, err := ctx.PkgInfo.InsertOne(&pkgInfo); err != nil {
		return err
	}

	return c.Redirect(http.StatusFound, "/gpnm/package/home")
}

// DoDeletePackage remove one record from database.
func DoDeletePackage(c echo.Context) error {
	ctx := c.(*middleware.AppContext)
	id := util.ParseInt(c.FormValue("id"))
	ctx.PkgInfo.DeleteById(id)

	return c.Redirect(http.StatusFound, "/gpnm/package/home")
}

// DoModifyPackage modify existed package information.
func DoModifyPackage(c echo.Context) error {
	ctx := c.(*middleware.AppContext)

	id := util.ParseInt(c.FormValue("id"))
	rootRepoUrl := c.FormValue("root_repo_url")

	if id < 0 || len(rootRepoUrl) == 0 {
		return showModifyPackageError(c)
	}

	ctx.PkgInfo.UpdateById(&models.PkgInfo{
		Id:          int32(id),
		RootRepoUrl: rootRepoUrl,
	})

	return c.Redirect(http.StatusFound, "/gpnm/package/home")
}

// ShowPackage will show package with meta data in `<head>` tag.
func ShowPackage(c echo.Context) error {
	ctx := c.(*middleware.AppContext)

	host := c.Request().Host
	uri := c.Request().RequestURI
	pkgName := uri[1:]
	index := strings.Index(uri, "?")
	if index > 0 {
		pkgName = uri[1:index]
	}

	var pkgInfo models.PkgInfo
	for {
		if err := ctx.PkgInfo.SelectByName(&pkgInfo, map[string]interface{}{
			"host": host,
			"name": pkgName,
		}); err == nil {
			break
		}

		index := strings.LastIndex(pkgName, "/")
		if index <= 0 {
			return c.String(http.StatusNotFound, "Not found")
		}
		pkgName = pkgName[:index]
	}

	return c.Render(http.StatusOK, "package.html", pkgInfo)
}

// ShowAllPackages will show all packages with pagination.
func ShowAllPackages(c echo.Context) error {
	ctx := c.(*middleware.AppContext)

	page := c.QueryParam("pagination")
	size := c.QueryParam("size")
	pagination, err := strconv.ParseInt(page, 0, 0)
	if err != nil {
		pagination = 1
	}
	pageCount, err := strconv.ParseInt(size, 0, 0)
	if err != nil {
		pageCount = 10
	}
	offset := (pagination - 1) * pageCount

	var pkgs []models.PkgInfo
	if err := ctx.PkgInfo.SelectByPage(&pkgs, map[string]interface{}{
		"host":   c.Request().Host,
		"limit":  pageCount,
		"offset": offset,
	}); err != nil {
		return err
	}

	var totalPkgs int64
	if err := ctx.PkgInfo.SelectTotalPackages(&totalPkgs); err != nil {
		return err
	}
	totalPages := totalPkgs / pageCount
	if totalPkgs%pageCount > 1 {
		totalPages++
	}

	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"Packages":    pkgs,
		"CurrentPage": pagination,
		"TotalPages":  totalPages,
	})
}

// ShowAddPackage will show `add` package html page.
func ShowAddPackage(c echo.Context) error {
	return c.Render(http.StatusOK, "add.html", nil)
}

// ShowModifyPackage will show `modify` package html page.
func ShowModifyPackage(c echo.Context) error {
	ctx := c.(*middleware.AppContext)
	id := util.ParseInt(c.QueryParam("id"))

	var pkg models.PkgInfo
	if err := ctx.PkgInfo.SelectById(&pkg, map[string]interface{}{
		"host": c.Request().Host,
		"id":   id,
	}); err != nil {
		return err
	}

	return c.Render(http.StatusOK, "modify.html", map[string]interface{}{
		"Package": pkg,
	})
}
