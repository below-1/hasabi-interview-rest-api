package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configs struct {
	PgUri  string
	AppEnv string
}

var configs *Configs

func GetConfig() *Configs {
	if configs == nil {
		configs = new(Configs)
		godotenv.Load()
		pg_conn_string, exists := os.LookupEnv("PG_URI")
		if !exists {
			log.Fatal("PG_URI not defined")
			os.Exit(1)
		}
		app_env, exists := os.LookupEnv("APP_ENV")
		if !exists {
			log.Fatal("APP_ENV not defined")
			os.Exit(1)
		}
		configs.PgUri = pg_conn_string
		configs.AppEnv = app_env
	}
	return configs
}
