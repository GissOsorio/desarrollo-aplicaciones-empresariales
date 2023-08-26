package models

import (
    "time"
    "github.com/google/uuid"
)

type Element struct {
    Id          uuid.UUID   `json:"id"`
	ListId      string      `json:"listId"`
	Date		time.Time	`json:"time"`
    Content     string      `json:"content"`
    Status      string      `json:"status"`
}