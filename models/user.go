package models

type UserBase struct {
	Email 		string 	`json:"email"`
	Username 	string 	`json:"username"`
}

type UserSignUp struct {
	UserBase
	Password 	string	`json:"password"`
	Name 		string 	`json:"name"`
	LastName 	string	`json:"lastName"`
}

type UserLogin struct {
	UserBase
	Password 	string	`json:"password"`
}

type User struct {
	UserBase
	Name 		string 	`json:"name"`
	LastName 	string	`json:"lastName"`	
}