package entity

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name       string `valid:"required"`
	Email      string
	CustomerID string // รหัสลูกค้าขึ8นต้นด้วย L หรือ M หรือ H แล้วตามด้วยตัวเลขจํานวน 7 ตัว
}
