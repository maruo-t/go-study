package main

import (
	"github.com/maruo-t/go-study/src/todo_app/config"
)

func main() {
	println(config.Config.Port)
	println(config.Config.SQLDriver)
	println(config.Config.DbName)
	println(config.Config.LogFile)
}
