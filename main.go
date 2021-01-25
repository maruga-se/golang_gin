package main

import (
	"gin_sample/db"
	"gin_sample/server"
)

func main() {
	db.Init()
	server.Init()
}
