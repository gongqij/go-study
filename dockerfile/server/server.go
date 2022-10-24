package server

import (
	"context"
	"fmt"
	"k8s-study/core"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type server struct {
	*http.Server
}

func NewServer(handler http.Handler) *server {
	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", core.AppConf.App.Port),
		Handler: handler,
	}

	return &server{
		s,
	}
}

func (s *server) Run() {
	go func() {
		fmt.Printf("Server starting at http://127.0.0.1:%d\n", core.AppConf.App.Port)
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("ListenAndServe err: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	fmt.Println("Shutdown Server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	fmt.Println("Server Exited Properly")
}
