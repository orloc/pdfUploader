package entity

import "time"

type Deck struct {
	Id int64 `json:"-"`
	Uuid string `json:"uuid"`
	CompanyName string `json:"company_name"`
	Images []string `json:"images"`
	CreatedAt time.Time `json:"created_at"`
}

func (r *Deck) TableName() string {
	return "decks"
}
