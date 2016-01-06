package main

import "github.com/m0t0k1ch1/ksatriya"

type BaseController struct {
	*ksatriya.Controller
}

func (c *BaseController) AddMyRoute(method, path string, hf HandlerFunc) {
	c.AddRoute(method, path, func(kctx ksatriya.Ctx) {
		ctx, _ := kctx.(*Context)
		hf(ctx)
	})
}

func NewBaseController() *BaseController {
	return &BaseController{ksatriya.NewController()}
}
