package main

import (
	"pocketmail-api/src/mail"

	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"

	_ "pocketmail-api/migrations"
)

var IsDev = strings.HasPrefix(os.Args[0], os.TempDir())

func main() {
	_ = godotenv.Load(".env")

    // loosely check if it was executed using "go run"
    isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())
	
	// Create a new pocketbase instance.
	app := pocketbase.New()
	s := mail.MailServer{ App: app }

	// register the migrate command plugin
    migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
        // enable auto creation of migration files when making collection changes in the Dashboard
        // (the isGoRun check is to enable it only during development)
        Automigrate: isGoRun,
    })

	// Start the SMTP server in a separate goroutine.
	go s.Serve()

	// Start pocketbase.
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
