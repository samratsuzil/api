package models

import "github.com/samratsuzil/api/database"

//GetAllPublications fetch all publications
func GetAllPublications(publication *[]Publication) (err error) {

	if err = database.ConnectDB().Preload("Book").Find(publication).Error; err != nil {
		return err
	}
	return nil
}

//GetPublication Endpoint
func GetPublication(publication *Publication, id uint) (err error) {
	if err = database.ConnectDB().First(&publication, "id=?", id).Error; err != nil {
		return err
	}
	return nil
}

//AddNewPublication endpoint hit
func AddNewPublication(publication *Publication) (err error) {
	if err = database.ConnectDB().Create(&publication).Error; err != nil {
		return err
	}
	return nil
}

//UpdatePublication endpoint models
func UpdatePublication(publication *Publication, id uint) (err error) {

	if err = database.ConnectDB().Model(&publication).Where("id=?", id).Updates(&publication).Error; err != nil {
		return err
	}
	return nil

}

//DeletePublication Endpoint
func DeletePublication(id uint) (err error) {
	if err = database.ConnectDB().Where("id = ?", id).Delete(&Publication{}).Error; err != nil {
		return err
	}
	return nil
}
