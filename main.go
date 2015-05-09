package main

import (
	"net"
	"net/http"

	"github.com/lestrrat/go-server-starter/listener"
	"github.com/m0t0k1ch1/ksatriya-sample/app"
)

func main() {
	app := ksatriyasample.NewApp()

	listeners, err := listener.ListenAll()
	if err != nil {
		panic(err)
	}
	var l net.Listener
	if len(listeners) == 0 {
		l, err = net.Listen("tcp", ":8080")
		if err != nil {
			panic(err)
		}
	} else {
		l = listeners[0]
	}

	server := &http.Server{Handler: app}
	server.Serve(l)
}
