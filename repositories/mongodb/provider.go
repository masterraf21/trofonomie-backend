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

type providerRepo struct {
	Instance    *mongo.Database
	CounterRepo models.CounterRepository
}

// NewProviderRepo will initiate object representing providerRepository
func NewProviderRepo(instance *mongo.Database, ctr models.CounterRepository) models.ProviderRepository {
	return &providerRepo{Instance: instance, CounterRepo: ctr}
}

func (r *providerRepo) Store(provider *models.Provider) (oid primitive.ObjectID, err error) {
	collectionName := "provider"
	identifier := "provider_id"

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	id, err := r.CounterRepo.Get(collectionName, identifier)
	if err != nil {
		return
	}

	collection := r.Instance.Collection(collectionName)
	provider.ProviderID = id

	result, err := collection.InsertOne(ctx, provider)
	if err != nil {
		return
	}

	_id := result.InsertedID
	oid = _id.(primitive.ObjectID)

	return
}

func (r *providerRepo) GetAll() (res []models.Provider, err error) {
	collection := r.Instance.Collection("provider")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return
	}

	if err = cursor.All(ctx, &res); err != nil {
		return
	}

	return
}

func (r *providerRepo) GetByID(id uint32) (res *models.Provider, err error) {
	collection := r.Instance.Collection("provider")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, bson.M{"provider_id": id}).Decode(&res)
	if err != nil {
		if strings.Contains(err.Error(), "mongo: no documents") {
			err = nil
			return
		}

		return
	}

	return
}

func (r *providerRepo) GetByOID(oid primitive.ObjectID) (res *models.Provider, err error) {
	collection := r.Instance.Collection("provider")

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

func (r *providerRepo) UpdateArbitrary(id uint32, key string, value interface{}) error {
	collection := r.Instance.Collection("provider")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	_, err := collection.UpdateOne(
		ctx,
		bson.M{"provider_id": id},
		bson.M{"$set": bson.M{key: value}},
	)
	if err != nil {
		return err
	}

	return nil
}
