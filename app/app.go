package app

import (
	"github.com/codegangsta/negroni"
	"github.com/m0t0k1ch1/ksatriya"
)

func New() *negroni.Negroni {
	n := negroni.Classic()

	k := ksatriya.New()
	k.RegisterController(NewController())
	n.UseHandler(k)

	return n
}
