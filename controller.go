package main

import (
	"net/http"

	"github.com/m0t0k1ch1/ksatriya"
)

type Controller struct {
	*ksatriya.Controller
}

func NewController() *Controller {
	c := &Controller{ksatriya.NewController()}

	c.AddBeforeFilter(c.Before)
	c.AddAfterFilter(c.After)

	c.GET("/", c.Index)
	c.GET("/user/:name", c.User)

	return c
}

func (c *Controller) Before(ctx *ksatriya.Context) {
	ctx.SetBaseTmplPath("layout.html")
}

func (c *Controller) After(ctx *ksatriya.Context) {
	ctx.SetRenderArg("title", "ksatriya-sample")
}

func (c *Controller) Index(ctx *ksatriya.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func (c *Controller) User(ctx *ksatriya.Context) {
	name := ctx.Param("name")
	ctx.HTML(http.StatusOK, "user.html", ksatriya.RenderArgs{
		"name": name,
	})
}
