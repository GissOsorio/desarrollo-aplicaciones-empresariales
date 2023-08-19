package mocks

import (
    "PROYECTO/pkg/models"
    "time"
)


var ItemElement = []models.Element{
    {Id: "1", ListId: "11111111-111-1111-1111-111111111111", Name: "Element 1", Status: "To Do"},
    {Id: "2", ListId: "11111111-111-1111-1111-111111111111", Name: "Element 2", Status: "Doing"},
    {Id: "3", ListId: "11111111-111-1111-1111-111111111111", Name: "Element 3", Status: "Done"},
}

var Elements = []models.Element{
    {Id: "qwe", Date: time.Now(), ListId: "11111111-111-1111-1111-111111111111", Name: "Element 1", Status: "To Do"},
    {Id: "asd", Date: time.Now(), ListId: "11111111-111-1111-1111-111111111111", Name: "Element 2", Status: "Doing"},
    {Id: "asd", Date: time.Now(), ListId: "11111111-111-1111-1111-111111111111", Name: "Element 3", Status: "Done"},
	{Id: "zxc", Date: time.Now(), ListId: "22222222-222-2222-2222-222222222222", Name: "Element 1", Status: "To Do"},
    {Id: "zxc", Date: time.Now(), ListId: "22222222-222-2222-2222-222222222222", Name: "Element 2", Status: "Doing"},
    {Id: "mnb", Date: time.Now(), ListId: "22222222-222-2222-2222-222222222222", Name: "Element 3", Status: "Done"},
}