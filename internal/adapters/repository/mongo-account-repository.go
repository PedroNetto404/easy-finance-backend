package repository

import (
	"github.com/PedroNetto404/easy-finance-backend/internal/domain/accounts"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoAccountRepository struct {
	MongoRepository[accounts.Account]
	collection *mongo.Collection
}

func NewMongoAccountRepository(collection *mongo.Collection) *mongoAccountRepository {
	return &mongoAccountRepository{
		collection: collection,
	}
}