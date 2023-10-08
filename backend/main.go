package main

import (
	"backend/api"
	"backend/database"
)

func main() {
	database.InitDatabase()
	api.InitAPI()

}
