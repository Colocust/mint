package main

import (
	"encoding/json"
	"net/http"
	"socket/server"
)

const CAPACITY = 100

var chs = make(chan server.Task, CAPACITY)

func init() {
	go server.Read(chs)
}

func main() {
	//http.HandleFunc("/send", run)
	//err := http.ListenAndServe("127.0.0.1:8006", nil)
	//if err != nil {
	//	panic(err)
	//}
}

func run(writer http.ResponseWriter, request *http.Request) {
	//body := make([]byte, request.ContentLength)
	//_, err := request.Body.Read(body)
	//if err != nil {
	//}
	//
	//task := new(server.Task)
	//json.Unmarshal(body, task)
	//
	//go server.Write(chs, *task)
}

