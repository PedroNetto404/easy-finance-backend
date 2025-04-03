package abstractions

import "github.com/PedroNetto404/easy-finance-backend/pkg/types"

type IRepository[T any] interface {
	FindById(id string) (*T, error)
	FindAll(args types.QueryArgs) (*types.PagedResult[T], error)
	Save(entity *T) error
	Delete(id string) error
	Update(entity *T) error
	Count(filter types.Filter) (int64, error)
	Exists(filter types.Filter) (bool, error)
}