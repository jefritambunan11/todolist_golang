package user

import "gorm.io/gorm"

type Repository interface {
	FindByEmail(email string) (User, error)
	FindByID(id int) (User, error)
	Save(user User) (User, error)
	Update(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	var err = r.db.Where("email = ?", email).Find(&user).Error

	if err != nil { return user, err }

	return user, nil
}

func (r *repository) FindByID(id int) (User, error) {
	var user User
	var err = r.db.Where("id = ?", id).Find(&user).Error

	if err != nil { return user, err }

	return user, nil
}

func (r *repository) Save(user User) (User, error) {
	var err = r.db.Create(&user).Error

	if err != nil { return user, err }

	return user, nil
}

func (r *repository) Update(user User) (User, error) {
	var err = r.db.Save(&user).Error

	if err != nil { return user, err }

	return user, nil
}
