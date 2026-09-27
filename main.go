package main

import (
	"toko-buku-api1/config"
	"toko-buku-api1/routes"
)

func main() {
	config.ConnectionDatabase()
	r := routes.SetupRouter()
	r.Run(":8080")
}
