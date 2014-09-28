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
	c.GET("/text", c.Text)
	c.GET("/json", c.JSON)
	c.GET("/redirect", c.Redirect)

	return c
}

func (c *Controller) Before(ctx ksatriya.Ctx) {
	ctx.View().SetTmplDirPath("app/view")
	ctx.View().SetBaseTmplPath("layout.html")
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

func (c *Controller) Text(ctx ksatriya.Ctx) {
	ctx.Text(http.StatusOK, "sample text")
}

func (c *Controller) JSON(ctx ksatriya.Ctx) {
	ctx.JSON(http.StatusOK, map[string]string{
		"key": "value",
	})
}

func (c *Controller) Redirect(ctx ksatriya.Ctx) {
	ctx.Redirect("/")
}
