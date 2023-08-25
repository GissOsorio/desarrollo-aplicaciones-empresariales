package mocks

import (
	"back/pkg/models"
	"time"
	"github.com/google/uuid"
)

var Elements = []models.Element{
	{Id: uuid.New(), Date: time.Now(), ListId: "474e6244-ec83-4013-aaa3-748cc5d17d87", Name: "Element 1", Status: "To Do"},
	{Id: uuid.New(), Date: time.Now(), ListId: "474e6244-ec83-4013-aaa3-748cc5d17d87", Name: "Element 2", Status: "Doing"},
	{Id: uuid.New(), Date: time.Now(), ListId: "474e6244-ec83-4013-aaa3-748cc5d17d87", Name: "Element 3", Status: "Done"},
	{Id: uuid.New(), Date: time.Now(), ListId: "93d5fa09-6798-4bdd-af8c-ed9b5da2fe33", Name: "Element 1", Status: "To Do"},
	{Id: uuid.New(), Date: time.Now(), ListId: "93d5fa09-6798-4bdd-af8c-ed9b5da2fe33", Name: "Element 2", Status: "Doing"},
	{Id: uuid.New(), Date: time.Now(), ListId: "93d5fa09-6798-4bdd-af8c-ed9b5da2fe33", Name: "Element 3", Status: "Done"},
}
