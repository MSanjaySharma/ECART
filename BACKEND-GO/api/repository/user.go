package repository

import (
	"ecommerce/infrastructure"
	"ecommerce/models"
	"ecommerce/util"
)

//UserRepository -> UserRepository resposible for accessing database
type UserRepository struct {
	db infrastructure.Database
}

//NewUserRepository -> creates a instance on UserRepository
func NewUserRepository(db infrastructure.Database) UserRepository {
	return UserRepository{
		db: db,
	}
}

//CreateUser -> method for saving user to database
func (u UserRepository) CreateUser(user models.UserRegister) error {

	var dbUser models.User
	dbUser.Name = user.Name
	dbUser.Email = user.Email
	dbUser.Password = user.Password
	return u.db.DB.Create(&dbUser).Error
}

//LoginUser -> method for returning user
func (u UserRepository) LoginUser(user models.UserLogin) (*models.User, error) {

	var dbUser models.User
	email := user.Email
	password := user.Password

	err := u.db.DB.Where("email = ?", email).First(&dbUser).Error
	if err != nil {
		return nil, err
	}

	hashErr := util.CheckPasswordHash(password, dbUser.Password)
	if hashErr != nil {
		return nil, hashErr
	}
	return &dbUser, nil
}

//FindAll -> Method for fetching all users from database
func (i UserRepository) FindAll(user models.User, keyword string) (*[]models.User, int64, error) {
	var users []models.User
	var totalRows int64 = 0

	queryBuider := i.db.DB.Order("created_at desc").Model(&models.User{})

	// Search parameter
	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuider = queryBuider.Where(
			i.db.DB.Where("user.title LIKE ? ", queryKeyword))
	}

	err := queryBuider.
		Where(user).
		Find(&users).
		Count(&totalRows).Error
	return &users, totalRows, err
}
