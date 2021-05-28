package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func indexGorillaHandlers(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing index handler")
	fmt.Fprintf(w, "Welcome!")
}

func aboutGorillaHandlers(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing about handler")
	fmt.Fprintf(w, "Go Middleware")
}

func iconHandlerGorillaHandlers(w http.ResponseWriter, r *http.Request) {
}

func main() {
	http.HandleFunc("/favicon.ico", iconHandlerGorillaHandlers)
	indexHandler := http.HandlerFunc(indexGorillaHandlers)
	aboutHandler := http.HandlerFunc(aboutGorillaHandlers)
	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	http.Handle("/", handlers.LoggingHandler(logFile, handlers.CompressHandler(indexHandler)))

	http.Handle("/about", handlers.LoggingHandler(logFile, handlers.CompressHandler(
		aboutHandler)))
	server := &http.Server{
		Addr: ":9000",
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
