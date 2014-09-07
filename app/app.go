package ksatriyasample

import (
	"github.com/codegangsta/negroni"
	"github.com/m0t0k1ch1/ksatriya"
)

func NewApp() *negroni.Negroni {
	k := ksatriya.New()
	k.RegisterController(NewController())

	n := negroni.Classic()
	n.UseHandler(k)

	return n
}
