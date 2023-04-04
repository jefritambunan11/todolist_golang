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
	// TODO 
	// 1. var todos ialah pengguanan multi struct todo
	// 2. mengkonekan ke tabel todo   
	// 3. return multi struct 
	
	// TODO 1
	var todos []Todo
	
	// TODO 2
	var err = r.db.Find(&todos).Error
	if err != nil {
		return todos, err
	}

	// TODO 3
	return todos, nil
}

func (r *repository) FindByUserID(UserID int) ([]Todo, error) {
	// TODO
	// 1. var todos ialah pengguanan multi struct todo
	// 2. mengkonekan ke tabel todo dan load jg tabel user
	// 3. return multi struct todo
	
	
	// TODO 1
	var todos []Todo

	// TODO 2
	var err = r.db.Preload("User").Where("user_id = ?", UserID).Find(&todos).Error
	if err != nil {
		return todos, err
	}
	
	// TODO 3
	return todos, nil
}

func (r *repository) FindByID(ID int) (Todo, error) {	
	// TODO
	// 1. penggunaan struct todo ke var todo
	// 2. mengkonekan ke tabel todo dan load jg tabel user
	// 3. return struct todo
	
	// TODO 1
	var todo Todo

	// TODO 2
	var err = r.db.Preload("User").Where("id = ?", ID).Find(&todo).Error
	if err != nil {
		return todo, err
	}
	
	// TODO 3
	return todo, nil
}

func (r *repository) Save(todo Todo) (Todo, error) {
	// TODO 
	// 1. mengakses dan insert ke tabel todo
	// 2. return struct todo 
	
	// TODO 1
	var err = r.db.Create(&todo).Error
	if err != nil {
		return todo, err
	}
	
	// TODO 2
	return todo, nil
}

func (r *repository) Update(todo Todo) (Todo, error) {
	// TODO
	// 1. mengakses dan update ke tabel todo
	// 2. return struct todo
	
	// TODO 1
	var err = r.db.Save(&todo).Error
	if err != nil {
		return todo, err
	}
	
	// TODO 2
	return todo, nil
}

func (r *repository) Delete(todo Todo) (Todo, error) {
	// TODO 
	// 1. mengakses dan delete ke tabel todo
	// 2. return struct todo
	
	// TODO 1
	var err = r.db.Delete(&todo).Error	
	if err != nil {
		return todo, err
	}
	
	// TODO 2
	return todo, nil
}
