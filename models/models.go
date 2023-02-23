package models

import "time"

type CustomerCompany struct {
	ID          uint   `json:"company_id" gorm:"primaryKey"`
	CompanyName string `json:"company_name" gorm:"not null"`
}

type Customer struct {
	UserID          string          `json:"user_id" gorm:"primaryKey"`
	Login           string          `json:"login" gorm:"not null"`
	Password        string          `json:"password" gorm:"not null"`
	Name            string          `json:"name" gorm:"not null"`
	CompanyID       uint            `json:"company_id" gorm:"not null"`
	CustomerCompany CustomerCompany `gorm:"foreignKey:CompanyID"`
	CreditCards     string          `json:"credit_cards" gorm:"not null"`
}

type Order struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time `json:"created_at" gorm:"not null"`
	OrderName  string    `json:"order_name" gorm:"not null"`
	CustomerID string    `json:"customer_id" gorm:"not null"`
	Customer   Customer  `gorm:"foreignKey:CustomerID"`
}

type OrderItem struct {
	ID           uint    `json:"id" gorm:"primaryKey"`
	OrderID      uint    `json:"order_id" gorm:"not null"`
	Order        Order   `gorm:"foreignKey:OrderID"`
	PricePerUnit float64 `json:"price_per_unit" gorm:"not null"`
	Quantity     int     `json:"quantity" gorm:"not null"`
	Product      string  `json:"product" gorm:"not null"`
}

type Delivery struct {
	ID                int       `json:"id" gorm:"primaryKey"`
	OrderItemID       uint      `json:"order_item_id" gorm:"not null"`
	OrderItem         OrderItem `gorm:"foreignKey:OrderItemID"`
	DeliveredQuantity int       `json:"delivered_quantity" gorm:"not null"`
}
