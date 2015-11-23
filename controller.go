package main

import (
	"net/http"

	"github.com/m0t0k1ch1/ksatriya"
)

type Controller struct {
	*ksatriya.Controller
}

func (c *Controller) PingHandler(kctx ksatriya.Ctx) {
	ctx := convertContext(kctx)
	c.ping(ctx)
}
func (c *Controller) ping(ctx *Context) {
	ctx.RenderJSON(http.StatusOK, NewResponse(ctx.conf.Message))
}

func NewController() *Controller {
	c := &Controller{ksatriya.NewController()}

	c.GET("/ping", c.PingHandler)

	return c
}
