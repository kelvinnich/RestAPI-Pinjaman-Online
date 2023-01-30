package config

import (
	"fmt"
	"os"
	"pinjaman-online/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

)


func ConnectDB() *gorm.DB{
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic(errEnv)
	}

	dbHost := os.Getenv("DBHOST")
	dbPort := os.Getenv("DBPORT")
	dbUser := os.Getenv("DBUSER")
	dbPassword := os.Getenv("DBPASSWORD")
	dbname := os.Getenv("DBNAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.Master_Customer{}, &model.Master_Document_Customer{}, &model.Master_Loan{}, &model.Transactions_Payment_Loan{}, &model.Master_Payment_History{}, &model.Master_Jobs_customers{})
	if err != nil {
		fmt.Println("Error migrating tables: ", err)
		os.Exit(1)
	}

	return db
}

func CloseDB(db *gorm.DB){
	dbSql, err := db.DB()
	if err != nil {
		panic(err)
	}
	dbSql.Close()
}
