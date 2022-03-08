package service

import (
	"errors"
	S "final-project/entities/service"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type ServiceRepository struct {
	db *gorm.DB
}

func NewServiceRepository(db *gorm.DB) *ServiceRepository {
	return &ServiceRepository{
		db: db,
	}
}

func (repo *ServiceRepository) Create(newService S.Services) (S.Services, error) {
	if err := repo.db.Create(&newService).Error; err != nil {
		log.Warn(err)
		return S.Services{}, errors.New("gagal membuat service baru")
	}
	return newService, nil
}

func (repo *ServiceRepository) Get() ([]S.Services, error) {
	var services []S.Services

	if err := repo.db.Find(&services).Error; err != nil {
		log.Warn(err)
		return nil, errors.New("tidak terdapat service sama sekali")
	}
	return services, nil
}

func (repo *ServiceRepository) GetDetail(ID uint) (S.Services, error) {
	var service S.Services

	if err := repo.db.First(&service, ID).Error; err != nil {
		log.Warn(err)
		return S.Services{}, errors.New("gagal mendapatkan detail service")
	}
	return service, nil
}

func (repo *ServiceRepository) Update(serviceUpdate S.Services) (S.Services, error) {
	if rowsAffected := repo.db.Model(&serviceUpdate).Updates(serviceUpdate).Error; rowsAffected != nil {
		return S.Services{}, errors.New("tidak ada data service yang diperbarui")
	}

	repo.db.First(&serviceUpdate)

	return serviceUpdate, nil
}

func (repo *ServiceRepository) Delete(ID uint) error {
	if rowsAffected := repo.db.Delete(&S.Services{}, ID).RowsAffected; rowsAffected == 0 {
		return errors.New("tidak ada service yang dihapus")
	}
	return nil
}
