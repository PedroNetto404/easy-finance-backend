package abstractions

import (
	"github.com/google/uuid"
)

type (
	IDomainEvent any

	IEntity interface {
		Id() string
	}

	Entity struct {
		id string
	}

	IAggregateRoot interface {
		HasDomainEvents() bool
		GetDomainEvents() []IDomainEvent
		ClearDomainEvents()
	}

	AggregateRoot struct {
		Entity
		domainEvents []IDomainEvent
	}
)

func NewEntity() IEntity {
	return &Entity{
		id: uuid.NewString(),
	}
}

func (e *Entity) Id() string {
	return e.id
}

func NewAggregateRoot() IAggregateRoot {
	return &AggregateRoot{
		Entity: Entity{
			id: uuid.NewString(),
		},
		domainEvents: make([]IDomainEvent, 0),
	}
}

func (a *AggregateRoot) HasDomainEvents() bool {
	return len(a.domainEvents) > 0
}

func (a *AggregateRoot) GetDomainEvents() []IDomainEvent {
	return a.domainEvents
}

func (a *AggregateRoot) ClearDomainEvents() {
	a.domainEvents = make([]IDomainEvent, 0)
}


