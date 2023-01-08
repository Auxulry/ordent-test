package entity

type Cart struct {
	ID            int `json:"id"`
	CartSessionID int `json:"cart_session_id"`
	ProductID     int `json:"product_id"`
	Quantity      int `json:"quantity"`
}
