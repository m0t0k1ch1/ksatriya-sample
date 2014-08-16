ksatriya-sample
===============

a sample web application, built with [ksatriya](https://github.com/m0t0k1ch1/ksatriya)

## Run

``` sh
$ go get github.com/m0t0k1ch1/ksatriya-sample
$ go build github.com/m0t0k1ch1/ksatriya-sample
$ ./ksatriya-sample
```

If you want to restart the application gracefully, use [go-server-starter-listenr](https://github.com/lestrrat/go-server-starter-listener).

``` sh
$ start_server --port 8080 -- ./ksatriya-sample
```
