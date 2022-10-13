package config

import (
	"os"
	"strconv"
)

const (
	DEVELOPER = "developer"
 	HOMOLOGATION = "homologation"
 	PRODUCTION = "production"
)

type Config struct {
	SRV_PORT string 
	WEB_UI bool   //publico pois esta tudo maisculo
	Mode string   //privado 
	OpenBrowser bool 

	DBConfig
}

type DBConfig struct  {
	DB_DRIVE string
	DB_HOST string
	DB_PORT string
	DB_USER string
	DB_PASS string
	DB_NAME string
	DB_DSN string
}

func NewConfig(config *Config) Config {
	var conf *Config

	if config == nil || config.SRV_PORT == "" {
		conf = defaultConfig()
	} else {
		conf = config
	}

	SRV_PORT := os.Getenv("SRV_PORT")
	if SRV_PORT != "" {
		conf.SRV_PORT = SRV_PORT
	}

	SRV_WEB_UI := os.Getenv("SRV_WEB_UI")
	if SRV_WEB_UI != "" {
		conf.WEB_UI, _ = strconv.ParseBool(SRV_WEB_UI)
	}

	SRV_OPEN_BROWSER := os.Getenv("SRV_OPEN_BROWSER")
	if SRV_OPEN_BROWSER != "" {
		conf.OpenBrowser, _ =
	}

}

func defaultConfig() *Config{
	default_config := Config{
		SRV_PORT: 	"8080",
		WEB_UI: 	 true,
		OpenBrowser: true,
		Mode: 		 PRODUCTION,
		DBConfig: DBConfig{
			DB_DRIVE: "sqlite3",
			DB_NAME: "db.sqlite3",
		},
	}

	return &default_config
}


