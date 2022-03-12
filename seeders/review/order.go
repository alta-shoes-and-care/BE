package review

import (
	"final-project/entities/review"
)

func ReviewSeeder() review.Reviews {
	return review.Reviews{
		ServiceID: 1,
		OrderID:   1,
		Rating:    5,
		Review:    "Mantap",
		UserID:    1,
	}
}
