package helper

import "github.com/google/uuid"

type UUID struct{}

func NewUUID() (*UUID, error){
	return &UUID{}, nil
}

func (g *UUID) GenerateUUID() string {
	return uuid.NewString()
}