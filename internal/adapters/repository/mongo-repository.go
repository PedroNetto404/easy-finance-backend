package repository

import (
	"context"

	"github.com/PedroNetto404/easy-finance-backend/pkg/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository[T any] struct {
	Collection *mongo.Collection
}

func NewMongoRepository[T any](collection *mongo.Collection) *MongoRepository[T] {
	return &MongoRepository[T]{
		Collection: collection,
	}
}

func (r *MongoRepository[T]) FindById(id string) (*T, error) {
	ctx := context.Background()
	var result T

	err := r.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *MongoRepository[T]) FindAll(args types.QueryArgs) (*types.PagedResult[T], error) {
	ctx := context.Background()
	args.CheckDefaults()

	filter := types.Filter{}
	for k, v := range args.Filter {
		filter[k] = v
	}

	sortOrder := int32(1)
	if !args.Ascending {
		sortOrder = -1
	}

	opts := options.Find().
		SetSort(bson.D{{Key: args.SortBy, Value: sortOrder}}).
		SetSkip(args.Offset * args.Limit).
		SetLimit(args.Limit)

	cursor, err := r.Collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var records []T
	if err = cursor.All(ctx, &records); err != nil {
		return nil, err
	}

	totalCount, err := r.Count(filter)
	if err != nil {
		return nil, err
	}

	pageCount := totalCount / args.Limit
	if totalCount%args.Limit > 0 {
		pageCount++
	}

	hasNext := pageCount > args.Offset+1
	hasPrev := args.Offset > 0

	return &types.PagedResult[T]{
		Meta: types.PagedResultMetadata{
			TotalCount: totalCount,
			PageCount:  pageCount,
			PageNumber: args.Offset + 1,
			PageSize:   args.Limit,
			HasNext:    hasNext,
			HasPrev:    hasPrev,
		},
		Records: records,
	}, nil
}

func (r *MongoRepository[T]) Save(entity *T) error {
	ctx := context.Background()
	_, err := r.Collection.InsertOne(ctx, entity)
	return err
}

func (r *MongoRepository[T]) Update(id string, entity *T) error {
	ctx := context.Background()
	_, err := r.Collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": entity})
	return err
}

func (r *MongoRepository[T]) Delete(id string) error {
	ctx := context.Background()
	_, err := r.Collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *MongoRepository[T]) Count(filter types.Filter) (int64, error) {
	ctx := context.Background()
	return r.Collection.CountDocuments(ctx, filter)
}

func (r *MongoRepository[T]) Exists(filter types.Filter) (bool, error) {
	ctx := context.Background()
	count, err := r.Collection.CountDocuments(ctx, filter)
	return count > 0, err
}