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
	

  db.AutoMigrate(&model.Nasabah{}, &model.Dokumen_nasabah{}, &model.Pinjaman{}, &model.Pembayaran{}, &model.HistoryPembayaran{}, &model.Pekerjaan_nasabah{} )


	return db
}

func CloseDB(db *gorm.DB){
	dbSql, err := db.DB()
	if err != nil {
		panic(err)
	}
	dbSql.Close()
}