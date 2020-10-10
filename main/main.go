package main

import (
	"tinyQ"
	"tinyQ/config"
)


func main() {
	config.Load()
	tinyQ.Boot()
}
