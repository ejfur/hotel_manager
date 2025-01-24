package models

import (
	"github.com/google/uuid"

)

type Room struct {
	ID           uuid.UUID `json:"id"`
	Room_num     int       `json:"room_num"`
	Stage        int       `json:"stage"`
	Last_cleaned string    `json:"last_cleaned"`
	Ocupied      bool      `json:"ocupied"`
}
