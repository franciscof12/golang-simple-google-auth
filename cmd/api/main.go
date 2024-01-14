package main

import (
	"autho-go-microservice/internal/oauth"
	"autho-go-microservice/internal/server"
	"fmt"
)

func main() {
	oauth.NewAuth()
	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
