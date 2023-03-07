package dto

import "github.com/google/uuid"

type PipeServerModel struct {
	ID     uuid.UUID `json:"id"`
	Input1 string    `json:"input_1"`
	INput2 string    `json:"input_2"`
}
