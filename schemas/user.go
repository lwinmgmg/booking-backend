package schemas

import "errors"

type UserCreate struct {
	FirstName       *string `json:"first_name"`
	LastName        *string `json:"last_name"`
	Email           *string
	Phone           *string
	Nrc             *string
	Address         *string
	Username        *string
	Password        *string
	ConfirmPassword *string
}

func (userCreate *UserCreate) Validate() error {
	if userCreate.FirstName == nil {
		tmpStr := ""
		userCreate.FirstName = &tmpStr
	}
	if userCreate.LastName == nil {
		tmpStr := ""
		userCreate.LastName = &tmpStr
	}
	if userCreate.Email == nil {
		tmpStr := ""
		userCreate.Email = &tmpStr
	}
	if userCreate.Phone == nil {
		tmpStr := ""
		userCreate.Phone = &tmpStr
	}
	if userCreate.Nrc == nil {
		tmpStr := ""
		userCreate.Nrc = &tmpStr
	}
	if userCreate.Address == nil {
		tmpStr := ""
		userCreate.Address = &tmpStr
	}
	if userCreate.Username == nil || *userCreate.Username == "" {
		return errors.New("username can't be empty")
	}
	if userCreate.Password == nil || *userCreate.Password == "" {
		return errors.New("password can't be empty")
	}
	return nil
}

type UserLogin struct {
	UserName *string
	Password *string
}

func (userLogin *UserLogin) Validate() error {
	if userLogin.UserName == nil {
		return errors.New("username can't be empty")
	}
	if userLogin.Password == nil {
		return errors.New("password can't be empty")
	}
	return nil
}
