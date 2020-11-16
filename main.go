package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mint/config"
	"net"
)

//开启tcp服务
//接收处理
//暂时一个处理开一个协程
func main() {
	address := config.Read("host").(string) + ":" + config.Read("port").(string)
	ln, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	for {
		conn, e := ln.Accept()
		if e != nil {
			log.Printf("accept err: %v", e)
			return
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	body := make([]byte, 1024)
	for {
		length, err := conn.Read(body)
		if err != nil {
			break
		}
		fmt.Println(string(body))
		req := make(map[string]interface{})
		json.Unmarshal(body[:length], &req)
		fmt.Println(req["host"])
	}
	fmt.Println("断开连接")
}
