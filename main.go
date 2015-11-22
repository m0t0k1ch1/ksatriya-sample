package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/braintree/manners"
	"github.com/codegangsta/negroni"
	"github.com/lestrrat/go-server-starter/listener"
)

func main() {
	var confPath = flag.String("conf", "config.tml", "config file path")
	flag.Parse()

	app := New(*confPath)

	n := negroni.Classic()
	n.UseHandler(app)

	run(n)
}

func run(handler http.Handler) {
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

	manners.Serve(l, handler)
}
