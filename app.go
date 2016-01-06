package main

import (
	"net/http"

	"github.com/m0t0k1ch1/ksatriya"
)

type HandlerFunc func(*Context)

type MyKsatriya struct {
	*ksatriya.Ksatriya
	conf *Config
}

func (mk *MyKsatriya) NewContext(w http.ResponseWriter, req *http.Request, args ksatriya.Args) ksatriya.Ctx {
	return &Context{
		Ctx:  ksatriya.NewContext(w, req, args),
		conf: mk.conf,
	}
}

func New(confPath string) *MyKsatriya {
	mk := &MyKsatriya{
		Ksatriya: ksatriya.New(),
		conf:     NewConfig(confPath),
	}

	mk.SetCtxBuilder(mk.NewContext)
	mk.RegisterController(NewAPIController())

	return mk
}
