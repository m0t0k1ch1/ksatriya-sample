package main

import "net/http"

type APIController struct {
	*BaseController
}

func (c *APIController) Ping(ctx *Context) {
	ctx.RenderJSON(http.StatusOK, NewResponse(ctx.conf.Message))
}

func NewAPIController() *APIController {
	c := &APIController{NewBaseController()}

	c.AddMyRoute("GET", "/api/ping", c.Ping)

	return c
}
