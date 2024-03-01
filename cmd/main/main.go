package main

import (
	"github.com/toastsandwich/CRUD-echo/pkg/config"
)

func init() {
	config.BootUp()
}

func main() {
	app := config.App
	app.Logger.Fatal(app.Start(":1234"))
}
