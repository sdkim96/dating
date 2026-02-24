package main

import (
	"flag"
	"fmt"

	"github.com/sdkim96/dating/internal/server"
)

func main() {
	mode := flag.String("mode", "stdio", "transport mode: stdio or http")
	addr := flag.String("addr", ":8080", "http listen address")
	flag.Parse()

	var err error
	switch *mode {
	case "stdio":
		s := server.NewStdio()
		err = s.Serve()
	case "http":
		fmt.Printf("listening on %s\n", *addr)
		s := server.NewHTTP(*addr)
		err = s.Serve()
	default:
		err = fmt.Errorf("unknown mode: %s", *mode)
	}

	if err != nil {
		fmt.Printf("server error: %v\n", err)
	}
}
