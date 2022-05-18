package main

import (
	internal "github.com/brunoshiroma/go-url-expander/internal"
)

func main() {
	var (
		l      = internal.Logger()
		server = internal.NewHttpServer(internal.Config.Port, internal.Config.BindHost)
	)
	l.Infof("Staring server on %s:%d", internal.Config.BindHost, internal.Config.Port)
	server.ListenAndServe()
}
