package todo

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Todo, error)
	FindByUserID(UserID int) ([]Todo, error)
	FindByID(ID int) (Todo, error)
	Save(todo Todo) (Todo, error)
	Update(todo Todo) (Todo, error)
	Delete(todo Todo) (Todo, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Todo, error) {
	var todo []Todo

	err := r.db.Find(&todo).Error

	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (r *repository) FindByUserID(UserID int) ([]Todo, error) {
	var todo []Todo

	err := r.db.Where("user_id = ?", UserID).Preload("User").Find(&todo).Error
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (r *repository) FindByID(ID int) (Todo, error) {
	var todo Todo

	err := r.db.Where("id = ?", ID).Preload("User").Find(&todo).Error
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (r *repository) Save(todo Todo) (Todo, error) {
	err := r.db.Create(&todo).Error
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (r *repository) Update(todo Todo) (Todo, error) {
	err := r.db.Save(&todo).Error
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (r *repository) Delete(todo Todo) (Todo, error) {
	err := r.db.Delete(&todo).Error
	if err != nil {
		return todo, err
	}
	return todo, nil
}
