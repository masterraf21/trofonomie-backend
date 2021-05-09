package models

// Counter represents counter for
type Counter struct {
	CustomerID uint32 `bson:"customer_id" json:"customer_id"`
	MenuID     uint32 `bson:"menu_id" json:"menu_id"`
	ProviderID uint32 `bson:"provider_id" json:"provider_id"`
	OrderID    uint32 `bson:"order_id" json:"order_id"`
}

// CounterRepository repo for counter
type CounterRepository interface {
	Get(collectionName string, identifier string) (uint32, error)
}
