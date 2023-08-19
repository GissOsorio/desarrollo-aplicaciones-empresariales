package models



type List struct {
    Id          string      `json:"id"`
    ClientId    string      `json:"clientId"`    
    Name        string      `json:"name"`
    Status      string      `json:"status"`
    Element     []Element   `json:"element"`
}

type Element struct {
    Id          string      `json:"id"`
    Name        string      `json:"name"`
    Status      string      `json:"desc"`

}