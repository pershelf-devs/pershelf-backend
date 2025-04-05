package main

import (
	"github.com/pershelf/pershelf/cmd"
	"github.com/pershelf/pershelf/cmd/database"
)

func main() {
	database.ConnectMongoDB()
	cmd.Run()
}
