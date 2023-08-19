package mocks

import "PROYECTO/pkg/models"

var ItemElement = []models.Element{
    {Id: "1", Name: "Element 1", Status: "To Do"},
    {Id: "2", Name: "Element 2", Status: "Doing"},
    {Id: "3", Name: "Element 3", Status: "Done"},
}
var Lists = []models.List{
    {Id: "6ba7b810-9dad-11d1-80b4-00c04fd430c8", ClientId: "1", Name: "List 1", Status: "Open", Element: ItemElement},
    {Id: "6ba7b810-9dad-11d1-80b4-00c04fd430c8", ClientId: "1", Name: "List 2", Status: "Open", Element: ItemElement},
}