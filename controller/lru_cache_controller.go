package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ujjwal8007/models"
	"github.com/ujjwal8007/services"
)

type LruCacheController struct {
	service services.LRUCacheService
}

func NewLruCacheController(s services.LRUCacheService) *LruCacheController {
	return &LruCacheController{service: s}
}

func (l *LruCacheController) SetKeyHandler(c *gin.Context) {
	var setKey models.SetKeyRequest
	if err := c.ShouldBindJSON(&setKey); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := l.service.Put(c, setKey); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (l *LruCacheController) GetKeyHandler(c *gin.Context) {
	var getKeyReq models.GetKeyRequest
	if err := c.ShouldBindJSON(&getKeyReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	value, err := l.service.Get(c, getKeyReq.Key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "key not found or the key is been expired"})
		return
	}

	c.JSON(http.StatusOK, value)
}
