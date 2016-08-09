ksatriya-sample
===============

a sample web application built with [ksatriya](https://github.com/m0t0k1ch1/ksatriya)

## Build & Run

``` sh
$ go get github.com/m0t0k1ch1/ksatriya-sample
$ ksatriya-sample --conf=$GOPATH/src/github.com/m0t0k1ch1/ksatriya-sample/config.tml
```

Using [go-server-starter](https://github.com/lestrrat/go-server-starter), you can restart the application gracefully.

``` sh
$ start_server --port 8080 -- ksatriya-sample --conf=$GOPATH/src/github.com/m0t0k1ch1/ksatriya-sample/config.tml
```
