package main

import(
	"github.com/leobrines/go-apirest-example/webserver"
)

func main () {
	webserver := new(webserver.Webserver)
	webserver.Configure()
	webserver.Start()
}

