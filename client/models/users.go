package models

type User struct {
	Phone    string
	Email    string
	Password string
}

type LoginUserReq struct {
	PhoneOrEmail string
	Password     string
}

type LogoutUserReq struct {
	// TODO
}
