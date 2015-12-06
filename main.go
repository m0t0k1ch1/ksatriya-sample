package main

import (
	"flag"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/codegangsta/negroni"
	"github.com/lestrrat/go-server-starter/listener"
	"github.com/shogo82148/go-gracedown"
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
				gracedown.Close()
			}
		}
	}()

	listeners, err := listener.ListenAll()
	if err != nil && err != listener.ErrNoListeningTarget {
		panic(err)
	}

	var l net.Listener
	if err == listener.ErrNoListeningTarget {
		l, err = net.Listen("tcp", ":8080")
		if err != nil {
			panic(err)
		}
	} else {
		l = listeners[0]
	}

	gracedown.Serve(l, handler)
}
