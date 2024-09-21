package database

import (
	"birthday/birthday"
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type fakeCollection struct {
}

func (f fakeCollection) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return &mongo.DeleteResult{}, nil
}

func (f fakeCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return &mongo.UpdateResult{}, nil
}

func (f fakeCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (cur *mongo.Cursor, err error) {
	c := &mongo.Cursor{}
	return c, nil
}
func (f fakeCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	return &mongo.SingleResult{}
}
func (f fakeCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return &mongo.InsertOneResult{}, nil
}
func TestDatabase(t *testing.T) {
	t.Run("Testing Add to database method", func(t *testing.T) {
		f := fakeCollection{}
		err := Add(f, birthday.Birthday{})
		if err != nil {
			t.Errorf("Error from Add method expected nil returned %v", err)
		}
	})

	t.Run("Testing Getall from database method", func(t *testing.T) {
		f := fakeCollection{}
		_, err := GetAll(f)
		if err != nil {
			t.Errorf("Error from Getall method expected nil returned %v", err)
		}
	})

	t.Run("Testing Delete from database method", func(t *testing.T) {
		f := fakeCollection{}
		err := Delete(f, birthday.Birthday{})
		if err == nil {
			t.Errorf("Error from delete method expected nil returned %v", err)
		}
	})

	t.Run("Testing Edit from database method", func(t *testing.T) {
		f := fakeCollection{}
		err := Edit(f, "test", 636536656, birthday.Birthday{})
		if err == nil {
			t.Errorf("Error from Getall method expected nil returned %v", err)
		}
	})
	t.Run("Testing connect db", func(t *testing.T) {
		_, err := ConnectDB("", "", "")
		if err == nil {
			t.Errorf("Error from connect db expected nil returned %v", err)
		}
	})
	t.Run("Testing Find By Name And Mobile", func(t *testing.T) {
		f := fakeCollection{}
		_,err := FindByNameAndMobile(f, 636536656,"test")
		if err == nil {
			t.Errorf("Error from Find By Name And Mobile method expected nil returned %v", err)
		}
	})
	t.Run("Testing Find For This Month", func(t *testing.T) {
		f := fakeCollection{}
		_,err := FindForThisMonth(f, "jan")
		if err != nil {
			t.Errorf("Error from Find For This Month method expected nil returned %v", err)
		}
	})
	t.Run("Testing Find For Today", func(t *testing.T) {
		f := fakeCollection{}
		_,err := FindForToday(f, "jan", 2)
		if err != nil {
			t.Errorf("Error from Find For Today method expected nil returned %v", err)
		}
	})
}
