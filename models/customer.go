package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Customer for customer table
type Customer struct {
	CustomerID uint32 `json:"customer_id" bson:"customer_id"`
	Name       string `json:"customer_name" bson:"customer_name"`
	Address    string `json:"customer_address" bson:"customer_address"`
}

// CustomerBody body for Customer
type CustomerBody struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

// CustomerRepository for repo
type CustomerRepository interface {
	Store(customer *Customer) (primitive.ObjectID, error)
	GetAll() ([]Customer, error)
	GetByID(id uint32) (*Customer, error)
	GetByOID(oid primitive.ObjectID) (*Customer, error)
	UpdateArbitrary(id uint32, key string, value interface{}) error
}

// CustomerUsecase for usecase
type CustomerUsecase interface {
	CreateCustomer(customer CustomerBody) (uint32, error)
	GetAll() ([]Customer, error)
	GetByID(id uint32) (*Customer, error)
}
