package domain

import (
	"encoding/hex"
	"fmt"
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
	Repr() string
}

func NewEntity(class string, version int) Entity {

	return Entity{uuid.New(), class, version, time.Now(), time.Now()}
}

func ResumeEntity(id uuid.UUID, class string, version int, createdAt time.Time, updatedAt time.Time) Entity {
	return Entity{id, class, version, createdAt, updatedAt}
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

func (e Entity) Repr() string {
	return fmt.Sprintf("id: %s\nclass: %s\nversion: %d\ncreated: %s\nupdated: %s",
		e.IdHexString(), e.Class(), e.Version(),
		e.CreatedAt().String(), e.UpdatedAt().String())
}
