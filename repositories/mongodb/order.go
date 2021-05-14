package mongodb

import (
	"context"
	"strings"
	"time"

	"github.com/masterraf21/trofonomie-backend/configs"
	"github.com/masterraf21/trofonomie-backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type orderRepo struct {
	Instance    *mongo.Database
	CounterRepo models.CounterRepository
}

// NewOrderRepo will initiate order repository object
func NewOrderRepo(instance *mongo.Database, ctr models.CounterRepository) models.OrderRepository {
	return &orderRepo{Instance: instance, CounterRepo: ctr}
}

func (r *orderRepo) Store(order *models.Order) (oid primitive.ObjectID, err error) {
	collectionName := "order"
	identifier := "order_id"

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	id, err := r.CounterRepo.Get(collectionName, identifier)
	if err != nil {
		return
	}

	collection := r.Instance.Collection(collectionName)
	order.ID = id

	result, err := collection.InsertOne(ctx, order)
	if err != nil {
		return
	}

	_id := result.InsertedID
	oid = _id.(primitive.ObjectID)

	return
}

func (r *orderRepo) GetAll() (res []models.Order, err error) {
	collection := r.Instance.Collection("order")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		if strings.Contains(err.Error(), "mongo: no documents") {
			res = make([]models.Order, 0)
			err = nil
			return
		}

		return
	}

	if err = cursor.All(ctx, &res); err != nil {
		return
	}

	return
}

func (r *orderRepo) GetByID(id uint32) (res *models.Order, err error) {
	collection := r.Instance.Collection("order")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, bson.M{"order_id": id}).Decode(&res)
	if err != nil {
		if strings.Contains(err.Error(), "mongo: no documents") {
			err = nil
			return
		}

		return
	}

	return
}

func (r *orderRepo) GetByOID(oid primitive.ObjectID) (res *models.Order, err error) {
	collection := r.Instance.Collection("order")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&res)
	if err != nil {
		if strings.Contains(err.Error(), "mongo: no documents") {
			err = nil
			return
		}

		return
	}

	return
}

func (r *orderRepo) UpdateArbitrary(id uint32, key string, value interface{}) error {
	collection := r.Instance.Collection("order")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	_, err := collection.UpdateOne(
		ctx,
		bson.M{"order_id": id},
		bson.M{"$set": bson.M{key: value}},
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *orderRepo) GetByProviderID(ProviderID uint32) (res []models.Order, err error) {
	collection := r.Instance.Collection("order")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{"provider_id": ProviderID})
	if err != nil {
		if strings.Contains(err.Error(), "mongo: no documents") {
			res = make([]models.Order, 0)
			err = nil
			return
		}

		return
	}

	if err = cursor.All(ctx, &res); err != nil {
		return
	}

	return
}

func (r *orderRepo) GetByCustomerID(CustomerID uint32) (res []models.Order, err error) {
	collection := r.Instance.Collection("order")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{"customer_id": CustomerID})
	if err != nil {
		if strings.Contains(err.Error(), "mongo: no documents") {
			res = make([]models.Order, 0)
			err = nil
			return
		}

		return
	}

	if err = cursor.All(ctx, &res); err != nil {
		return
	}

	return
}

func (r *orderRepo) GetByCustomerIDAndStatus(CustomerID uint32, status string) (res []models.Order, err error) {
	collection := r.Instance.Collection("order")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{"customer_id": CustomerID, "status": status})
	if err != nil {
		if strings.Contains(err.Error(), "mongo: no documents") {
			res = make([]models.Order, 0)
			err = nil
			return
		}

		return
	}

	if err = cursor.All(ctx, &res); err != nil {
		return
	}

	return
}

func (r *orderRepo) GetByProviderIDAndStatus(ProviderID uint32, status string) (res []models.Order, err error) {
	collection := r.Instance.Collection("order")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{"provider_id": ProviderID, "status": status})
	if err != nil {
		if strings.Contains(err.Error(), "mongo: no documents") {
			res = make([]models.Order, 0)
			err = nil
			return
		}

		return
	}

	if err = cursor.All(ctx, &res); err != nil {
		return
	}

	return
}

func (r *orderRepo) GetByStatus(status string) (res []models.Order, err error) {
	collection := r.Instance.Collection("order")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{"status": status})
	if err != nil {
		if strings.Contains(err.Error(), "mongo: no documents") {
			res = make([]models.Order, 0)
			err = nil
			return
		}

		return
	}

	if err = cursor.All(ctx, &res); err != nil {
		return
	}

	return
}
