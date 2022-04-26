package example_app_db_connection

import (
	"context"
	"fmt"
	"os"

	"github.com/teejays/clog"
	"github.com/teejays/goku-util/client/db"
)

// InitDatabaseConnections initializes all of the service DB connections for this app
func InitDatabaseConnections(ctx context.Context) error {
	var err error

	// Ensure neccessary ENV variables are set
	if os.Getenv("DATABASE_HOST") == "" {
		return fmt.Errorf("ENV variable DATABASE_HOST not set")
	}
	if os.Getenv("POSTGRES_USERNAME") == "" {
		return fmt.Errorf("ENV variable POSTGRES_USERNAME not set")
	}
	if os.Getenv("PGPASSWORD") == "" {
		return fmt.Errorf("ENV variable PGPASSWORD not set")
	}
	clog.Warnf("Initializing database connection to database %s", "pharmacy")
	err = db.InitDatabase(ctx, db.Options{
		Host:     os.Getenv("DATABASE_HOST"),
		Port:     db.DEFAULT_POSTGRES_PORT,
		Database: "pharmacy",
		User:     os.Getenv("POSTGRES_USERNAME"),
		SSLMode:  "disable",
	})
	if err != nil {
		return fmt.Errorf("Initalizing database `pharmacy`: %w", err)
	}
	clog.Warnf("Initializing database connection to database %s", "users")
	err = db.InitDatabase(ctx, db.Options{
		Host:     os.Getenv("DATABASE_HOST"),
		Port:     db.DEFAULT_POSTGRES_PORT,
		Database: "users",
		User:     os.Getenv("POSTGRES_USERNAME"),
		SSLMode:  "disable",
	})
	if err != nil {
		return fmt.Errorf("Initalizing database `users`: %w", err)
	}
	err = db.CheckConnection("pharmacy")
	if err != nil {
		return fmt.Errorf("Failed to verify connection to database `pharmacy`: %w", err)
	}
	err = db.CheckConnection("users")
	if err != nil {
		return fmt.Errorf("Failed to verify connection to database `users`: %w", err)
	}

	return nil
}
