package db

import (
	"fmt"
	"log"
	"os"
	"salimon/nexus/types"

	_ "github.com/golang-migrate/migrate/v4/source/file" // Import file driver
	_ "github.com/lib/pq"                                // Import PostgreSQL driver
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// setup database connection
func SetupDatabase() {
	DB = initGormConnection()
	DB.AutoMigrate(types.User{})
	DB.AutoMigrate(types.Verification{})
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

func initGormConnection() *gorm.DB {
	connectionString := generateConnectionString()
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
