package valueObject

import "github.com/google/uuid"

type IDType = uuid.UUID

func ID() IDType {
	return IDType(uuid.New())
}

func ParseID(id string) (IDType, error) {
	uuid, err := uuid.Parse(id)
	return IDType(uuid), err
}
