package repository

import (
	"log"
	"pinjaman-online/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)


type NasabahRepository interface{
	CreateNasabah(nasabah model.Master_Customer)(model.Master_Customer, error)
	UpdateNasabah(id int, nasabah model.Master_Customer) (model.Master_Customer,error)
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	IsDuplicateNIk(noKtp string) (tx *gorm.DB)
	FindByNIK(NIK string) model.Master_Customer
	ProfileNasabah(nasabahID string) model.Master_Customer
}

type nasabahConnection struct{
	db *gorm.DB
}

func NewNasabahRepository(db *gorm.DB) NasabahRepository{
	return &nasabahConnection{
		db: db,
	}
}

func(db *nasabahConnection)CreateNasabah(nasabah model.Master_Customer) (model.Master_Customer, error){
	nasabah.Password = HashPassword([]byte(nasabah.Password))
	if err := db.db.Create(nasabah).Error; err != nil{
		panic(err)
	}
	return nasabah,nil
}

func(db *nasabahConnection)UpdateNasabah(id int,nasabah model.Master_Customer)( model.Master_Customer, error){
	if nasabah.Password != ""{
		nasabah.Password = HashPassword([]byte(nasabah.Password))
	}else {
		var nasabahTemp model.Master_Customer
		db.db.Find(&nasabahTemp, nasabah.Id)
		nasabah.Password = nasabahTemp.Password
	}
	if err := db.db.Model(&model.Master_Customer{}).Where("id = $1", id).Updates(nasabah).Error; err != nil {
		panic(err)
	}
	return nasabah,nil
}

func(db *nasabahConnection)VerifyCredential(email string, password string) interface{}{
	var nasabah model.Master_Customer
	res := db.db.Where("email = $1", email).Take(&nasabah)
	if res.Error == nil {
		return res
	}
	return nil
}

func(db *nasabahConnection)IsDuplicateEmail(email string) (dB *gorm.DB){
	var nasabah model.Master_Customer
	return db.db.Where("email = $1", email).Take(&nasabah)
}

func(db *nasabahConnection)IsDuplicateNIk(noKtp string) (dB *gorm.DB){
	var nasabah model.Master_Customer
	return db.db.Where("noKtp = $1", noKtp).Take(&nasabah)
}

func(db *nasabahConnection)FindByNIK(nik string) model.Master_Customer{
	var nasabah model.Master_Customer
	db.db.Where("no_ktp = $1", nik).Take(&nasabah)
	return nasabah
}

func(db *nasabahConnection)ProfileNasabah(nasabahID string) model.Master_Customer{
	var nasabah model.Master_Customer
	db.db.Find(&nasabah, nasabahID)
	return nasabah
}


func HashPassword(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}