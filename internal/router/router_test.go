package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type mockHandler struct{}

func (m *mockHandler) GetCountries(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func (m *mockHandler) GetCategories(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func (m *mockHandler) GetChannels(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func (m *mockHandler) GetStream(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ok": true})
}