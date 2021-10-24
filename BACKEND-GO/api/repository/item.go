package repository

import (
	"ecommerce/infrastructure"
	"ecommerce/models"
)

//ItemRepository -> ItemRepository
type ItemRepository struct {
	db infrastructure.Database
}

// NewItemRepository : fetching database
func NewItemRepository(db infrastructure.Database) ItemRepository {
	return ItemRepository{
		db: db,
	}
}

//Save -> Method for saving item to database
func (i ItemRepository) Save(item models.Item) error {
	return i.db.DB.Create(&item).Error
}

//FindAll -> Method for fetching all items from database
func (i ItemRepository) FindAll(item models.Item, keyword string) (*[]models.Item, int64, error) {
	var items []models.Item
	var totalRows int64 = 0

	queryBuider := i.db.DB.Order("created_at desc").Model(&models.Item{})

	// Search parameter
	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuider = queryBuider.Where(
			i.db.DB.Where("item.title LIKE ? ", queryKeyword))
	}

	err := queryBuider.
		Where(item).
		Find(&items).
		Count(&totalRows).Error
	return &items, totalRows, err
}

//Update -> Method for updating Item
func (i ItemRepository) Update(item models.Item) error {
	return i.db.DB.Save(&item).Error
}

//Find -> Method for fetching item by id
func (i ItemRepository) Find(item models.Item) (models.Item, error) {
	var items models.Item
	err := i.db.DB.
		Debug().
		Model(&models.Item{}).
		Where(&item).
		Take(&items).Error
	return items, err
}

//Delete Deletes Item
func (i ItemRepository) Delete(item models.Item) error {
	return i.db.DB.Delete(&item).Error
}
