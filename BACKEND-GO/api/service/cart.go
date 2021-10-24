package service

import (
	"ecommerce/api/repository"
	"ecommerce/models"
)

//CartService CartService struct
type CartService struct {
	repository repository.CartRepository
}

//NewCartService : returns the CartService struct instance
func NewCartService(r repository.CartRepository) CartService {
	return CartService{
		repository: r,
	}
}

//Save -> calls cart repository save method
func (c CartService) Save(cart models.Cart) error {
	return c.repository.Save(cart)
}

//FindAll -> calls cart repo find all method
func (c CartService) FindAll(cart models.Cart, keyword string) (*[]models.Cart, int64, error) {
	return c.repository.FindAll(cart, keyword)
}

// Update -> calls cartrepo update method
func (c CartService) Update(cart models.Cart) error {
	return c.repository.Update(cart)
}

// Delete -> calls cart repo delete method
func (c CartService) Delete(id int64) error {
	var cart models.Cart
	cart.ID = id
	return c.repository.Delete(cart)
}

// Find -> calls cart repo find method
func (c CartService) Find(cart models.Cart) (models.Cart, error) {
	return c.repository.Find(cart)
}
