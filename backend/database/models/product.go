package models

import "encoding/json"

type Product struct {
	Id          string          `json:"id"`
	Label       string          `json:"label"`
	Description string          `json:"description"`
	Price       float64         `json:"price"`
	Stock       int             `json:"stock"`
	CategoryId  string          `json:"category_id"`
	Options     json.RawMessage `json:"options"`
	OldPrice float64 `json:"old_price"`
}