package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Provider represents provider data
type Provider struct {
	ProviderID  uint32 `json:"provider_id" bson:"provider_id"`
	Name        string `json:"provider_name" bson:"provider_name"`
	Address     string `json:"provider_address" bson:"provider_address"`
	PhoneNumber string `json:"provider_phone" bson:"provider_phone"`
	OpenHour    int64  `json:"open_hour" bson:"open_hour"`
	ClosedHour  int64  `json:"closed_hour" bson:"closed_hour"`
}

// ProviderBody body for buyer
type ProviderBody struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	OpenHour    int64  `json:"open_hour"`
	ClosedHour  int64  `json:"closed_hour"`
}

// ProviderRepository represents repo functions for Provider
type ProviderRepository interface {
	Store(provider *Provider) (primitive.ObjectID, error)
	GetAll() ([]Provider, error)
	GetByID(id uint32) (*Provider, error)
	GetByOID(oid primitive.ObjectID) (*Provider, error)
	UpdateArbitrary(id uint32, key string, value interface{}) error
}

// ProviderUsecase for Provider usecase
type ProviderUsecase interface {
	CreateProvider(provider ProviderBody) (uint32, error)
	GetAll() ([]Provider, error)
	GetByID(id uint32) (*Provider, error)
}
