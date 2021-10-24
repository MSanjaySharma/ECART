package repository

import (
	"ecommerce/infrastructure"
	"ecommerce/models"
)

//CartRepository -> CartRepository
type CartRepository struct {
	db infrastructure.Database
}

// NewCartRepository : fetching database
func NewCartRepository(db infrastructure.Database) CartRepository {
	return CartRepository{
		db: db,
	}
}

//Save -> Method for saving cart to database
func (c CartRepository) Save(cart models.Cart) error {
	return c.db.DB.Create(&cart).Error
}

//FindAll -> Method for fetching all carts from database
func (c CartRepository) FindAll(cart models.Cart, keyword string) (*[]models.Cart, int64, error) {
	var carts []models.Cart
	var totalRows int64 = 0

	queryBuider := c.db.DB.Order("created_at desc").Model(&models.Cart{})

	// Search parameter
	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuider = queryBuider.Where(
			c.db.DB.Where("cart.title LIKE ? ", queryKeyword))
	}

	err := queryBuider.
		Where(cart).
		Find(&carts).
		Count(&totalRows).Error
	return &carts, totalRows, err
}

//Update -> Method for updating Cart
func (c CartRepository) Update(cart models.Cart) error {
	return c.db.DB.Save(&cart).Error
}

//Find -> Method for fetching cart by id
func (c CartRepository) Find(cart models.Cart) (models.Cart, error) {
	var carts models.Cart
	err := c.db.DB.
		Debug().
		Model(&models.Cart{}).
		Where(&cart).
		Take(&carts).Error
	return carts, err
}

//Delete Deletes Cart
func (c CartRepository) Delete(cart models.Cart) error {
	return c.db.DB.Delete(&cart).Error
}
