package domain

import "github.com/google/uuid"

type Repository struct{}

type IRepository interface {
	List() []Entity
	Retrieve(uuid.UUID) Entity
	Add(Entity)
	Remove(uuid.UUID)
	Filter(func(Entity) bool) []Entity
}
