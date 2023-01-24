package dto

import "time"

type CreatePembayaranDTO struct {
	Loan_id int `form:"loan_id" json:"loan_id" binding:"required"`
	Loan_Amount int `form:"loan_amount" json:"loan_amount" binding:"required"`
	Monthly_Payments int `form:"monthly_payment" json:"monthly_payment" binding:"required"`
	Payment_Date time.Time `form:"payment_date" json:"payment_date" binding:"required"`
}

type UpdatePembayaranDTO struct {
	Id int `form:"id" json:"id"`
	Loan_id int `form:"loan_id" json:"loan_id" binding:"required"`
	Loan_Amount int `form:"loan_amount" json:"loan_amount" binding:"required"`
	Monthly_Payments int `form:"monthly_payment" json:"monthly_payment" binding:"required"`
	Payment_Date time.Time `form:"payment_date" json:"payment_date" binding:"required"`
}

