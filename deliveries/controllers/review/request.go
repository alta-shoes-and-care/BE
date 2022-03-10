package review

import (
	R "final-project/entities/review"

	"gorm.io/gorm"
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

type RequestUpdateReview struct {
	Rating int    `json:"rating"`
	Review string `json:"review"`
}

func (Req RequestUpdateReview) ToEntityReview(ID, UserID uint) R.Reviews {
	return R.Reviews{
		Model:  gorm.Model{ID: ID},
		UserID: UserID,
		Rating: Req.Rating,
		Review: Req.Review,
	}
}
