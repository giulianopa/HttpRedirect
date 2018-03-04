# Http Redirect in Golang
This is meant mainly as a Go playground. The idea is to run a redirect service
that logs the IP address of the remote host.

## Build and run
To build:
```
$ cd HttpRedirect/
$ go get github.com/gorilla/mux
$ go build redirect.go
```

To run:
```
$ ./redirect
```

Then, from a browser: `localhost:8080/r/<URL_to_redirect_to>`
To see the log: `localhost:8080/log`


## Useful links
* Effective Go: https://golang.org/doc/effective_go.html
* Actual Go playground: https://play.golang.org/
* If you need something more complex, like a customizable HTTP proxy:
https://github.com/elazarl/goproxy
