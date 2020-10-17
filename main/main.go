package main

import (
	"mint/config"
	"mint/route"
	"mint/task"
	"net/http"
)

func init() {
	config.Load()
	task.Boot()
}

type (
	App struct {
		Ip     string
		Port   string
		Router *route.Router
	}
)

func NewApp() *App {
	ip := config.Read("ip").(string)
	port := config.Read("port").(string)
	return &App{
		Ip:     ip,
		Port:   port,
		Router: route.NewRouter(),
	}
}

func (app *App) Run() error {
	server := &http.Server{
		Addr:    app.Ip + ":" + app.Port,
		Handler: app.Router,
	}
	err := server.ListenAndServe()
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
