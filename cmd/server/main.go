package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/voltgizerz/iptv-browser/internal/handler"
	"github.com/voltgizerz/iptv-browser/internal/repository"
	"github.com/voltgizerz/iptv-browser/internal/service"
)

const serverAddr = ":8080"

func main() {

	repo := repository.NewIPTVRepository()

	svc := service.NewIPTVService(repo)

	h := handler.NewIPTVHandler(svc)

	r := gin.Default()

	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.File("templates/index.html")
	})

	api := r.Group("/api")
	{
		api.GET("/countries", h.GetCountries)
		api.GET("/categories", h.GetCategories)
		api.GET("/channels", h.GetChannels)
		api.GET("/stream/:id", h.GetStream)
	}

	server := &http.Server{
		Addr:    serverAddr,
		Handler: r,
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	go func() {
		log.Println("server listening on http://localhost" + serverAddr)

		if err := server.ListenAndServe(); err != nil &&
			!errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server failed: %v", err)
		}
	}()

	<-ctx.Done()
	stop()

	shutdownCtx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	log.Println("shutting down server...")

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}

	log.Println("server stopped")
}
