package accounts

import "github.com/google/uuid"

type Account struct {
	Id string `json:"id" bson:"_id"`
}

func NewAccount() *Account {
	return &Account{
		Id: uuid.NewString(),
	}
}

