package main

import (
	"backend.com/m/v2/database"
	"backend.com/m/v2/router"
)

func main() {
	database.Setup()
	r := router.Setup()

	r.Run()
}
