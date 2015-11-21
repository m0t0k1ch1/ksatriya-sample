package main

import (
	"net/http"

	"github.com/m0t0k1ch1/ksatriya"
)

type Context struct {
	ksatriya.Ctx
	client *http.Client
}

func NewContext(w http.ResponseWriter, req *http.Request, args ksatriya.Args) ksatriya.Ctx {
	return &Context{
		Ctx:    ksatriya.NewContext(w, req, args),
		client: &http.Client{},
	}
}

func convertContext(kctx ksatriya.Ctx) *Context {
	ctx, _ := kctx.(*Context)
	return ctx
}
