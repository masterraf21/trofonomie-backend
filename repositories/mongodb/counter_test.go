package mongodb

import (
	"context"
	"testing"
	"time"

	"github.com/masterraf21/trofonomie-backend/configs"
	"github.com/masterraf21/trofonomie-backend/models"
	"github.com/masterraf21/trofonomie-backend/utils/mongodb"
	testUtil "github.com/masterraf21/trofonomie-backend/utils/test"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type counterRepoTestSuite struct {
	suite.Suite
	Instance    *mongo.Database
	CounterRepo models.CounterRepository
}

func TestCounterRepository(t *testing.T) {
	suite.Run(t, new(counterRepoTestSuite))
}

func (s *counterRepoTestSuite) SetupSuite() {
	instance := mongodb.ConfigureMongo()
	s.Instance = instance
	s.CounterRepo = NewCounterRepo(instance)
}

func (s *counterRepoTestSuite) TearDownTest() {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	err := testUtil.DropMenu(ctx, s.Instance)
	testUtil.HandleError(err)
	err = testUtil.DropCounter(ctx, s.Instance)
	testUtil.HandleError(err)
}

func (s *counterRepoTestSuite) TearDownSuite() {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
	defer cancel()

	err := testUtil.DropMenu(ctx, s.Instance)
	testUtil.HandleError(err)
	err = testUtil.DropCounter(ctx, s.Instance)
	testUtil.HandleError(err)
}

func (s *counterRepoTestSuite) TestGetEmpty() {
	s.Run("Get id with empty document", func() {
		collectionName := "menu"
		identifier := "menu_id"

		id, err := s.CounterRepo.Get(collectionName, identifier)
		testUtil.HandleError(err)

		s.Equal(uint32(1), id)
	})
}

func (s *counterRepoTestSuite) TestGetExisting() {
	s.Run("Get id with existing document", func() {
		ctx, cancel := context.WithTimeout(context.Background(), configs.Constant.TimeoutOnSeconds*time.Second)
		defer cancel()

		collectionName := "menu"
		identifier := "menu_id"

		collection := s.Instance.Collection(collectionName)
		menu := models.Menu{
			MenuID:      1,
			ProviderID:  1,
			Name:        "test",
			Description: "test",
			Price:       120.1,
			Calorie:     100.1,
			ImageURL:    "http",
		}
		_, err := collection.UpdateOne(
			ctx,
			bson.M{identifier: menu.MenuID},
			bson.M{"$set": menu},
			options.Update().SetUpsert(true),
		)
		testUtil.HandleError(err)

		id, err := s.CounterRepo.Get(collectionName, identifier)
		testUtil.HandleError(err)

		s.Assert().EqualValues(1, id)
	})
}
