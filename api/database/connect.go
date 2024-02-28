package database

import (
	"fmt"
	"github.com/apudiu/alfurqan/internal/model"
	"log"
	"strconv"

	"github.com/apudiu/alfurqan/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB database variable
var DB *gorm.DB

// ConnectDB connect to db
func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("Idiot")
	}

	// Connection URL to connect to Postgres Database
	// user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_HOST"),
		port,
		config.Config("DB_NAME"),
	)
	// Connect to the DB and initialize the DB variable
	DB, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	// auto migrate
	mgErr := DB.AutoMigrate(
		&model.User{},
		&model.Media{},
		&model.Surah{},
		&model.Ayat{},
		&model.Accent{},
		&model.Language{},
	)
	if mgErr != nil {
		fmt.Println("Auto migration err", mgErr.Error())
		return
	}

	fmt.Println("Database Migrated")
}
