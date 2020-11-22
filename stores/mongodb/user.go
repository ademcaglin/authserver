package mongodb

import (
	"context"
	"time"

	"github.com/ademcaglin/authserver/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	IsActive  bool      `json:"is_active"`
}

type MongoUserStore struct {
	userCollection *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client) models.UserStore {
	return &MongoUserStore{
		userCollection: client.Database("sample").Collection("users"),
	}
}

func MapEntityToModel(entity User) (model models.User) {
	return models.User{
		Username:  entity.Username,
		Name:      entity.Name,
		CreatedAt: entity.CreatedAt,
		IsActive:  entity.IsActive,
	}
}

func (store MongoUserStore) GetOne(ctx context.Context, username string) (models.User, error) {
	filter := bson.M{"username": username}
	var entity User
	store.userCollection.FindOne(context.TODO(), filter).Decode(&entity)
	return MapEntityToModel(entity), nil
}

func (store MongoUserStore) Save(ctx context.Context, username string, name string) error {
	entity := User{
		Username:  username,
		Name:      name,
		IsActive:  true,
		CreatedAt: time.Now(),
	}
	store.userCollection.InsertOne(context.TODO(), entity)
	return nil
}
