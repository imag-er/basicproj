package models

import "time"

type ExchangeRate struct {
	ID           uint      `gorm:"primaryKey" form:"id,required"`
	FromCurrency string    `form:"from_currency,required"`
	ToCurrency   string    `form:"to_currency,required"`
	Rate         float64   `form:"rate,required"`
	Date         time.Time `form:"date,required"`
}
	