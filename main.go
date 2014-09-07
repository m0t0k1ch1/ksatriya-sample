package main

import (
	"net"
	"net/http"

	"github.com/lestrrat/go-server-starter-listener"
	"github.com/m0t0k1ch1/ksatriya-sample/app"
)

func main() {
	app := ksatriyasample.NewApp()

	listener, _ := ss.NewListener()
	if listener == nil {
		var err error
		listener, err = net.Listen("tcp", ":8080")
		if err != nil {
			panic(err)
		}
	}

	server := &http.Server{Handler: app}
	server.Serve(listener)
}
