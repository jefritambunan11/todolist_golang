package todo

import (

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Todo, error)
	FindByUserID(UserID int, page_number int) ([]Todo, error)
	FindByID(ID int, UserID int) (Todo, error)
	CountRowUserID(UserID int) (int64, error)
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
	
	var todos []Todo	

	var err = r.db.Find(&todos).Error
	if err != nil { return todos, err }
	
	return todos, nil
}

func (r *repository) FindByUserID(UserID int, page_number int) ([]Todo, error) {

	var todos []Todo

	if page_number > 0 {
		var page = page_number
		var pageSize = 5
		var offset = (page - 1) * pageSize
		var err = r.db.Offset(offset).Limit(pageSize).Where("user_id = ?", UserID).Find(&todos).Error
		if err != nil { return todos, err }	
	}else{
		var err = r.db.Where("user_id = ?", UserID).Find(&todos).Error
		if err != nil { return todos, err }	
	}

	return todos, nil
}

func (r *repository) FindByID(ID int, UserID int) (Todo, error) {	
	
	var todo Todo
	
	var err = r.db.Where("id = ? and user_id = ?", ID, UserID).Find(&todo).Error
	if err != nil { return todo, err }
	
	return todo, nil
}


func (r *repository) CountRowUserID(UserID int) (int64, error) {

	var todo Todo
	var total_data int64

	var err = r.db.Where("user_id = ?", UserID).Find(&todo).Count(&total_data).Error
	if err != nil { return 0, err }	

	return total_data, nil
}


func (r *repository) Save(todo Todo) (Todo, error) {
	
	var err = r.db.Create(&todo).Error
	if err != nil { return todo, err }
	
	return todo, nil
}

func (r *repository) Update(todo Todo) (Todo, error) {
	
	var err = r.db.Save(&todo).Error
	if err != nil { return todo, err }	
	
	return todo, nil
}

func (r *repository) Delete(todo Todo) (Todo, error) {
	
	var err = r.db.Delete(&todo).Error	
	if err != nil { return todo, err }
	
	return todo, nil
}
