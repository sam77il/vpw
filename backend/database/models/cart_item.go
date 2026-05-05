package models

import "encoding/json"

type CartItem struct {
	Id        int    `json:"id"`
	UserId    string `json:"user_id"`
	ProductId string `json:"product_id"`
	Amount    int    `json:"amount"`
	Metadata  json.RawMessage `json:"metadata"`
}