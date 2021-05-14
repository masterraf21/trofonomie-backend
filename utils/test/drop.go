package test

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// DropCounter will drop counter table
func DropCounter(ctx context.Context, db *mongo.Database) (err error) {
	collection := db.Collection("counter")
	err = collection.Drop(ctx)
	return
}

// DropMenu will drop buyer table
func DropMenu(ctx context.Context, db *mongo.Database) (err error) {
	collection := db.Collection("menu")
	err = collection.Drop(ctx)
	return
}

// DropProvider will drop seller table
func DropProvider(ctx context.Context, db *mongo.Database) (err error) {
	collection := db.Collection("provider")
	err = collection.Drop(ctx)
	return
}

// DropCustomer will drop product table
func DropCustomer(ctx context.Context, db *mongo.Database) (err error) {
	collection := db.Collection("customer")
	err = collection.Drop(ctx)
	return
}

// DropOrder will drop order table
func DropOrder(ctx context.Context, db *mongo.Database) (err error) {
	collection := db.Collection("order")
	err = collection.Drop(ctx)
	return
}
