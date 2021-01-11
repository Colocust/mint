package main

import (
	"encoding/json"
	"errors"
	"log"
	"mint/server/http/server"
	"mint/server/route"
	"mint/server/util/config"
	"net"
	"os"
	"strings"
)

var (
	requireConfig = []string{"host", "port", "retry_max_time"}
)

type (
	App struct {
		Address string
		Router  *route.Router
	}
)

func main() {
	//加载配置
	args := os.Args
	if err := loadConfig(args); err != nil {
		panic(err)
	}
	//检查必填配置
	if err := checkConfig(); err != nil {
		panic(err)
	}

	app := newApp()

	//注册路由
	route.Register(app.Router)

	ln, err := net.Listen("tcp", app.Address)
	if err != nil {
		panic(err)
	}
	log.Println("server started successfully")
	for {
		conn, _ := ln.Accept()
		go boot(conn)
	}
}

func newApp() *App {
	app := &App{
		Address: config.Get("host").(string) + ":" + config.Get("port").(string),
		Router:  route.GetInstance(),
	}
	return app
}

func boot(conn net.Conn) {
	defer func() {
		_ = conn.Close()
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
		conn.Write(buf)
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
			return errors.New("the" + key + "configuration item was not found")
		}
	}

	return nil
}
