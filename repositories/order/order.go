package order

import (
	"errors"
	O "final-project/entities/order"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (repo *OrderRepository) Create(newOrder O.Orders) (ResponseOrder, error) {
	var order ResponseOrder

	if err := repo.db.Create(&newOrder).Error; err != nil {
		log.Warn(err)
		return ResponseOrder{}, errors.New("gagal membuat order baru")
	}

	repo.db.Table("orders as o").
		Select("o.id as ID, o.user_id as UserID, o.service_id as ServiceID, s.title as ServiceTitle, s.price as Price, o.qty as Qty, pm.id as PaymentMethodID, pm.name as PaymentMethodName, o.date as Date, o.address as Address, o.city as City, o.phone as Phone, o.status as Status, o.is_paid as IsPaid, o.url as Url").
		Joins("inner join services as s on s.id = o.service_id").
		Joins("inner join payment_methods as pm on pm.id = o.payment_method_id").
		Where("o.user_id = ? AND o.service_id = ? AND o.qty = ?", newOrder.UserID, newOrder.ServiceID, newOrder.Qty).
		Last(&order)
	return order, nil
}

func (repo *OrderRepository) Get() ([]ResponseOrder, error) {
	var orders []ResponseOrder

	if rowsAffected := repo.db.Table("orders as o").Select("o.id as ID, o.user_id as UserID, o.service_id as ServiceID, s.title as ServiceTitle, s.price as Price, o.qty as Qty, pm.id as PaymentMethodID, pm.name as PaymentMethodName, o.date as Date, o.address as Address, o.city as City, o.phone as Phone, o.status as Status, o.is_paid as IsPaid, o.url as Url").Joins("inner join services as s on s.id = o.service_id").Joins("inner join payment_methods as pm on pm.id = o.payment_method_id").Find(&orders).RowsAffected; rowsAffected == 0 {
		return nil, errors.New("tidak terdapat order sama sekali")
	}
	return orders, nil
}

func (repo *OrderRepository) GetByUserID(UserID uint) ([]ResponseOrder, error) {
	var orders []ResponseOrder

	if rowsAffected := repo.db.Table("orders as o").Select("o.id as ID, o.user_id as UserID, o.service_id as ServiceID, s.title as ServiceTitle, s.price as Price, o.qty as Qty, pm.id as PaymentMethodID, pm.name as PaymentMethodName, o.date as Date, o.address as Address, o.city as City, o.phone as Phone, o.status as Status, o.is_paid as IsPaid, o.url as Url").Joins("inner join services as s on s.id = o.service_id").Joins("inner join payment_methods as pm on pm.id = o.payment_method_id").Where("o.user_id = ?", UserID).Find(&orders).RowsAffected; rowsAffected == 0 {
		return nil, errors.New("tidak terdapat order sama sekali")
	}
	return orders, nil
}

func (repo *OrderRepository) GetByID(ID uint) (ResponseOrder, error) {
	var order ResponseOrder

	if rowsAffected := repo.db.Table("orders as o").Select("o.id as ID, o.user_id as UserID, o.service_id as ServiceID, s.title as ServiceTitle, s.price as Price, o.qty as Qty, pm.id as PaymentMethodID, pm.name as PaymentMethodName, o.date as Date, o.address as Address, o.city as City, o.phone as Phone, o.status as Status, o.is_paid as IsPaid, o.url as Url").Joins("inner join services as s on s.id = o.service_id").Joins("inner join payment_methods as pm on pm.id = o.payment_method_id").Where("o.id = ?", ID).First(&order).RowsAffected; rowsAffected == 0 {
		return ResponseOrder{}, errors.New("gagal mendapatkan detail order")
	}
	return order, nil
}

func (repo *OrderRepository) InsertUrl(ID uint, url string) (ResponseOrder, error) {
	var order ResponseOrder

	if rowsAffected := repo.db.Table("orders").Where("id = ?", ID).Update("url", url).RowsAffected; rowsAffected == 0 {
		return ResponseOrder{}, errors.New("gagal menambahkan url pembayaran")
	}

	repo.db.Table("orders as o").Select("o.id as ID, o.user_id as UserID, o.service_id as ServiceID, s.title as ServiceTitle, s.price as Price, o.qty as Qty, pm.id as PaymentMethodID, pm.name as PaymentMethodName, o.date as Date, o.address as Address, o.city as City, o.phone as Phone, o.status as Status, o.is_paid as IsPaid, o.url as Url").
		Joins("inner join services as s on s.id = o.service_id").
		Joins("inner join payment_methods as pm on pm.id = o.payment_method_id").
		Where("o.id = ?", ID).First(&order)
	return order, nil
}

func (repo *OrderRepository) SetPaid(ID uint) (ResponseOrder, error) {
	var order ResponseOrder

	if rowsAffected := repo.db.Table("orders").Where("id = ?", ID).Update("is_paid", true).RowsAffected; rowsAffected == 0 {
		return ResponseOrder{}, errors.New("gagal mengubah status pembayaran menjadi paid")
	}

	repo.db.Table("orders as o").Select("o.id as ID, o.user_id as UserID, o.service_id as ServiceID, s.title as ServiceTitle, s.price as Price, o.qty as Qty, pm.id as PaymentMethodID, pm.name as PaymentMethodName, o.date as Date, o.address as Address, o.city as City, o.phone as Phone, o.status as Status, o.is_paid as IsPaid, o.url as Url").
		Joins("inner join services as s on s.id = o.service_id").
		Joins("inner join payment_methods as pm on pm.id = o.payment_method_id").
		Where("o.id = ?", ID).First(&order)
	return order, nil
}

func (repo *OrderRepository) SetAccepted(ID uint) (ResponseOrder, error) {
	var order ResponseOrder

	if rowsAffected := repo.db.Table("orders").Where("id = ?", ID).Update("status", "accepted").RowsAffected; rowsAffected == 0 {
		return ResponseOrder{}, errors.New("gagal mengubah status order menjadi accepted")
	}

	repo.db.Table("orders as o").Select("o.id as ID, o.user_id as UserID, o.service_id as ServiceID, s.title as ServiceTitle, s.price as Price, o.qty as Qty, pm.id as PaymentMethodID, pm.name as PaymentMethodName, o.date as Date, o.address as Address, o.city as City, o.phone as Phone, o.status as Status, o.is_paid as IsPaid, o.url as Url").
		Joins("inner join services as s on s.id = o.service_id").
		Joins("inner join payment_methods as pm on pm.id = o.payment_method_id").
		Where("o.id = ?", ID).First(&order)
	return order, nil
}

func (repo *OrderRepository) SetRejected(ID uint) (ResponseOrder, error) {
	var order ResponseOrder

	if rowsAffected := repo.db.Table("orders").Where("id = ?", ID).Update("status", "rejected").RowsAffected; rowsAffected == 0 {
		return ResponseOrder{}, errors.New("gagal mengubah status order menjadi rejected")
	}

	repo.db.Table("orders as o").Select("o.id as ID, o.user_id as UserID, o.service_id as ServiceID, s.title as ServiceTitle, s.price as Price, o.qty as Qty, pm.id as PaymentMethodID, pm.name as PaymentMethodName, o.date as Date, o.address as Address, o.city as City, o.phone as Phone, o.status as Status, o.is_paid as IsPaid, o.url as Url").
		Joins("inner join services as s on s.id = o.service_id").
		Joins("inner join payment_methods as pm on pm.id = o.payment_method_id").
		Where("o.id = ?", ID).First(&order)
	return order, nil
}

func (repo *OrderRepository) SetOnProcess(ID uint) (ResponseOrder, error) {
	var order ResponseOrder

	if rowsAffected := repo.db.Table("orders").Where("id = ?", ID).Update("status", "on process").RowsAffected; rowsAffected == 0 {
		return ResponseOrder{}, errors.New("gagal mengubah status order menjadi on process")
	}

	repo.db.Table("orders as o").Select("o.id as ID, o.user_id as UserID, o.service_id as ServiceID, s.title as ServiceTitle, s.price as Price, o.qty as Qty, pm.id as PaymentMethodID, pm.name as PaymentMethodName, o.date as Date, o.address as Address, o.city as City, o.phone as Phone, o.status as Status, o.is_paid as IsPaid, o.url as Url").
		Joins("inner join services as s on s.id = o.service_id").
		Joins("inner join payment_methods as pm on pm.id = o.payment_method_id").
		Where("o.id = ?", ID).First(&order)
	return order, nil
}

func (repo *OrderRepository) SetDelivering(ID uint) (ResponseOrder, error) {
	var order ResponseOrder

	if rowsAffected := repo.db.Table("orders").Where("id = ?", ID).Update("status", "delivering").RowsAffected; rowsAffected == 0 {
		return ResponseOrder{}, errors.New("gagal mengubah status order menjadi delivering")
	}

	repo.db.Table("orders as o").Select("o.id as ID, o.user_id as UserID, o.service_id as ServiceID, s.title as ServiceTitle, s.price as Price, o.qty as Qty, pm.id as PaymentMethodID, pm.name as PaymentMethodName, o.date as Date, o.address as Address, o.city as City, o.phone as Phone, o.status as Status, o.is_paid as IsPaid, o.url as Url").
		Joins("inner join services as s on s.id = o.service_id").
		Joins("inner join payment_methods as pm on pm.id = o.payment_method_id").
		Where("o.id = ?", ID).First(&order)
	return order, nil
}

func (repo *OrderRepository) SetCancel(ID uint) (ResponseOrder, error) {
	var order ResponseOrder

	if rowsAffected := repo.db.Table("orders").Where("id = ?", ID).Update("status", "cancel").RowsAffected; rowsAffected == 0 {
		return ResponseOrder{}, errors.New("gagal mengubah status order menjadi cancel")
	}

	repo.db.Table("orders as o").Select("o.id as ID, o.user_id as UserID, o.service_id as ServiceID, s.title as ServiceTitle, s.price as Price, o.qty as Qty, pm.id as PaymentMethodID, pm.name as PaymentMethodName, o.date as Date, o.address as Address, o.city as City, o.phone as Phone, o.status as Status, o.is_paid as IsPaid, o.url as Url").
		Joins("inner join services as s on s.id = o.service_id").
		Joins("inner join payment_methods as pm on pm.id = o.payment_method_id").
		Where("o.id = ?", ID).First(&order)
	return order, nil
}

func (repo *OrderRepository) SetDone(ID uint) (ResponseOrder, error) {
	var order ResponseOrder

	if rowsAffected := repo.db.Table("orders").Where("id = ?", ID).Update("status", "done").RowsAffected; rowsAffected == 0 {
		return ResponseOrder{}, errors.New("gagal mengubah status order menjadi done")
	}

	repo.db.Table("orders as o").Select("o.id as ID, o.user_id as UserID, o.service_id as ServiceID, s.title as ServiceTitle, s.price as Price, o.qty as Qty, pm.id as PaymentMethodID, pm.name as PaymentMethodName, o.date as Date, o.address as Address, o.city as City, o.phone as Phone, o.status as Status, o.is_paid as IsPaid, o.url as Url").
		Joins("inner join services as s on s.id = o.service_id").
		Joins("inner join payment_methods as pm on pm.id = o.payment_method_id").
		Where("o.id = ?", ID).First(&order)
	return order, nil
}
