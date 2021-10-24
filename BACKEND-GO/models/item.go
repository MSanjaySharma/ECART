package models

import "time"

// Item Model
type Item struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:200" json:"title"`
	Carts     []*Cart   `gorm:"many2many:cart_item;"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// TableName method sets table name for Item model
func (item *Item) TableName() string {
	return "item"
}

//ResponseMap -> response map method of Post
func (item *Item) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = item.ID
	resp["name"] = item.Name
	resp["carts"] = item.Carts
	resp["created_at"] = item.CreatedAt
	resp["updated_at"] = item.UpdatedAt
	return resp
}
