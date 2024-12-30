package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // Import file driver
	_ "github.com/lib/pq"                                // Import PostgreSQL driver
)

var DB *sql.DB

// setup database connection
func SetupDatabase() {
	connectToDb()
	migrateToDb()
}

// generate connection string from  environment variables
func generateConnectionString() string {
	host := os.Getenv("PGSQL_HOST")
	port := os.Getenv("PGSQL_PORT")
	dbname := os.Getenv("PGSQL_DBNAME")
	username := os.Getenv("PGSQL_USERNAME")
	password := os.Getenv("PGSQL_PASSWORD")
	connectionStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, dbname)
	return connectionStr
}

// connects to db, if fails, terminates the service
func connectToDb() {
	connectionStr := generateConnectionString()
	var err error
	DB, err = sql.Open("postgres", connectionStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected to database")
}

func migrateToDb() {
	driver, err := postgres.WithInstance(DB, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
	log.Println("Migration applied successfully!")
}
