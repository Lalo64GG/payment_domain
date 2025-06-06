package services

type UUIDGenerator interface {
	GenerateUUID() string
}