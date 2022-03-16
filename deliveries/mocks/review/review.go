package review

import (
	"errors"
	R "final-project/entities/review"
	"final-project/repositories/review"
	U "final-project/entities/user"

	"gorm.io/gorm"
)

type MockAuthUserRepository struct{}

func (repo *MockAuthUserRepository) Login(email, password string) (U.Users, error) {
	return U.Users{Model: gorm.Model{ID: 1}, Email: "ucup@ucup.com", Password: "ucup123", IsAdmin: false}, nil
}

type MockAuthAdminRepository struct{}

func (repo *MockAuthAdminRepository) Login(email, password string) (U.Users, error) {
	return U.Users{Model: gorm.Model{ID: 1}, Email: "ucup@ucup.com", Password: "ucup123", IsAdmin: true}, nil
}

type MockReviewRepository struct{}

func (repo *MockReviewRepository) Insert(newReview R.Reviews) (review.FormatReview, error) {
	return review.FormatReview{ID: 1, ServiceID: 1, OrderID: 1, UserID: 1, Name: "Yusuf", Rating: 5, Review: "Bagus"}, nil
}

func (repo *MockReviewRepository) Get() ([]review.FormatReview, error) {
	review1 := review.FormatReview{
		ID: 1,
		ServiceID: 1,
		OrderID: 1,
		UserID: 1,
		Name: "Yusuf",
		Rating: 5,
		Review: "Bagus",
	}

	review2 := review.FormatReview{
		ID: 2,
		ServiceID: 1,
		OrderID: 2,
		UserID: 2,
		Name: "Frans",
		Rating: 5,
		Review: "Bagus",
	}

	return []review.FormatReview{review1, review2}, nil
}

func (repo *MockReviewRepository) Update(reviewUpdate R.Reviews) (review.FormatReview, error) {
	return review.FormatReview{ID: 1, ServiceID: 1, OrderID: 1, UserID: 1, Name: "Yusuf", Rating: 5, Review: "Bagus"}, nil
}

func (repo *MockReviewRepository) Delete(ID, UserID uint) error {
	return nil
}

type MockFalseReviewRepository struct{}

func (repo *MockFalseReviewRepository) Insert(newReview R.Reviews) (review.FormatReview, error) {
	return review.FormatReview{}, errors.New("false insert review")
}

func (repo *MockFalseReviewRepository) Get() ([]review.FormatReview, error) {
	return nil, errors.New("false get all reviews")
}

func (repo *MockFalseReviewRepository) Update(reviewUpdate R.Reviews) (review.FormatReview, error) {
	return review.FormatReview{}, errors.New("false update review")
}

func (repo *MockFalseReviewRepository) Delete(ID, UserID uint) error {
	return errors.New("false delete review")
}

