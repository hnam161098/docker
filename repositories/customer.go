package repositories

import (
	"context"
	"docker/connections"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collectionName = "customer"
var db = connections.Mongodb.Collection(collectionName)

type CustomerModel struct {
	Id        primitive.ObjectID `bson:"_id" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Address   string             `bson:"address" json:"addess"`
	Age       int64              `bson:"age" json:"age"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

func NewCustomerModel() CustomerModel {
	return CustomerModel{}
}

func (CustomerModel) Collection() string {
	return collectionName
}

func (c *CustomerModel) InsertOneCustomer(ctx context.Context, model CustomerModel) error {
	if model.Id.IsZero() {
		model.Id = primitive.NewObjectID()
	}
	_, errC := db.InsertOne(ctx, model)
	return errC
}

func (c *CustomerModel) FindCustomer(ctx context.Context) ([]CustomerModel, error) {
	var result []CustomerModel
	cursor, errF := db.Find(ctx, bson.M{})
	if errF != nil {
		return result, errF
	}
	err := cursor.All(ctx, &result)
	return result, err
}
