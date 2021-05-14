package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Order represents order
type Order struct {
	ID              uint32        `bson:"order_id" json:"order_id"`
	CustomerID      uint32        `bson:"customer_id" json:"customer_id"`
	Customer        *Customer     `bson:"customer" json:"customer"`
	ProviderID      uint32        `bson:"provider_id" json:"provider_id"`
	Provider        *Provider     `bson:"provider" json:"provider"`
	SourceAddress   string        `bson:"source_address" json:"source_address"`
	DeliveryAddress string        `bson:"delivery_address" json:"delivery_address"`
	OrderDetails    []OrderDetail `bson:"order_details" json:"order_details"`
	TotalPrice      float32       `bson:"total_price" json:"total_price"`
	Status          string        `bson:"status" json:"status"`
}

// OrderDetail will detail the order
type OrderDetail struct {
	MenuID     uint32  `json:"menu_id" bson:"menu_id"`
	Menu       *Menu   `json:"menu" bson:"menu"`
	Quantity   uint32  `json:"quantity" bson:"quantity"`
	TotalPrice float32 `json:"total_price" bson:"total_price"`
}

// MenuDetail for body
type MenuDetail struct {
	MenuID   uint32 `json:"menu_id"`
	Quantity uint32 `json:"quantity"`
}

// OrderBody body from json
type OrderBody struct {
	CustomerID uint32       `json:"customer_id"`
	ProviderID uint32       `json:"provider_id"`
	Menus      []MenuDetail `json:"menus"`
}

// OrderRepository reprresents repo functions for order
type OrderRepository interface {
	Store(order *Order) (primitive.ObjectID, error)
	GetAll() ([]Order, error)
	GetByID(id uint32) (*Order, error)
	GetByOID(oid primitive.ObjectID) (*Order, error)
	UpdateArbitrary(id uint32, key string, value interface{}) error
	GetByProviderID(ProviderID uint32) ([]Order, error)
	GetByCustomerID(CustomerID uint32) ([]Order, error)
	GetByCustomerIDAndStatus(CustomerID uint32, status string) ([]Order, error)
	GetByProviderIDAndStatus(ProviderID uint32, status string) ([]Order, error)
	GetByStatus(status string) ([]Order, error)
}

// OrderUsecase usecase for order
type OrderUsecase interface {
	CreateOrder(order OrderBody) (uint32, error)
	AcceptOrder(id uint32) error
	GetAll() ([]Order, error)
	GetByID(id uint32) (*Order, error)
	GetByProviderID(ProviderID uint32) ([]Order, error)
	GetByCustomerID(CustomerID uint32) ([]Order, error)
	GetByCustomerIDAndStatus(CustomerID uint32, status string) ([]Order, error)
	GetByProviderIDAndStatus(ProviderID uint32, status string) ([]Order, error)
	GetByStatus(status string) ([]Order, error)
}
