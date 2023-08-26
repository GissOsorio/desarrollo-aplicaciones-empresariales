package mocks

import (
	"back/pkg/models"
	"time"
	"github.com/google/uuid"
)

var Elements = []models.Element{
	{Id: uuid.New(), Date: time.Now(), ListId: "f525c177-bf96-4d5b-a536-dc89ff85e23e", Content: "Element 1", Status: "todo"},
	{Id: uuid.New(), Date: time.Now(), ListId: "f525c177-bf96-4d5b-a536-dc89ff85e23e", Content: "Element 2", Status: "doing"},
	{Id: uuid.New(), Date: time.Now(), ListId: "f525c177-bf96-4d5b-a536-dc89ff85e23e", Content: "Element 3", Status: "done"},
	{Id: uuid.New(), Date: time.Now(), ListId: "adbce07e-211f-4095-bba7-123e5ce5dcd1", Content: "Element 1", Status: "todo"},
	{Id: uuid.New(), Date: time.Now(), ListId: "adbce07e-211f-4095-bba7-123e5ce5dcd1", Content: "Element 2", Status: "doing"},
	{Id: uuid.New(), Date: time.Now(), ListId: "adbce07e-211f-4095-bba7-123e5ce5dcd1", Content: "Element 3", Status: "done"},
}
