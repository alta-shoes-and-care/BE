package review

import R "final-project/entities/review"

type Review interface {
	Insert(newReview R.Reviews) (FormatReview, error)
	Get() ([]FormatReview, error)
}
