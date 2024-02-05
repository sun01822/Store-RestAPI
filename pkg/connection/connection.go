package connection

import (
	"Store_RestAPI/pkg/config"
	"Store_RestAPI/pkg/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func Connect() {
	dbConfig := config.LocalConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.DBUser, dbConfig.DBPass, dbConfig.DBIp, dbConfig.DBName)
	// dsn := "root:password@tcp(localhost:3306)/store?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// print sql query
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("Error connecting to database")
		panic(err)
	}
	db = d
}

// creating new table to store database
func migrate() {
	db.Migrator().AutoMigrate(&models.Product{})
	db.Migrator().AutoMigrate(&models.User{})
	db.Migrator().AutoMigrate(&models.Cart{})
	db.Migrator().AutoMigrate(&models.ProductsInCart{})
}

// calling to connect fucntion to initialize connection
func GetDB() *gorm.DB {
	if db == nil {
		Connect()
	}
	migrate()
	return db
}