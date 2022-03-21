package review

import (
	"errors"
	R "final-project/entities/review"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type ReviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) *ReviewRepository {
	return &ReviewRepository{
		db: db,
	}
}

func (repo *ReviewRepository) Insert(newReview R.Reviews) (FormatReview, error) {
	var checkReview R.Reviews
	var review FormatReview
	if rowsAffected := repo.db.Where("order_id = ? AND user_id = ?", newReview.OrderID, newReview.UserID).
		First(&checkReview).RowsAffected; rowsAffected > 0 {
		return FormatReview{}, errors.New("review sudah ada")
	}

	if err := repo.db.Create(&newReview).Error; err != nil {
		log.Warn(err)
		return FormatReview{}, errors.New("gagal membuat review baru")
	}

	repo.db.Table("reviews as r").
		Select("r.id as ID, r.user_id as UserID, r.service_id as ServiceID, r.order_id as OrderID, u.name as Name, r.rating as Rating, r.review as Review").
		Joins("inner join users as u on r.user_id = u.id").
		Where("r.user_id = ? AND r.service_id = ? AND r.order_id = ?", newReview.UserID, newReview.ServiceID, newReview.OrderID).
		Last(&review)
	return review, nil
}

func (repo *ReviewRepository) Get() ([]FormatReview, error) {
	var reviews []FormatReview

	if rowsAffected := repo.db.Table("reviews as r").
		Select("r.id as ID, r.user_id as UserID, r.service_id as ServiceID, r.order_id as OrderID, u.name as Name, r.rating as Rating, r.review as Review").
		Joins("inner join users as u on r.user_id = u.id").
		Find(&reviews).RowsAffected; rowsAffected == 0 {
		return nil, errors.New("tidak terdapat review sama sekali")
	}
	return reviews, nil
}
