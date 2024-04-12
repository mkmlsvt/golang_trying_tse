package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

type httpServer struct {
	port     string
	fiberApp *fiber.App
}

func NewHttpServer(port string, fiberApp *fiber.App) httpServer {
	return httpServer{port: port, fiberApp: fiberApp}
}

func (s httpServer) Start(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-ctx.Done():
			print("http server shutting down ...")
			if err := s.fiberApp.ShutdownWithTimeout(5 * time.Second); err != nil {
				fmt.Printf("http server shutdown failed: %v", err)
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Printf("http server starting on %s", s.port)
		if err := s.fiberApp.Listen(s.port); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf(err.Error())
		}
	}()
}
