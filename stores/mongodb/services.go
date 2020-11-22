package mongodb

import (
	"github.com/ademcaglin/authserver/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoOptions struct {
	Constr string
}

type MongoServices struct {
	UserStore   models.UserStore
	ClientStore models.UserStore
}

func GetServices(appOptions MongoOptions) *MongoServices {
	client, _ := mongo.NewClient(options.Client().ApplyURI(appOptions.Constr))
	return &MongoServices{
		UserStore: NewMongoUserStore(client),
	}
}
