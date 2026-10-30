package app

import (
	config "AuthInGo/config/env"
	"AuthInGo/router"
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Addr string
}

type Application struct {
	Config Config
}

// constructor for config
func NewConfig() Config {
	addr := config.GetString("PORT", ":8080")
	return Config{
		Addr: addr,
	}
}

// constructor for Application
func NewApplication(cnf Config) *Application {
	return &Application{
		Config: cnf,
	}
}

func (app *Application) Run() error {
	server := &http.Server{
		Addr:    app.Config.Addr,
		Handler: router.SetupRouter(),

		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	fmt.Println("Server Running at port ", app.Config.Addr)

	return server.ListenAndServe()

}
