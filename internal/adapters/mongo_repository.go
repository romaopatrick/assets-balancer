package adapters

import (
	"balancer/internal/ports"
	"context"

	"reflect"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDbRepository[T interface{}] struct {
	collection *mongo.Collection
}

func NewMongoDbRepository[T interface{}](
	db *mongo.Database) ports.Repository[T] {

	var r T
	coll := db.Collection(strings.ToLower(reflect.TypeOf(r).Elem().Name()))
	return &MongoDbRepository[T]{
		collection: coll,
	}
}

func (r *MongoDbRepository[T]) GetAll(
	ctx context.Context,
	filter map[string]interface{}) []T {

	cur, err := r.collection.Find(ctx, filter)
	if err != nil {
		panic(err)
	}
	result := []T{}
	for cur.Next(ctx) {
		var el T
		err = cur.Decode(&el)
		if err != nil {
			panic(err)
		}
		result = append(result, el)
	}

	return result
}

func (r *MongoDbRepository[T]) GetAllSkipTake(
	ctx context.Context,
	filter map[string]interface{},
	skip int64,
	take int64) []T {

	op := options.Find()
	op.SetSkip(skip)
	op.SetLimit(take)
	cur, err := r.collection.Find(ctx, filter, op)

	if err != nil {
		panic(err)
	}
	result := []T{}
	for cur.Next(ctx) {
		var el T
		err = cur.Decode(&el)
		if err != nil {
			panic(err)
		}
		result = append(result, el)
	}

	return result
}

func (r *MongoDbRepository[T]) GetFirst(
	ctx context.Context,
	filter map[string]interface{}) T {
	var el T
	err := r.collection.FindOne(ctx, filter).Decode(&el)

	if err == mongo.ErrNoDocuments {
		return el
	}

	if err != nil {
		panic(err)
	}

	return el
}

func (r *MongoDbRepository[T]) Insert(
	ctx context.Context,
	entity T) {
	_, err := r.collection.InsertOne(ctx, entity)
	if err != nil {
		panic(err)
	}
}

func (r *MongoDbRepository[T]) Replace(
	ctx context.Context,
	filter map[string]interface{},
	entity T) {

	_, err := r.collection.ReplaceOne(ctx, filter, entity)
	if err != nil {
		panic(err)
	}
}

func (r *MongoDbRepository[T]) DeleteAll(
	ctx context.Context,
	filter map[string]interface{}) {
	_, err := r.collection.DeleteMany(ctx, filter)
	if err != nil {
		panic(err)
	}
}
