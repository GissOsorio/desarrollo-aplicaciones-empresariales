package models

import (
    "time"
    "github.com/google/uuid"
)

type List struct {
    Id          uuid.UUID   `json:"id"`
    Date		time.Time	`json:"time"`
    UserId      string      `json:"userId"`    
    Name        string      `json:"name"`
    Status      string      `json:"status"`
}