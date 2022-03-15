package review

import (
	R "final-project/entities/review"
)

func ReviewSeeder() R.Reviews {
	return R.Reviews{
		ServiceID: 1,
		OrderID:   1,
		Rating:    5,
		Review:    "Mantap",
		UserID:    1,
	}
}
