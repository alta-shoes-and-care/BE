package configs

import (
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

type AppConfig struct {
	PORT        int16
	DB_DRIVER   string
	DB_NAME     string
	DB_PORT     int16
	DB_HOST     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_LOC      string
	S3_REGION   string
	S3_KEY      string
	S3_SECRET   string
}

var synchronizer = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig(isTest bool) *AppConfig {
	synchronizer.Lock()
	defer synchronizer.Unlock()

	if appConfig == nil {
		appConfig = initConfig(isTest)
	}
	return appConfig
}

func initConfig(isTest bool) *AppConfig {
	if err := godotenv.Load(".env"); err != nil {
		log.Info("tidak dapat memuat env file", err)
	}

	var defaultAppConfig AppConfig

	getEnv(&defaultAppConfig, isTest)

	log.Info("connected to:\n", defaultAppConfig)

	return &defaultAppConfig
}

func getEnv(appConfig *AppConfig, isTest bool) {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Warn(err)
	}

	db_port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Warn(err)
	}

	appConfig.PORT = int16(port)
	appConfig.DB_DRIVER = os.Getenv("DB_DRIVER")
	appConfig.DB_NAME = os.Getenv("DB_NAME")
	appConfig.DB_PORT = int16(db_port)
	appConfig.DB_HOST = os.Getenv("DB_HOST")
	appConfig.DB_USERNAME = os.Getenv("DB_USERNAME")
	appConfig.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	appConfig.DB_LOC = os.Getenv("DB_LOC")
	appConfig.S3_REGION = os.Getenv("S3_REGION")
	appConfig.S3_KEY = os.Getenv("S3_KEY")
	appConfig.S3_SECRET = os.Getenv("S3_SECRET")

	if isTest {
		appConfig.DB_NAME = "immersive6"
	}
}
