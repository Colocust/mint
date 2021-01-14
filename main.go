package main

import (
	"encoding/json"
	"errors"
	"log"
	"mint/http/server"
	"mint/job/delay"
	"mint/route"
	"mint/util/config"
	"net"
	"os"
	"strings"
)

var (
	requireConfig = []string{"host", "port"}
)

type (
	App struct {
		Address string
		Router  *route.Router
	}
)

func main() {
	args := os.Args
	if err := loadConfig(args); err != nil {
		panic(err)
	}
	if err := checkConfig(); err != nil {
		panic(err)
	}

	app := newApp()
	route.Register(app.Router)

	go jobBoot()

	ln, err := net.Listen("tcp", app.Address)
	if err != nil {
		panic(err)
	}
	log.Println("server started successfully")
	for {
		conn, _ := ln.Accept()
		go serverBoot(conn)
	}
}

func newApp() *App {
	app := &App{
		Address: config.Get("host").(string) + ":" + config.Get("port").(string),
		Router:  route.GetInstance(),
	}
	return app
}

func serverBoot(conn net.Conn) {
	defer func() {
		conn.Close()
		log.Println(conn.RemoteAddr().String() + "断开连接")
	}()
	for {
		body := make([]byte, 1024)
		length, err := conn.Read(body)
		if err != nil {
			break
		}

		request := strings.Split(string(body[:length]), " ")
		resp := &server.Response{}

		route.GetInstance().Handle(request, resp)

		buf, _ := json.Marshal(resp)
		_, err = conn.Write(buf)
	}
}

func loadConfig(args []string) error {
	wd, _ := os.Getwd()
	var name = wd + "/config.json"

	if len(args) == 2 {
		name = args[1]
	}

	if err := config.Load(name); err != nil {
		return err
	}
	return nil
}

func checkConfig() error {
	for _, key := range requireConfig {
		if value := config.Get(key); value == nil {
			err := errors.New("the" + key + "configuration item was not found")
			return err
		}
	}
	return nil
}

func jobBoot() {
	delay.Scan()
}
