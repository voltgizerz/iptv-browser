package router

import (
	"github.com/gin-gonic/gin"
	"github.com/voltgizerz/iptv-browser/internal/handler"
	"github.com/voltgizerz/iptv-browser/internal/middleware"
)

type Router struct {
	handler *handler.IPTVHandler
}

func NewRouter(h *handler.IPTVHandler) *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())

	// static
	r.Static("/static", "./static")
	r.StaticFile("/sw.js", "./static/sw.js")

	r.GET("/manifest.webmanifest", func(c *gin.Context) {
		c.Header("Content-Type", "application/manifest+json")
		c.File("static/manifest.webmanifest")
	})

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

	return r
}
