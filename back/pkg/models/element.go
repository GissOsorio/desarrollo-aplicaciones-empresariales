package models

import (
    "time"
    "github.com/google/uuid"
)

type Element struct {
    Id          uuid.UUID   `json:"id"`
	ListId      string      `json:"listId"`
	Date		time.Time	`json:"time"`
    Name        string      `json:"name"`
    Status      string      `json:"status"`
}