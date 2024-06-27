package models

type User struct {
	ID        uint
	Phone     string
	Email     string
	Password  string
	FirstName string
	LastName  string
}

type LoginUserReq struct {
	PhoneOrEmail string
	Password     string
}

type LogoutUserReq struct {
	// TODO
}
