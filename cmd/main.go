package main

import (
	"go-contacts-api/config"
	"go-contacts-api/controller"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("conf")
	return viper.ReadInConfig()
}

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error loading env variables: %s", err)
	}

	db := config.InitDB(config.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetInt("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLmode:  viper.GetString("db.sslmode"),
	})

	h := controller.NewHandler(db)

	s := &http.Server{
		Addr:    ":8000",
		Handler: h.InitRoutes(),
	}
	s.ListenAndServe()
}
