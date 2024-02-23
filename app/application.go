package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/ujjwal8007/config"
	"github.com/ujjwal8007/database"
	"gorm.io/gorm"
)

var router = gin.Default()

// StartApplication Start an application.
func StartApplication() {
	router = gin.Default()
	router.HandleMethodNotAllowed = true

	BootstrapApp()

	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(config.App.Port),
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal()
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal()
	}

}

func BootstrapApp() {
	postgresqlGormDB := Init()

	postgresqlDB := database.NewDB(postgresqlGormDB)

	setUpRoutes(router, postgresqlDB)
}

func Init() *gorm.DB {

	config.LoadConfig()

	postgresqlDB, err := config.ConnectToPostgreSQL()
	if err != nil {
		fmt.Println("error connecting to postgresql database")
		log.Fatal(err.Error())
	}
	return postgresqlDB
}
