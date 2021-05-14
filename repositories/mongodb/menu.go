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

type menuRepo struct {
	Instance    *mongo.Database
	CounterRepo models.CounterRepository
}

// NewMenuRepo will initiate product repo
func NewMenuRepo(instance *mongo.Database, ctr models.CounterRepository) models.MenuRepository {
	return &menuRepo{Instance: instance, CounterRepo: ctr}
}

func (r *menuRepo) Store(menu *models.Menu) (oid primitive.ObjectID, err error) {
	collectionName := "menu"
	identifier := "menu_id"

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	id, err := r.CounterRepo.Get(collectionName, identifier)
	if err != nil {
		return
	}

	collection := r.Instance.Collection(collectionName)
	menu.MenuID = id

	result, err := collection.InsertOne(ctx, menu)
	if err != nil {
		return
	}

	_id := result.InsertedID
	oid = _id.(primitive.ObjectID)

	return
}

func (r *menuRepo) GetAll() (res []models.Menu, error error) {
	collection := r.Instance.Collection("menu")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		if strings.Contains(err.Error(), "mongo: no documents") {
			res = make([]models.Menu, 0)
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

func (r *menuRepo) GetByProviderID(providerID uint32) (res []models.Menu, err error) {
	collection := r.Instance.Collection("menu")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{"provider_id": providerID})
	if err != nil {
		if strings.Contains(err.Error(), "mongo: no documents") {
			res = make([]models.Menu, 0)
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

func (r *menuRepo) GetByID(id uint32) (res *models.Menu, err error) {
	collection := r.Instance.Collection("menu")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, bson.M{"menu_id": id}).Decode(&res)
	if err != nil {
		if strings.Contains(err.Error(), "mongo: no documents") {
			err = nil
			return
		}

		return
	}

	return
}

func (r *menuRepo) GetByOID(oid primitive.ObjectID) (res *models.Menu, err error) {
	collection := r.Instance.Collection("menu")

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

func (r *menuRepo) UpdateArbitrary(id uint32, key string, value interface{}) error {
	collection := r.Instance.Collection("menu")

	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	_, err := collection.UpdateOne(
		ctx,
		bson.M{"menu_id": id},
		bson.M{"$set": bson.M{key: value}},
	)
	if err != nil {
		return err
	}

	return nil
}
