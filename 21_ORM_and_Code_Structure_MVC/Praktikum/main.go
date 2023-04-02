package main

import (
	"test/21_ORM_and_Code_Structure_MVC/Praktikum/config"
	"test/21_ORM_and_Code_Structure_MVC/Praktikum/routes"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// create a new echo instance
	config.InitDB()
	e := routes.New()

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))

}
