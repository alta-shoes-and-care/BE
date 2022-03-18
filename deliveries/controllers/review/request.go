package review

import (
	R "final-project/entities/review"
)

type RequestInsertReview struct {
	ServiceID uint   `json:"service_id"`
	OrderID   uint   `json:"order_id"`
	Rating    int    `json:"rating"`
	Review    string `json:"review"`
}

func (Req RequestInsertReview) ToEntityReview(userID uint) R.Reviews {
	return R.Reviews{
		ServiceID: Req.ServiceID,
		OrderID:   Req.OrderID,
		Rating:    Req.Rating,
		Review:    Req.Review,
		UserID:    userID,
	}
}
