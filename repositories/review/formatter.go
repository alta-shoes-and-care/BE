package review

type FormatReview struct {
	ID        uint   `json:"id"`
	ServiceID uint   `json:"service_id"`
	UserID    uint   `json:"user_id"`
	OrderID   uint   `json:"order_id"`
	Name      string `json:"name"`
	Rating    int    `json:"rating"`
	Review    string `json:"review"`
}
