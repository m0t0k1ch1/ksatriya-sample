package main

import (
	"net/http"

	"github.com/m0t0k1ch1/ksatriya"
)

type Controller struct {
	*ksatriya.Controller
}

func (c *Controller) BeforeHandler(ctx ksatriya.Ctx) {
	if ctx.Req().URL.Path == "/redirect" {
		ctx.Redirect("/")
	}
}

func (c *Controller) IndexHandler(ctx ksatriya.Ctx) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func (c *Controller) UserHandler(ctx ksatriya.Ctx) {
	name := ctx.Arg("name")
	ctx.HTML(http.StatusOK, "user.html", ksatriya.RenderArgs{
		"name": name,
	})
}

func (c *Controller) RedirectHandler(ctx ksatriya.Ctx) {}

func (c *Controller) AfterHandler(ctx ksatriya.Ctx) {
	v := ctx.View()
	v.SetRenderArg("title", "ksatriya-sample")

	conf := v.RenderConfig()
	conf.SetTmplDirPath("view")
	conf.SetBaseTmplPath("layout.html")
}

func NewController() *Controller {
	c := &Controller{ksatriya.NewController()}

	c.AddBeforeFilter(c.BeforeHandler)
	c.GET("/", c.IndexHandler)
	c.GET("/user/:name", c.UserHandler)
	c.GET("/redirect", c.RedirectHandler)
	c.AddAfterFilter(c.AfterHandler)

	return c
}
