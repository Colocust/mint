package main

import (
	"encoding/json"
	"log"
	"mint/api"
	"mint/config"
	"mint/route"
	"net"
	"strings"
)

func init() {
	go api.Consume()
}

//在消息生产后通知消费者来消费
func main() {
	address := config.Read("host").(string) + ":" + config.Read("port").(string)
	ln, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	for {
		conn, _ := ln.Accept()
		go run(conn)
	}
}

func run(conn net.Conn) {
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
		resp := route.GetInstance().Handle(request)
		b, _ := json.Marshal(resp)

		_, err = conn.Write(b)
		if err != nil {
			log.Println(err.Error())
		}
	}
}
