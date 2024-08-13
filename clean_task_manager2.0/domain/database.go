package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CollectionInterface defines the methods used in UserRepo
type CollectionInterface interface {
	FindOne(context.Context, interface{}, ...*options.FindOneOptions) SingleResultInterface
	Find(context.Context, interface{}, ...*options.FindOptions) (CursorInterface, error)
	InsertOne(context.Context, interface{}, ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	UpdateOne(context.Context, interface{}, interface{}, ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	DeleteOne(context.Context, interface{}, ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	Indexes() mongo.IndexView
}

// CursorInterface defines the methods used on the cursor
type CursorInterface interface {
	Next(context.Context) bool
	Decode(interface{}) error
	Close(context.Context) error
}

// SingleResultInterface defines the methods used on a single result
type SingleResultInterface interface {
	Decode(v interface{}) error
}

// DeleteResultInterface defines the methods used on a delete result (optional, mostly for testing purposes)
type DeleteResultInterface interface {
	DeletedCount() int64
}