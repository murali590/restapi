package service

import (
	"fmt"
	"log"

	"github.com/murali590/restapi/config"
)

//GetAllusers - fetch all users
func GetAllusers(user *[]model.user) (err error) {
	if err = config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

//Createuser - creates an user
func Createuser(user *model.user) (err error) {
	log.Println("Album Input:", user)
	if err = config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

//Getuser fetch one  user
func Getuser(user *model.user, id string) (err error) {
	if err := config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

//Updateuser - modifies an user
func Updateuser(user *model.user, id string) (err error) {
	fmt.Println(user)
	config.DB.Save(user)
	return nil
}

//Deleteuser deletes a given user
func Deleteuser(user *model.user, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(user)
	return nil
}
