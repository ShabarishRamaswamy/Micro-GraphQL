package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ShabarishRamaswamy/Micro-GraphQL/src/router"
)

var PORT string = ":8000"

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	r := router.GetNewRouter()

	go func() {
		log.Fatal(http.ListenAndServe(PORT, r))
	}()

	<-quit
}
