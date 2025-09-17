package domain

import (
	"encoding/hex"
	"time"

	"github.com/google/uuid"
)

type Entity struct {
	id        uuid.UUID
	class     string
	version   int
	createdAt time.Time
	updatedAt time.Time
}

type IEntity interface {
	Id() uuid.UUID
	IdString() string
	IdHexString() string
	Class() string
	Version() int
	CreatedAt() time.Time
	UpdatedAt() time.Time
}

func NewEntity(class string, version int) Entity {
	return Entity{uuid.New(), class, version, time.Now(), time.Now()}
}

func (e Entity) Id() uuid.UUID {
	return e.id
}

func (e Entity) IdString() string {
	return e.id.String()
}

func (e Entity) IdHexString() string {
	data, _ := e.id.MarshalBinary()
	return hex.EncodeToString(data)
}

func (e Entity) Class() string {
	return e.class
}

func (e Entity) Version() int {
	return e.version
}

func (e Entity) CreatedAt() time.Time {
	return e.createdAt
}

func (e Entity) UpdatedAt() time.Time {
	return e.updatedAt
}
