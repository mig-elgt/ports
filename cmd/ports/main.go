package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"ports/handler"
	"ports/memo"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	addr := flag.Int("addr", 8080, "server address")
	flag.Parse()
	storage := memo.New()
	defer storage.Close()

	handler := handler.New(storage)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", *addr),
		Handler: handler,
	}
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("listen: %s\n", err)
		}
	}()

	logrus.Printf("Server listeing on %v", *addr)
	<-done
	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
}
