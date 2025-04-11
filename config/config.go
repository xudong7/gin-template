package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name string
		Port string
	}

	Front struct {
		Url string
	}

	Database struct {
		Dsn          string
		MaxIdleConns int
		MaxOpenConns int
	}

	Redis struct {
		Host     string
		Port     string
		Password string
		DB       int
	}

	OSS struct {
		Endpoint        string
		AccessKeyId     string
		AccessKeySecret string
		BucketName      string
		URLPrefix       string
		Directory       string
	}
}

var AppConfig *Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file", err)
	}

	AppConfig = &Config{}

	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatal("Unable to decode into struct", err)
	}

	initDB()
	initRedis()
	if err := InitOSS(); err != nil {
		log.Printf("Failed to initialize OSS client: %v", err)
	} else {
		log.Println("OSS client initialized successfully")
	}
}
