package main

import (
	"net"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/lestrrat/go-server-starter-listener"
	"github.com/m0t0k1ch1/ksatriya"
)

func Index(c *ksatriya.Context) {
	c.RenderHTML(http.StatusOK, "index", nil)
}

func User(c *ksatriya.Context) {
	name := c.Param("name")
	c.RenderHTML(http.StatusOK, "user", ksatriya.RenderData{
		"name": name,
	})
}

func main() {
	k := ksatriya.New()
	k.GET("/", Index)
	k.GET("/user/:name", User)

	n := negroni.Classic()
	n.UseHandler(k)

	listener, _ := ss.NewListener()
	if listener == nil {
		var err error
		listener, err = net.Listen("tcp", ":8080")
		if err != nil {
			panic(err)
		}
	}

	server := &http.Server{Handler: n}
	server.Serve(listener)
}
