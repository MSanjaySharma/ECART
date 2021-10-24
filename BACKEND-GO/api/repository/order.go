package repository

import (
	"ecommerce/infrastructure"
	"ecommerce/models"
)

//OrderRepository -> OrderRepository
type OrderRepository struct {
	db infrastructure.Database
}

// NewOrderRepository : fetching database
func NewOrderRepository(db infrastructure.Database) OrderRepository {
	return OrderRepository{
		db: db,
	}
}

//Save -> Method for saving order to database
func (o OrderRepository) Save(order models.Order) error {
	return o.db.DB.Create(&order).Error
}

//FindAll -> Method for fetching all orders from database
func (o OrderRepository) FindAll(order models.Order, keyword string) (*[]models.Order, int64, error) {
	var orders []models.Order
	var totalRows int64 = 0

	queryBuider := o.db.DB.Order("created_at desc").Model(&models.Order{})

	// Search parameter
	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuider = queryBuider.Where(
			o.db.DB.Where("order.title LIKE ? ", queryKeyword))
	}

	err := queryBuider.
		Where(order).
		Find(&orders).
		Count(&totalRows).Error
	return &orders, totalRows, err
}

//Update -> Method for updating Order
func (o OrderRepository) Update(order models.Order) error {
	return o.db.DB.Save(&order).Error
}

//Find -> Method for fetching order by id
func (o OrderRepository) Find(order models.Order) (models.Order, error) {
	var orders models.Order
	err := o.db.DB.
		Debug().
		Model(&models.Order{}).
		Where(&order).
		Take(&orders).Error
	return orders, err
}

//Delete Deletes Order
func (o OrderRepository) Delete(order models.Order) error {
	return o.db.DB.Delete(&order).Error
}
