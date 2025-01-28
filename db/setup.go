package db

import (
	"fmt"
	"log"
	"os"
	"salimon/nexus/types"
	"time"

	_ "github.com/golang-migrate/migrate/v4/source/file" // Import file driver
	"github.com/google/uuid"
	_ "github.com/lib/pq" // Import PostgreSQL driver
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// setup database connection
func SetupDatabase() {
	DB = initGormConnection()
	DB.AutoMigrate(types.User{})
	DB.AutoMigrate(types.Verification{})
	DB.AutoMigrate(types.Entity{})
	insertE2EEntity()
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

func insertE2EEntity() {
	entity := types.Entity{
		Id:          uuid.New(),
		Name:        "e2e",
		Description: "e2e testing entity for nexus operations",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err := InsertEntity(&entity)
	if err != nil {
		fmt.Println(err.Error())
	}
}
