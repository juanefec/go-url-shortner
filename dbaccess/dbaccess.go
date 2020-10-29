package dbaccess

import (
	"context"
	"fmt"

	"github.com/jkomyno/nanoid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	username = "creative"
	password = "asdfgh123456"

	dbname = "testdb"
)

type URLStore struct {
	urlID string `bson: "url_id", json: "url_id"`
	raw   string `bson: "raw", json: "raw"`
}

var clientOptions *options.ClientOptions

func StoreURL(url string) (string, error) {
	client, e := mongo.Connect(context.TODO(), clientOptions)
	if e != nil {
		return "", e
	}

	urlstore := client.Database("testdb").Collection("new_url_store")

	id, e := nanoid.Nanoid(10)
	if e != nil {
		return "", e
	}

	r, e := urlstore.InsertOne(context.TODO(), bson.D{
		{Key: "raw", Value: url},
		{Key: "url_id", Value: id},
	})

	if e != nil {
		return "", e
	}

	_, ok := r.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", e
	}
	return id, nil
}

func GetURL(id string) (string, error) {
	client, e := mongo.Connect(context.TODO(), clientOptions)
	if e != nil {
		return "", e
	}

	urlstore := client.Database("testdb").Collection("new_url_store")

	filter := bson.D{{"url_id", id}}
	var res URLStore
	e = urlstore.FindOne(context.TODO(), filter).Decode(&res)
	if e != nil {
		return "", e
	}
	fmt.Println("res.raw: " + res.raw)
	return res.raw, nil
}

func init() {

	// Set client options
	clientOptions = options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://creative:%v@cluster0.n0fo5.mongodb.net/%v?retryWrites=true&w=majority", password, dbname))

}
