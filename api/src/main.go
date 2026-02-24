package main

import (
	"log"
	"os"
	"strings"
	
	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase"
)

var IsDev = strings.HasPrefix(os.Args[0], os.TempDir())

func main() {
	_ = godotenv.Load(".env")
	
	// Create a new pocketbase instance.
	app := pocketbase.New()

	// Start pocketbase.
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
