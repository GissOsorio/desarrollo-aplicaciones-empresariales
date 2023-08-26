package mocks

import (
	"back/pkg/models"
	"time"
	"github.com/google/uuid"
)

var Elements = []models.Element{
	{Id: uuid.New(), Date: time.Now(), ListId: "38e614c4-ea9b-49b6-85ad-f294c6e25fe4", Content: "Element 1", Status: "todo"},
	{Id: uuid.New(), Date: time.Now(), ListId: "38e614c4-ea9b-49b6-85ad-f294c6e25fe4", Content: "Element 2", Status: "doing"},
	{Id: uuid.New(), Date: time.Now(), ListId: "38e614c4-ea9b-49b6-85ad-f294c6e25fe4", Content: "Element 3", Status: "done"},
	{Id: uuid.New(), Date: time.Now(), ListId: "fd99648f-61e6-4437-bad8-1410c0dcef3b", Content: "Element 1", Status: "todo"},
	{Id: uuid.New(), Date: time.Now(), ListId: "fd99648f-61e6-4437-bad8-1410c0dcef3b", Content: "Element 2", Status: "doing"},
	{Id: uuid.New(), Date: time.Now(), ListId: "fd99648f-61e6-4437-bad8-1410c0dcef3b", Content: "Element 3", Status: "done"},
}
