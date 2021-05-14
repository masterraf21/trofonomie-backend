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

type customerRepo struct {
	Instance    *mongo.Database
	CounterRepo models.CounterRepository
}

// NewCustomerRepo will create an object representing CustomerRepository
func NewCustomerRepo(instance *mongo.Database, ctr models.CounterRepository) models.CustomerRepository {
	return &customerRepo{Instance: instance, CounterRepo: ctr}
}

func (r *customerRepo) Store(customer *models.Customer) (uid primitive.ObjectID, err error) {
	collectionName := "customer"
	identifier := "customer_id"

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	id, err := r.CounterRepo.Get(collectionName, identifier)
	if err != nil {
		return
	}

	collection := r.Instance.Collection(collectionName)
	customer.CustomerID = id

	result, err := collection.InsertOne(ctx, customer)
	if err != nil {
		return
	}

	_id := result.InsertedID
	uid = _id.(primitive.ObjectID)

	return
}

func (r *customerRepo) GetAll() (res []models.Customer, err error) {
	collection := r.Instance.Collection("customer")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		if strings.Contains(err.Error(), "mongo: no documents") {
			res = make([]models.Customer, 0)
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

func (r *customerRepo) GetByOID(oid primitive.ObjectID) (res *models.Customer, err error) {
	collection := r.Instance.Collection("customer")

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

func (r *customerRepo) GetByID(id uint32) (res *models.Customer, err error) {
	collection := r.Instance.Collection("customer")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, bson.M{"customer_id": id}).Decode(&res)

	if err != nil {
		if strings.Contains(err.Error(), "mongo: no documents") {
			err = nil
			return
		}

		return
	}

	return
}

func (r *customerRepo) UpdateArbitrary(id uint32, key string, value interface{}) error {
	collection := r.Instance.Collection("customer")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	_, err := collection.UpdateOne(
		ctx,
		bson.M{"customer_id": id},
		bson.M{"$set": bson.M{key: value}},
	)
	if err != nil {
		return err
	}

	return nil
}
