package service

import (
	"ecommerce/api/repository"
	"ecommerce/models"
)

//OrderService OrderService struct
type OrderService struct {
	repository repository.OrderRepository
}

//NewOrderService : returns the OrderService struct instance
func NewOrderService(r repository.OrderRepository) OrderService {
	return OrderService{
		repository: r,
	}
}

//Save -> calls order repository save method
func (o OrderService) Save(order models.Order) error {
	return o.repository.Save(order)
}

//FindAll -> calls order repo find all method
func (o OrderService) FindAll(order models.Order, keyword string) (*[]models.Order, int64, error) {
	return o.repository.FindAll(order, keyword)
}

// Update -> calls orderrepo update method
func (o OrderService) Update(order models.Order) error {
	return o.repository.Update(order)
}

// Delete -> calls order repo delete method
func (o OrderService) Delete(id int64) error {
	var order models.Order
	order.ID = id
	return o.repository.Delete(order)
}

// Find -> calls order repo find method
func (o OrderService) Find(order models.Order) (models.Order, error) {
	return o.repository.Find(order)
}
