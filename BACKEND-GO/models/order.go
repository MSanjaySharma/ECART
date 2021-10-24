package models

import "time"

// Order Model
type Order struct {
	ID        int64 `gorm:"primary_key;auto_increment" json:"id"`
	CartID    int64
	Cart      Cart
	UserID    int64
	User      User
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// TableName method sets table name for Item model
func (order *Order) TableName() string {
	return "order"
}

//ResponseMap -> response map method of Post
func (order *Order) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = order.ID
	resp["created_at"] = order.CreatedAt
	resp["updated_at"] = order.UpdatedAt
	return resp
}
