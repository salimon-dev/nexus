package db

import (
	"crypto/md5"
	"encoding/hex"
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
	DB.AutoMigrate(types.Invitation{})
	insertKeymaker()
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

	entity, err := FindEntity("name = 'e2e'")
	if err != nil {
		fmt.Println(err)
		return
	}
	if entity != nil {
		fmt.Println("e2e entity exists")
		return
	}
	e2eEntity := types.Entity{
		Id:          uuid.New(),
		Name:        "e2e",
		Description: "e2e testing entity for nexus operations",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err = InsertEntity(&e2eEntity)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func insertKeymaker() {
	username := os.Getenv("KEYMAKER_USERNAME")
	passwordPlain := os.Getenv("KEYMAKER_PASSWORD")

	user, err := FindUser("username = ? AND role = ?", username, types.UserRoleKeyMaker)
	if err != nil {
		fmt.Println(err)
		return
	}
	if user != nil {
		fmt.Println("keymaker user exists")
		return
	}

	passwordHash := md5.Sum([]byte(passwordPlain))
	password := hex.EncodeToString(passwordHash[:])

	keymaker := types.User{
		Id:           uuid.New(),
		Username:     username,
		Password:     password,
		Credit:       90000,
		Usage:        0,
		InvitationId: uuid.Nil,
		Role:         types.UserRoleKeyMaker,
		Status:       types.UserStatusActive,
		RegisteredAt: time.Now(),
		UpdatedAt:    time.Now(),
	}

	err = InsertUser(&keymaker)
	if err != nil {
		fmt.Println(err)
	}
}
