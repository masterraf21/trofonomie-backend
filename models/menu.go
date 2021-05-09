package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Menu represents menu data
type Menu struct {
	MenuID      uint32    `json:"menu_id" bson:"menu_id"`
	ProviderID  uint32    `json:"provider_id" bson:"provider_id"`
	Provider    *Provider `json:"provider" bson:"provider"`
	Name        string    `json:"name" bson:"name"`
	Description string    `bson:"description" json:"description"`
	Price       float32   `json:"price" bson:"price"`
	Calorie     float32   `json:"calorie" bson:"calorie"`
	ImageURL    string    `json:"image_url" bson:"image_url"`
}

// MenuBody for receiving body grom json
type MenuBody struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Calorie     float32 `json:"calorie"`
	ImageURL    string  `json:"image_url"`
	ProviderID  uint32  `json:"id_provider"`
}

// MenuRepository represents repo functions for Menu
type MenuRepository interface {
	Store(menu *Menu) (primitive.ObjectID, error)
	GetAll() ([]Menu, error)
	GetByProviderID(providerID uint32) ([]Menu, error)
	GetByID(id uint32) (*Menu, error)
	GetByOID(oid primitive.ObjectID) (*Menu, error)
	UpdateArbitrary(id uint32, key string, value interface{}) error
}

// MenuUsecase usecase for Menu
type MenuUsecase interface {
	CreateMenu(Menu MenuBody) (uint32, error)
	GetAll() ([]Menu, error)
	GetByProviderID(ProviderID uint32) ([]Menu, error)
	GetByID(id uint32) (*Menu, error)
}
