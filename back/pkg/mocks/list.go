package mocks

import (
    "time"
	"back/pkg/models"
	"github.com/google/uuid"
)

var Lists = []models.List{
	{Id: uuid.New(), Date: time.Now(), UserId: "google-oauth2|116444648959900838283", Name: "Lista de Compras", Status: "Open"},
	{Id: uuid.New(), Date: time.Now(), UserId: "google-oauth2|116444648959900838283", Name: "Lista de Ciudades", Status: "Open"},
}
