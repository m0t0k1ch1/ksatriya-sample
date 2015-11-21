package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/m0t0k1ch1/ksatriya"
)

func PingHandler(kctx ksatriya.Ctx) {
	Ping(convertContext(kctx))
}
func Ping(ctx *Context) {
	res, err := ctx.client.Get("http://127.0.0.1:8080/pong")
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	ctx.Text(http.StatusOK, string(body))
}

func PongHandler(ctx ksatriya.Ctx) {
	ctx.Text(http.StatusOK, "pong")
}

func NewApp() *negroni.Negroni {
	k := ksatriya.New()
	k.SetCtxBuilder(NewContext)
	k.RegisterController(NewController())

	k.GET("/ping", PingHandler)
	k.GET("/pong", PongHandler)

	n := negroni.Classic()
	n.UseHandler(k)

	return n
}
