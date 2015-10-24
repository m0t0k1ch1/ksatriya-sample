package app

import (
	"github.com/codegangsta/negroni"
	"github.com/m0t0k1ch1/ksatriya"
)

func New() *negroni.Negroni {
	k := ksatriya.New()
	k.SetContextInitializer(func(ctx *ksatriya.Context) {
		ctx.Stash["app_name"] = "ksatriya-sample"
	})
	k.RegisterController(NewController())

	n := negroni.Classic()
	n.UseHandler(k)

	return n
}
