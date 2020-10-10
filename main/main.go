package main

import (
	"net/http"
	"tinyQ/config"
	"tinyQ/route"
)

func init() {
	config.Load()
}

type (
	App struct {
		Ip     string
		Port   string
		Router *route.Router
		Server *http.Server
	}
)

func NewApp() *App {
	return &App{
		Ip:     "127.0.0.1",
		Port:   "8899",
		Router: route.NewRouter(),
	}
}

func (app *App) Run() error {
	if app.Server == nil {
		app.Server = &http.Server{
			Addr:    app.Ip + ":" + app.Port,
			Handler: app.Router,
		}
	}
	err := app.Server.ListenAndServe()
	return err
}

func main() {
	app := NewApp()
	route.Register(app.Router)

	err := app.Run()
	if err != nil {
		panic(err)
	}
}
