package models

import "github.com/google/uuid"

type Guest struct {
	ID               uuid.UUID `json:"id"`
	First_name       string    `json:"first_name"`
	Second_name      string    `json:"second_name"`
	Current_room_num uuid.UUID `json:"current_room_num"`
	Created_at       string    `json:"created_at"`
	Deleted_at       string    `json:"deleted_at"`
}