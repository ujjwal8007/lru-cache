package app

import (
	"github.com/gin-gonic/gin"
	"github.com/ujjwal8007/controller"
	"github.com/ujjwal8007/database"
	"github.com/ujjwal8007/services"
	"github.com/ujjwal8007/store"
)

func setUpRoutes(router *gin.Engine, postgresqlGormDB database.DB) {
	lruCacheStore := store.NewLRUCacheStore(postgresqlGormDB)
	lruCacheService := services.NewLRUCache(lruCacheStore, 1024)
	lruCacheController := controller.NewLruCacheController(lruCacheService)
	lruCacheRoutes(router, lruCacheController)
}

func lruCacheRoutes(router *gin.Engine, bureauHandler *controller.LruCacheController) {
	router.POST("/set", bureauHandler.SetKeyHandler)
	router.POST("/get", bureauHandler.GetKeyHandler)
}
