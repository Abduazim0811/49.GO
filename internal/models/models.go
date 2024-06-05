package models

type CatFact struct {
    Fact string `json:"fact"`
}

type CatFactResponse struct {
    ID int `json:"id"`
}