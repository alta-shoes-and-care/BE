package review

import (
	R "final-project/entities/review"
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

func (repo *MockReviewRepository) Insert(newReview R.Reviews) (FormatReview, error) {
	return FormatReview{ID: 1, ServiceID: 1, OrderID: 1, UserID: 1, Name: "Yusuf", Rating: 5, Review: "Bagus"}, nil
}

func (repo *MockReviewRepository) Get() ([]FormatReview, error) {
	review1 := FormatReview{
		ID: 1,
		ServiceID: 1,
		OrderID: 1,
		UserID: 1,
		Name: "Yusuf",
		Rating: 5,
		Review: "Bagus",
	}

	review2 := FormatReview{
		ID: 2,
		ServiceID: 1,
		OrderID: 2,
		UserID: 2,
		Name: "Frans",
		Rating: 5,
		Review: "Bagus",
	}

	return []FormatReview{review1, review2}, nil
}

func (repo *MockReviewRepository) Update(reviewUpdate R.Reviews) (FormatReview, error) {
	return FormatReview{ID: 1, ServiceID: 1, OrderID: 1, UserID: 1, Name: "Yusuf", Rating: 5, Review: "Bagus"}, nil
}

func (repo *MockReviewRepository) Delete(ID, UserID uint) error {
	return nil
}
