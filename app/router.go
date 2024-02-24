package app

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/ujjwal8007/controller"
	"github.com/ujjwal8007/database"
	"github.com/ujjwal8007/services"
	"github.com/ujjwal8007/store"
)

func setUpRoutes(router *gin.Engine, postgresqlGormDB database.DB) {
	lruCacheStore := store.NewLRUCacheStore(postgresqlGormDB)
	lruCacheService := services.NewLRUCache(lruCacheStore, 3)
	lruCacheController := controller.NewLruCacheController(lruCacheService)
	ctx := context.Background()
	lruCacheStore.StartExpiredKeysDeletion(ctx)
	lruCacheRoutes(router, lruCacheController)
}

func lruCacheRoutes(router *gin.Engine, bureauHandler *controller.LruCacheController) {
	router.POST("/set", bureauHandler.SetKeyHandler)
	router.POST("/get", bureauHandler.GetKeyHandler)
}
