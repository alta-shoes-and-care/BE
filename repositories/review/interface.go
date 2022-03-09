package review

import R "final-project/entities/review"

type Review interface {
	Insert(newReview R.Reviews) (FormatReview, error)
	Get() ([]FormatReview, error)
	Update(reviewUpdate R.Reviews) (FormatReview, error)
	Delete(ID uint) error
}
