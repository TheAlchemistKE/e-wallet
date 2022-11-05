package repository

import (
	"blog/infrastructure"
	"blog/models"
)

//UserRepository -> UserRepository
type UserRepository struct {
	db infrastructure.Database
}

// NewUserRepository : fetching database
func NewUserRepository(db infrastructure.Database) UserRepository {
	return UserRepository{
		db: db,
	}
}

//Save -> Method for saving user to database
func (u UserRepository) Save(user models.User) error {
	return u.db.DB.Create(&user).Error
}

//FindAll -> Method for fetching all users from database
func (u UserRepository) FindAll(user models.User, keyword string) (*[]models.User, int64, error) {
	var users []models.User
	var totalRows int64 = 0

	queryBuider := u.db.DB.Order("created_at desc").Model(&models.User{})

	// Search parameter
	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuider = queryBuider.Where(
			u.db.DB.Where("users.email LIKE ? ", queryKeyword))
	}

	err := queryBuider.
		Where(user).
		Find(&users).
		Count(&totalRows).Error
	return &users, totalRows, err
}

//Update -> Method for updating user
func (u UserRepository) Update(user models.User) error {
	return u.db.DB.Save(&user).Error
}

//Find -> Method for fetching user by id
func (u UserRepository) Find(user models.User) (models.User, error) {
	var users models.User
	err := u.db.DB.
		Debug().
		Model(&models.User{}).
		Where(&user).
		Take(&users).Error
	return users, err
}

//Delete Deletes User
func (u UserRepository) Delete(user models.User) error {
	return u.db.DB.Delete(&user).Error
}
