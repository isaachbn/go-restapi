package db

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
	"os"
	"restapi/domain/books"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
		panic(err)
	}
}

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(openDriver(), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting database %v", err)
		panic(err)
	}

	// Migrate the schema
	db.AutoMigrate(&books.Book{})

	return db
}

func openDriver() gorm.Dialector {
	driver := os.Getenv("driver")
	dns := os.Getenv("dns")

	switch driver {
	case "mysql":
		return mysql.Open(dns)
	case "postgres":
		return postgres.Open(dns)
	case "sqlite":
		return sqlite.Open(dns)
	case "sqlserver":
		return sqlserver.Open(dns)
	default:
		panic("Database drive is not supported")
	}
}
