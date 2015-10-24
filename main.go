package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/braintree/manners"
	"github.com/lestrrat/go-server-starter/listener"
	"github.com/m0t0k1ch1/ksatriya-sample/app"
)

func main() {
	app := app.New()

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

	var l net.Listener
	listeners, err := listener.ListenAll()
	if err != nil {
		if err == listener.ErrNoListeningTarget {
			l, err = net.Listen("tcp", ":8080")
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
	} else {
		l = listeners[0]
	}

	manners.Serve(l, app)
}
