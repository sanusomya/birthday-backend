package database

import (
	"birthday/birthday"
	"birthday/utils"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(uri string, database string, collection string) (Icollection, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return &mongo.Collection{}, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		return &mongo.Collection{}, err
	}

	db := client.Database(database)
	coll := db.Collection(collection)
	return coll, nil
}

func GetAll(coll Icollection) ([]birthday.Birthday, error) {
	var birth []birthday.Birthday
	ans, err := coll.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return []birthday.Birthday{}, err
	}
	if ans.RemainingBatchLength() == 0 {
		return []birthday.Birthday{}, nil
	}
	for ans.Next(context.TODO()) {
		var temp birthday.Birthday
		ans.Decode(&temp)
		birth = append(birth, temp)
	}
	ans.Close(context.TODO())
	return birth, nil
}

func Add(coll Icollection, b birthday.Birthday) error {
	err := findDuplicate(coll, b.Name, b.Mobile)
	if err != nil {
		return err
	}
	_, err = coll.InsertOne(context.TODO(), b)
	return err
}

func Delete(coll Icollection, b birthday.Birthday) error {
	res, err := coll.DeleteOne(context.TODO(), bson.D{{"name", b.Name}, {"month", b.Month}, {"date", b.Date}, {"mobile", b.Mobile}})
	if res.DeletedCount == 0{
		return utils.NotFound{}
	}
	return err
}

func Edit(coll Icollection, name string, mobile int64, b birthday.Birthday) error {
	filter := bson.D{{"name", name}, {"mobile", mobile}}
	update := bson.D{{"$set", bson.D{{"name", b.Name}, {"date", b.Date}, {"month", b.Month}, {"mobile", b.Mobile}}}}

	res, err := coll.UpdateOne(context.TODO(), filter, update, nil)
	if res.MatchedCount == 0{
		return utils.NotFound{}
	}
	return err
}



func findDuplicate(coll Icollection, name string, mobile int64) error {
	ans := coll.FindOne(context.TODO(), bson.D{{"name", name}, {"mobile", mobile}})
	var temp = birthday.Birthday{}
	ans.Decode(&temp)
	if temp.Date != 0 {
		return utils.AlreadyFind{}
	}
	return nil
}

func FindByNameAndMobile(coll Icollection, phone int64, name string) (birthday.Birthday, error) {
	var bday birthday.Birthday
	ans := coll.FindOne(context.TODO(), bson.D{{"mobile", phone}, {"name", name}})
	ans.Decode(&bday)
	if bday.Date == 0{
		return birthday.Birthday{}, utils.NotFound{}
	}
	return bday, nil
}

func FindForThisMonth(coll Icollection, mon string) ([]birthday.Birthday, error) {
	var birth []birthday.Birthday
	ans, err := coll.Find(context.TODO(), bson.D{{"month", mon}})
	if err != nil {
		return []birthday.Birthday{}, err
	}
	if ans.RemainingBatchLength() == 0 {
		return []birthday.Birthday{}, nil
	}
	for ans.Next(context.TODO()) {
		var temp birthday.Birthday
		ans.Decode(&temp)
		birth = append(birth, temp)
	}
	ans.Close(context.TODO())
	return birth, nil
}

func FindForToday(coll Icollection, mon string, date int8) ([]birthday.Birthday, error) {
	var birth []birthday.Birthday
	ans, err := coll.Find(context.TODO(), bson.D{{"month", mon}, {"date", date}})
	if err != nil {
		return []birthday.Birthday{}, err
	}
	if ans.RemainingBatchLength() == 0 {
		return []birthday.Birthday{}, nil
	}
	for ans.Next(context.TODO()) {
		var temp birthday.Birthday
		ans.Decode(&temp)
		birth = append(birth, temp)
	}
	ans.Close(nil)
	return birth, nil
}
