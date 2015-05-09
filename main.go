package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/braintree/manners"
	"github.com/lestrrat/go-server-starter/listener"
	"github.com/m0t0k1ch1/ksatriya-sample/app"
)

func main() {
	app := ksatriyasample.NewApp()

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGTERM)
	go func() {
		for {
			s := <-signalChan
			if s == syscall.SIGTERM {
				manners.Close()
			}
		}
	}()

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

	manners.Serve(l, app)
}
