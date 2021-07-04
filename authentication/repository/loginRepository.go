package repository

import (
	"github.com/architagr/golang-microservice-tutorial/authentication/models"
)

type Login struct{
	
}

var user []models.User

func Init() *Login{
	
	user = make([]models.User, 0,2)

	user = append(user, 
		models.User{
			Id: 1,
			Name: "Steve",
			UserName: "steve@yopmail.com",
			Password: "steve123",
			Roles: []int{1,2,3},
		}, 
		models.User{
			Id: 2,
		 	Name: "Mark",
		 	UserName: "mark@yopmail.com",
		 	Password: "mark@123",
		 	Roles: []int{4},
		},
	)

	return &Login{}
}

func (l *Login) GetUserByUserName(userName, password string) (models.User, *models.ErrorDetail){
	for _, value:= range user{
		if value.UserName == userName && value.Password == password{
			return value, nil
		} 
	}

	return models.User{}, &models.ErrorDetail{
		ErrorType: models.ErrorTypeError,
		ErrorMessage: "UserName and password not valid",

	}
}


