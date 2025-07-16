package main

import (
	"log"
	"todo-cli/cmd"

	"github.com/joho/godotenv"
	_ "modernc.org/sqlite"
)

func main() {
	// Env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Cmd
	cmd.Execute()
}
