package main

import (
	"pocketmail-api/src/mail"

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
	s := mail.MailServer{ App: app }

	// Start the SMTP server in a separate goroutine.
	go s.Serve()

	// Start pocketbase.
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
