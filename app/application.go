package app

import (
	"net/http"
	"time"
	"fmt"
)

type Config struct{
	Addr string
}

type Application struct {
	Config Config
}
 
func (app *Application) Run() error{
	server := &http.Server{
		Addr: app.Config.Addr,
		Handler: nil,// TODO : add chi routes
		
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 15 * time.Second,
	} 
	fmt.Println("Server Running at port ", app.Config.Addr)

	return server.ListenAndServe()

}
