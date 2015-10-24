package app

import (
	"net/http"

	"github.com/m0t0k1ch1/ksatriya"
)

type Controller struct {
	*ksatriya.Controller
}

func (c *Controller) Before(ctx *ksatriya.Context) {
	if ctx.Req().URL.Path == "/redirect" {
		ctx.Redirect("/")
	}
}

func (c *Controller) Index(ctx *ksatriya.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func (c *Controller) User(ctx *ksatriya.Context) {
	name := ctx.Arg("name")
	ctx.HTML(http.StatusOK, "user.html", ksatriya.RenderArgs{
		"name": name,
	})
}

func (c *Controller) Text(ctx *ksatriya.Context) {
	ctx.Text(http.StatusOK, "sample text")
}

func (c *Controller) JSON(ctx *ksatriya.Context) {
	ctx.JSON(http.StatusOK, map[string]string{
		"key": "value",
	})
}

func (c *Controller) Redirect(ctx *ksatriya.Context) {}

func (c *Controller) After(ctx *ksatriya.Context) {
	v := ctx.View()
	v.SetRenderArg("title", "ksatriya-sample")

	conf := v.RenderConfig()
	conf.SetTmplDirPath("app/view")
	conf.SetBaseTmplPath("layout.html")
}

func NewController() *Controller {
	c := &Controller{ksatriya.NewController()}

	c.AddBeforeFilter(c.Before)
	c.GET("/", c.Index)
	c.GET("/user/:name", c.User)
	c.GET("/text", c.Text)
	c.GET("/json", c.JSON)
	c.GET("/redirect", c.Redirect)
	c.AddAfterFilter(c.After)

	return c
}
