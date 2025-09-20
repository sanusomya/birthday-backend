package database

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Icollection interface {
	Find(svc dynamodb.DynamoDB,filter interface{}) (item interface{}, err error)
	InsertOne(svc dynamodb.DynamoDB,filter interface{}) (item interface{}, err error)
	DeleteOne(svc dynamodb.DynamoDB,filter interface{}) (item interface{}, err error)
	UpdateOne(svc dynamodb.DynamoDB,filter interface{}) (item interface{}, err error)
}
