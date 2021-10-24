package service

import (
	"ecommerce/api/repository"
	"ecommerce/models"
)

//ItemService ItemService struct
type ItemService struct {
	repository repository.ItemRepository
}

//NewItemService : returns the ItemService struct instance
func NewItemService(r repository.ItemRepository) ItemService {
	return ItemService{
		repository: r,
	}
}

//Save -> calls item repository save method
func (i ItemService) Save(item models.Item) error {
	return i.repository.Save(item)
}

//FindAll -> calls item repo find all method
func (i ItemService) FindAll(item models.Item, keyword string) (*[]models.Item, int64, error) {
	return i.repository.FindAll(item, keyword)
}

// Update -> calls itemrepo update method
func (i ItemService) Update(item models.Item) error {
	return i.repository.Update(item)
}

// Delete -> calls item repo delete method
func (i ItemService) Delete(id int64) error {
	var item models.Item
	item.ID = id
	return i.repository.Delete(item)
}

// Find -> calls item repo find method
func (i ItemService) Find(item models.Item) (models.Item, error) {
	return i.repository.Find(item)
}
