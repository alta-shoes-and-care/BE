package service

import S "final-project/entities/service"

type Service interface {
	Create(newService S.Services) (S.Services, error)
	Get() ([]S.Services, error)
	GetDetails(ID uint) (S.Services, error)
	Update(serviceUpdate S.Services) (S.Services, error)
	UpdateImage(ID uint, image string) (S.Services, error)
	Delete(ID uint) error
}
