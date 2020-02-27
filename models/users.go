package models

import "github.com/samratsuzil/api/database"

//GetAllUsers fetch all user
func GetAllUsers(book *[]User) (err error) {
	if err = database.ConnectDB().Find(book).Error; err != nil {
		return err
	}
	return nil
}

//GetUser Endpoint
func GetUser(book *User, id uint) (err error) {
	if err = database.ConnectDB().First(&book, "id=?", id).Error; err != nil {
		return err
	}
	return nil
}

//AddNewUser endpoint hit
func AddNewUser(book *User) (err error) {
	if err = database.ConnectDB().Create(&book).Error; err != nil {
		return err
	}
	return nil
}

//UpdateUser endpoint models
func UpdateUser(book *User, id uint) (err error) {

	if err = database.ConnectDB().Model(&book).Where("id=?", id).Updates(&book).Error; err != nil {
		return err
	}
	return nil

}

//DeleteUser Endpoint
func DeleteUser(id uint) (err error) {
	if err = database.ConnectDB().Where("id = ?", id).Delete(&User{}).Error; err != nil {
		return err
	}
	return nil
}
