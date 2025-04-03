package repository

import (
	"github.com/PedroNetto404/easy-finance-backend/internal/domain/accounts"
	"github.com/PedroNetto404/easy-finance-backend/pkg/abstractions"
)

type AccountRepository interface {
	abstractions.IRepository[accounts.Account]
}