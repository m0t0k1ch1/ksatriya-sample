package main

import "github.com/m0t0k1ch1/ksatriya"

type Context struct {
	ksatriya.Ctx
	conf *Config
}

func convertContext(kctx ksatriya.Ctx) *Context {
	ctx, _ := kctx.(*Context)
	return ctx
}
