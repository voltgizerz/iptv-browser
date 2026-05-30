package main

import (
	"html/template"

	"github.com/gin-gonic/gin"

	"github.com/voltgizerz/iptv-browser/internal/handler"
	"github.com/voltgizerz/iptv-browser/internal/repository"
	"github.com/voltgizerz/iptv-browser/internal/service"
)

func main() {

	repo := repository.NewIPTVRepository()

	svc := service.NewIPTVService(repo)

	h := handler.NewIPTVHandler(svc)

	r := gin.Default()

	r.Static("/static", "./static")

	r.SetHTMLTemplate(
		template.Must(
			template.ParseFiles(
				"templates/index.html",
			),
		),
	)

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	api := r.Group("/api")
	{
		api.GET("/countries", h.GetCountries)
		api.GET("/channels", h.GetChannels)
		api.GET("/stream/:id", h.GetStream)
	}

	r.Run(":8080")
}