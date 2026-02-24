package main

import (
	"flag"
	"fmt"

	"github.com/sdkim96/dating/internal/server"
)

func main() {
	mode := flag.String("mode", "stdio", "transport mode: stdio, http, or http-stateful")
	addr := flag.String("addr", ":8080", "http listen address")
	flag.Parse()

	var err error
	switch *mode {
	case "stdio":
		err = server.RunStdio()
	case "http":
		fmt.Printf("listening on %s (stateless)\n", *addr)
		err = server.RunHTTPStateless(*addr)
	case "http-stateful":
		fmt.Printf("listening on %s (stateful)\n", *addr)
		err = server.RunHTTPStateful(*addr)
	default:
		err = fmt.Errorf("unknown mode: %s", *mode)
	}

	if err != nil {
		fmt.Printf("server error: %v\n", err)
	}
}
