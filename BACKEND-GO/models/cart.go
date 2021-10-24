package models

import "time"

// Cart Model
type Cart struct {
	ID          int64     `gorm:"primary_key;auto_increment" json:"id"`
	IsPurchased bool      `json:"isPurchased"`
	Items       []*Item   `gorm:"many2many:cart_item;"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	UserId      int64
}

// TableName method sets table name for Item model
func (cart *Cart) TableName() string {
	return "cart"
}

//ResponseMap -> response map method of Post
func (cart *Cart) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = cart.ID
	resp["items"] = cart.Items
	resp["created_at"] = cart.CreatedAt
	resp["updated_at"] = cart.UpdatedAt
	resp["is_purchased"] = cart.IsPurchased
	return resp
}
