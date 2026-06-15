package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	ProdEnabled bool   `mapstructure:"prod_enabled"`
	UploadUrl   string `mapstructure:"upload_url"`
	HttpPort    int    `mapstructure:"http_port"`

	DbConnection string `mapstructure:"db_connection"`
	DbHost       string `mapstructure:"db_host"`
	DbPort       string `mapstructure:"db_port"`
	DbDatabase   string `mapstructure:"db_database"`
	DbUsername   string `mapstructure:"db_username"`
	DbPassword   string `mapstructure:"db_password"`
}

var App Config
var RequestLocation *time.Location

func LoadConfig() {
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
	if err := viper.Unmarshal(&App); err != nil {
		log.Fatalln(err)
	}

	if App.HttpPort == 0 {
		App.HttpPort = 8000
	}

	// // init the loc
	RequestLocation, _ = time.LoadLocation("Asia/Ashgabat")
	UTCLoc, _ := time.LoadLocation("UTC")
	time.Local = UTCLoc
	// // set timezone,
	// os.Setenv("TZ", "Asia/Ashgabat")
}
