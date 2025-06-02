package domain

import "github.com/google/uuid"

type Team struct {
	Id      uuid.UUID
	Name    string
	Country string
	Players []Player
}
