package config

import (
	"path/filepath"
	"runtime"

	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	Port int

	ProjectRootPath string
}

// App ..
var App *AppConfig

func LoadConfig() {
	// Get current file full path from runtime
	_, b, _, _ := runtime.Caller(0)

	// Root folder of this project
	projectRootPath := filepath.Join(filepath.Dir(b), "../")

	// load .env file
	err := godotenv.Load(projectRootPath + "/prod.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	loadAppConfig()

}

func RootPath() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), "../")
}

func loadAppConfig() {
	App = &AppConfig{}
	// var s AppConfig
	err := envconfig.Process("app", App)

	// Get current file full path from runtime
	_, b, _, _ := runtime.Caller(0)

	// Root folder of this project
	App.ProjectRootPath = filepath.Join(filepath.Dir(b), "../")

	if err != nil {
		log.Fatal(err.Error())
	}
}
