package model

import "time"

type Transactions_Payment_Loan struct {
	ID int` gorm:"primary_key;column:id;type:serial" json:"id"`
	Loan_id int `gorm:"not null" json:"-"`
	Loan Master_Loan `gorm:"foreignKey:Loan_id;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"loan_id"`
	Loan_Amount int `gorm:"type:integer" json:"loan_amount"`
	Monthly_Payments int `gorm:"type:integer" json:"monthly_payment"`
	Payment_Status bool `gorm:"type:boolean" json:"payment_status"`
	Payment_Date time.Time `gorm:"type:timestamp" json:"payment_date"`
	}