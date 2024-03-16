package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("สวัสดีครับผม")
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`world`))
	})
	// Nothing
	http.HandleFunc("/greeting", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`earth`))
	})
	server := http.Server{
		Addr:    ":2567",
		Handler: mux,
	}
	go func() {
		log.Fatal(server.ListenAndServe())
	}()
	fmt.Println("server starting at :2567")
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	<-shutdown
	fmt.Println("shutting down...")
	if err := server.Shutdown(context.Background()); err != nil {
		fmt.Println("shutdown err:", err)
	}
	fmt.Println("bye bye")
}
