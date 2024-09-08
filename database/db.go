package database

import (
	"birthday/birthday"
	"context"
	"os"
	"time"
	"birthday/utils"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() (mongo.Collection, error) {
	godotenv.Load()
	var uri = os.Getenv("db_url")
	var database = os.Getenv("database")
	var collection = os.Getenv("db_coll")
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return mongo.Collection{}, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		return mongo.Collection{}, err
	}

	db := client.Database(database)
	coll := db.Collection(collection)
	return *coll, nil
}

func GetAll(coll mongo.Collection) ([]birthday.Birthday, error) {
	var birth []birthday.Birthday
	ans, err := coll.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return []birthday.Birthday{}, err
	}
	for ans.Next(context.TODO()) {
		var temp birthday.Birthday
		ans.Decode(&temp)
		birth = append(birth, temp)
	}

	return birth, nil
}

func Add(coll mongo.Collection, b birthday.Birthday) error {
	temp := birthday.Birthday{}
	res := coll.FindOne(context.TODO(), bson.D{{"name", b.Name}, {"mobile", b.Mobile}})
	err := res.Decode(&temp)
	if err == nil {
		return utils.AlreadyFind{}
	}
	_, err = coll.InsertOne(context.TODO(), b)
	return err
}

func Delete(coll mongo.Collection, b birthday.Birthday) error {
	_, err := coll.DeleteOne(context.TODO(), bson.D{{"name", b.Name}, {"month", b.Month}, {"date", b.Date}, {"mobile", b.Mobile}})
	return err
}

func Edit(coll mongo.Collection, name string, mobile int64, b birthday.Birthday) error {
	filter := bson.D{{"name", name}, {"mobile", mobile}}
	update := bson.D{{"$set", bson.D{{"name", b.Name}, {"date", b.Date}, {"month", b.Month}, {"mobile", b.Mobile}}}}

	_, err := coll.UpdateOne(context.TODO(), filter, update, nil)
	return err
}
