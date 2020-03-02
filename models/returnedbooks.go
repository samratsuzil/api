package models

import "github.com/samratsuzil/api/database"

//GetAllReturnedBooks fetch all return
func GetAllReturnedBooks(returnedbook *[]ReturnedBook) (err error) {
	if err = database.ConnectDB().Find(returnedbook).Error; err != nil {
		return err
	}
	return nil
}

//GetReturnedBook Endpoint
func GetReturnedBook(returnedbook *ReturnedBook, id uint) (err error) {
	if err = database.ConnectDB().First(&returnedbook, "id=?", id).Error; err != nil {
		return err
	}
	return nil
}

//AddNewReturnedBook endpoint hit
func AddNewReturnedBook(returnedbook *ReturnedBook) (err error) {
	if err = database.ConnectDB().Create(&returnedbook).Error; err != nil {
		return err
	}
	return nil
}

//UpdateReturnedBook endpoint models
func UpdateReturnedBook(returnedbook *ReturnedBook, id uint) (err error) {

	if err = database.ConnectDB().Model(&returnedbook).Where("id=?", id).Updates(&returnedbook).Error; err != nil {
		return err
	}
	return nil

}

//DeleteReturnedBook Endpoint
func DeleteReturnedBook(id uint) (err error) {
	if err = database.ConnectDB().Where("id = ?", id).Delete(&ReturnedBook{}).Error; err != nil {
		return err
	}
	return nil
}
