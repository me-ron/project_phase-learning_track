package database

import (
	"context"
	"task_manager/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoCollection struct {
	*mongo.Collection
}

func (c *MongoCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) domain.SingleResultInterface {
	return &domain.MongoSingleResult{SingleResult: c.Collection.FindOne(ctx, filter, opts...)}
}

func (c *MongoCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (domain.CursorInterface, error) {
	cursor, err := c.Collection.Find(ctx, filter, opts...)
	return &domain.MongoCursor{Cursor: cursor}, err
}

func (c *MongoCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return c.Collection.InsertOne(ctx, document, opts...)
}

func (c *MongoCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return c.Collection.UpdateOne(ctx, filter, update, opts...)
}

func (c *MongoCollection) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (domain.DeleteResultInterface, error) {
	result, err := c.Collection.DeleteOne(ctx, filter, opts...)
	return &domain.MongoDeleteResult{DeleteResult: result}, err
}

func (c *MongoCollection) Indexes() mongo.IndexView {
	return c.Collection.Indexes()
}
