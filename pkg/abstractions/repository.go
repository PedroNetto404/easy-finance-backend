package abstractions

import "github.com/PedroNetto404/easy-finance-backend/pkg/types"

type IRepository[T IAggregateRoot] interface {
	FindById(id string) (T, error)
	FindAll(args types.QueryArgs) (*types.PagedResult[T], error)
	Save(entity T) error
	Delete(id string) error
	Count() (int64, error)
}