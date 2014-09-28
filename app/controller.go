package ksatriyasample

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

func (c *Controller) Before(ctx ksatriya.Ctx) {
	ctx.SetTmplDirPath("app/view")
	ctx.SetBaseTmplPath("layout.html")
}

func (c *Controller) After(ctx ksatriya.Ctx) {
	ctx.SetRenderArg("title", "ksatriya-sample")
}

func (c *Controller) Index(ctx ksatriya.Ctx) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func (c *Controller) User(ctx ksatriya.Ctx) {
	name := ctx.Param("name")
	ctx.HTML(http.StatusOK, "user.html", ksatriya.RenderArgs{
		"name": name,
	})
}
