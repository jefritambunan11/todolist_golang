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
<<<<<<< HEAD
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
=======
	
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
	
>>>>>>> 9ede165dd324e1863802b8cdb43c54dc29b7457e
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
<<<<<<< HEAD
	// TODO 
	// 1. mengakses dan insert ke tabel todo
	// 2. return struct todo 
	
	// TODO 1
	var err = r.db.Create(&todo).Error
	if err != nil {
		return todo, err
	}
	
	// TODO 2
=======
	
	var err = r.db.Create(&todo).Error
	if err != nil { return todo, err }
	
>>>>>>> 9ede165dd324e1863802b8cdb43c54dc29b7457e
	return todo, nil
}

func (r *repository) Update(todo Todo) (Todo, error) {
<<<<<<< HEAD
	// TODO
	// 1. mengakses dan update ke tabel todo
	// 2. return struct todo
	
	// TODO 1
	var err = r.db.Save(&todo).Error
	if err != nil {
		return todo, err
	}
	
	// TODO 2
=======
	
	var err = r.db.Save(&todo).Error
	if err != nil { return todo, err }	
	
>>>>>>> 9ede165dd324e1863802b8cdb43c54dc29b7457e
	return todo, nil
}

func (r *repository) Delete(todo Todo) (Todo, error) {
<<<<<<< HEAD
	// TODO 
	// 1. mengakses dan delete ke tabel todo
	// 2. return struct todo
	
	// TODO 1
	var err = r.db.Delete(&todo).Error	
	if err != nil {
		return todo, err
	}
	
	// TODO 2
=======
	
	var err = r.db.Delete(&todo).Error	
	if err != nil { return todo, err }
	
>>>>>>> 9ede165dd324e1863802b8cdb43c54dc29b7457e
	return todo, nil
}
