package models

import "time"

//User -> User struct to save user on database
type User struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Token     string
	Cart      Cart
}

//TableName -> returns the table name of User Model
func (user *User) TableName() string {
	return "user"
}

//UserLogin -> Request Binding for User Login
type UserLogin struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

//UserRegister -> Request Binding for User Register
type UserRegister struct {
	Name     string `form:"name"`
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

//ResponseMap -> response map method of User
func (user *User) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = user.ID
	resp["name"] = user.Name
	resp["email"] = user.Email
	resp["token"] = user.Token
	resp["created_at"] = user.CreatedAt
	resp["updated_at"] = user.UpdatedAt
	return resp
}
