package models

import (
    "time"
)

type Element struct {
    Id          string      `json:"id"`
	ListId      string      `json:"listId"`
	Date		time.Time	`json:"time"`
    Name        string      `json:"name"`
    Status      string      `json:"status"`
}