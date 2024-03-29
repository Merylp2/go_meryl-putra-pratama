package main

import (
	"praktikum/config"
	"praktikum/middlewares"
	"praktikum/routes"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// create a new echo instance
	config.InitDB()
	e := routes.New()

	middlewares.LogMiddlewares(e)
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))

}
