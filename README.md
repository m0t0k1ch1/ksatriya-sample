ksatriya-sample
===============

a sample web application built with [ksatriya](https://github.com/m0t0k1ch1/ksatriya)

## Run

``` sh
$ go get github.com/m0t0k1ch1/ksatriya-sample
$ go build github.com/m0t0k1ch1/ksatriya-sample
$ ./ksatriya-sample
```

Using [Server::Starter](http://search.cpan.org/~kazuho/Server-Starter-0.17/lib/Server/Starter.pm), you can restart the application gracefully.

``` sh
$ start_server --port 8080 -- ./ksatriya-sample
```
