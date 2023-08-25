package mocks

import (
    "time"
	"back/pkg/models"
	"github.com/google/uuid"
)

var Lists = []models.List{
	{Id: uuid.New(), Date: time.Now(), UserId: "1", Name: "Lista de Compras", Status: "Open"},
	{Id: uuid.New(), Date: time.Now(), UserId: "1", Name: "Lista de Ciudades", Status: "Open"},
}
