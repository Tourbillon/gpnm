// Copyright (c) 2018-present Anbillon Team (anbillonteam@gmail.com).

package render

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type TemplateRender struct {
	tpl       *template.Template
	localizer *i18n.Localizer
}

// NewRenderer returns a new TemplateRender.
func NewRenderer() *TemplateRender {
	r := TemplateRender{}
	funcMap := template.FuncMap{
		"minus": r.minus,
		"add":   r.add,
		"tr":    r.translate,
	}
	r.tpl = template.Must(
		template.New("").Funcs(funcMap).ParseGlob("public/template/*.html"))

	return &r
}

// Render implement Renderer in echo.
func (t *TemplateRender) Render(w io.Writer, name string,
	data interface{}, c echo.Context) error {
	localizer, ok := c.Get("localizer").(*i18n.Localizer)
	if ok {
		t.localizer = localizer
	}
	return t.tpl.ExecuteTemplate(w, name, data)
}

func (t *TemplateRender) translate(id string) string {
	if t.localizer == nil {
		return id
	}

	l, err := t.localizer.Localize(&i18n.LocalizeConfig{
		MessageID: id,
	})
	if err != nil {
		return id
	}
	return l
}

func (t *TemplateRender) minus(x, y int64) int64 {
	return x - y
}

func (t *TemplateRender) add(x, y int64) int64 {
	return x + y
}
