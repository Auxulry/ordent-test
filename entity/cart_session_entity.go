package entity

type CartSession struct {
	ID     int     `json:"id"`
	UserID int     `json:"user_id"`
	Total  float64 `json:"total"`
}
