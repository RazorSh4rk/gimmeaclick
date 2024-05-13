package main

import (
	"razorsh4rk.github.io/urlshorten/api"
	"razorsh4rk.github.io/urlshorten/db"
)

func main() {
	db.Connect()
	api.Setup()
	api.Run()
}
