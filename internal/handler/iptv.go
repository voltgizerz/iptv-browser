package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/voltgizerz/iptv-browser/internal/service"
)

type IPTVHandler struct {
	service *service.IPTVService
}

func NewIPTVHandler(
	service *service.IPTVService,
) *IPTVHandler {
	return &IPTVHandler{
		service: service,
	}
}

func (h *IPTVHandler) GetCountries(c *gin.Context) {
	data, err := h.service.GetCountries(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *IPTVHandler) GetChannels(c *gin.Context) {

	data, err := h.service.GetChannels(
		c,
		c.Query("country"),
		c.Query("q"),
	)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}

func (h *IPTVHandler) GetStream(c *gin.Context) {

	url, err := h.service.GetStream(
		c,
		c.Param("id"),
	)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"url": url,
	})
}