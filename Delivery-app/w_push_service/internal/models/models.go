package models

type Product struct {
	ID          int     `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Count       uint32  `json:"count,omitempty"`
}

type Message struct {
	ChatID int64          `json:"chat_id"`
	Text   MessageDetails `json:"text"`
}

type MessageDetails struct {
	OrderId  uint32    `json:"order_id"`
	Products []Product `json:"products"`
	TotalSum float64   `json:"total"`
}
