package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gotamboon/servers"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: <project_cmd> <csv_file>")
		return
	}

	server := servers.NewServer(os.Args[1])
	server.Start()

}
