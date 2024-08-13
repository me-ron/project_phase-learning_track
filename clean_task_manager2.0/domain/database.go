package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CollectionInterface interface {
	FindOne(context.Context, interface{}, ...*options.FindOneOptions) SingleResultInterface
	Find(context.Context, interface{}, ...*options.FindOptions) (CursorInterface, error)
	InsertOne(context.Context, interface{}, ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	UpdateOne(context.Context, interface{}, interface{}, ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	DeleteOne(context.Context, interface{}, ...*options.DeleteOptions) (DeleteResultInterface, error)
	Indexes() IndexView
}

type IndexView interface{
	CreateOne(ctx context.Context, model mongo.IndexModel, opts ...*options.CreateIndexesOptions) (string, error)
}

type CursorInterface interface {
	Next(context.Context) bool
	Decode(interface{}) error
	Close(context.Context) error
}

type SingleResultInterface interface {
	Decode(v interface{}) error
}

type DeleteResultInterface interface {
	DeletedCount() int64
}
type MongoCursor struct {
	*mongo.Cursor
}

type MongoIndexView struct{
	mongo.IndexView
}

func (MI *MongoIndexView) CreateOne(ctx context.Context, model mongo.IndexModel, opts ...*options.CreateIndexesOptions) (string, error){
	return MI.IndexView.CreateOne(ctx, model , opts ...)
}

func (c *MongoCursor) Next(ctx context.Context) bool {
	return c.Cursor.Next(ctx)
}

func (c *MongoCursor) Decode(v interface{}) error {
	return c.Cursor.Decode(v)
}

func (c *MongoCursor) Close(ctx context.Context) error {
	return c.Cursor.Close(ctx)
}

type MongoSingleResult struct {
	*mongo.SingleResult
}

func (r *MongoSingleResult) Decode(v interface{}) error {
	return r.SingleResult.Decode(v)
}

type MongoDeleteResult struct {
	*mongo.DeleteResult
}

func (r *MongoDeleteResult) DeletedCount() int64 {
	return r.DeleteResult.DeletedCount
}