package service

import (
	"errors"
	S "final-project/entities/service"
	"mime/multipart"

	"gorm.io/gorm"
)

type MockServiceTrueRepository struct{}

func (repo *MockServiceTrueRepository) Create(newService S.Services) (S.Services, error) {
	return S.Services{Model: gorm.Model{ID: 1}, Title: "Service 1", Description: "Layanan 1", Price: 30000, Image: "https://foo.com/bar.jpeg", UserID: 1}, nil
}

func (repo *MockServiceTrueRepository) Get() ([]S.Services, error) {
	service1 := S.Services{
		Model:       gorm.Model{ID: 1},
		Title:       "Service 1",
		Description: "Layanan 1",
		Price:       30000,
		Image:       "https://foo.com/bar.jpeg",
		UserID:      1}
	service2 := S.Services{
		Model:       gorm.Model{ID: 2},
		Title:       "Service 2",
		Description: "Layanan 2",
		Price:       20000,
		Image:       "https://foo.com/bar2.jpeg",
		UserID:      1}

	return []S.Services{service1, service2}, nil
}

func (repo *MockServiceTrueRepository) GetDetails(ID uint) (S.Services, error) {
	return S.Services{Model: gorm.Model{ID: 1}, Title: "Service 1", Description: "Layanan 1", Price: 30000, Image: "https://foo.com/bar.jpeg", UserID: 1}, nil
}

func (repo *MockServiceTrueRepository) Update(serviceUpdate S.Services) (S.Services, error) {
	return S.Services{Model: gorm.Model{ID: 1}, Title: "Service 1 Updated", Description: "Layanan 1", Price: 30000, Image: "https://foo.com/bar.jpeg", UserID: 1}, nil
}

func (repo *MockServiceTrueRepository) Delete(ID uint) error {
	return nil
}

type MockServiceFalseRepository struct{}

func (repo *MockServiceFalseRepository) Create(newService S.Services) (S.Services, error) {
	return S.Services{}, errors.New("create error")
}

func (repo *MockServiceFalseRepository) Get() ([]S.Services, error) {
	return nil, errors.New("get error")
}

func (repo *MockServiceFalseRepository) GetDetails(ID uint) (S.Services, error) {
	return S.Services{}, errors.New("get details error")
}

func (repo *MockServiceFalseRepository) Update(serviceUpdate S.Services) (S.Services, error) {
	return S.Services{}, errors.New("update error")
}

func (repo *MockServiceFalseRepository) Delete(ID uint) error {
	return errors.New("delete error")
}

type MockAWSStructTrue struct{}

func (awsStruct *MockAWSStructTrue) DoUpload(region, bucket string, file *multipart.FileHeader) (string, error) {
	return "https://foo.com/bar.jpeg", nil
}

func (awsStruct *MockAWSStructTrue) DoDelete(fileName, bucket string) error {
	return nil
}

type MockAWSStructFalse struct{}

func (awsStruct *MockAWSStructFalse) DoUpload(region, bucket string, file *multipart.FileHeader) (string, error) {
	return "https://foo.com/bar.jpeg", nil
}

func (awsStruct *MockAWSStructFalse) DoDelete(fileName, bucket string) error {
	return nil
}
