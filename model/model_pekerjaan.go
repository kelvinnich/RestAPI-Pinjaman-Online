package model


type Master_Jobs_customers struct {
	ID         int `gorm:"primary_key;auto_increment" json:"id"`
	Customer_Id  int ` gorm:"not null " json:"-"`
	Company_Address string `gorm:"type:varchar(255)" json:"company_address"`
	Payday_Date string `gorm:"type:varchar(50)" json:"payDay_date"`
	Job_Position string `gorm:"type:varchar(255)" json:"job_position"`
}