package models

import (
    "time"
)

type User struct {
    Id          string      `json:"id"`
	Date		time.Time	`json:"time"`
    Name        string      `json:"name"`
}