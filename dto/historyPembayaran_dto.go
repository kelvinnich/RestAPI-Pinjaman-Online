package dto

import "time"


type CreateHistoryPembayaranDTO struct {
	Loan_id int `json:"loan_id" form:"loan_id" binding:"required"`
	payment_id int `json:"payment_id" form:"payment_id" binding:"required"`
	Date time.Time `json:"date" form:"date" binding:"required"`
}

type UpdateHistoryPembayaranDTO struct {
	Id uint64 `json:"id" form:"id"`
	Loan_id int `json:"loan_id" form:"loan_id" binding:"required"`
	payment_id int `json:"payment_id" form:"payment_id" binding:"required"`
	Date time.Time `json:"date" form:"date" binding:"required"`
}
