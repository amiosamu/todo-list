package main

import "github.com/amiosamu/todo-list/internal/app"

const appConfig = "config/config.yml"

func main() {
	app.Run(appConfig)
}
