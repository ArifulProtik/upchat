package data

import "ArifulProtik/UpChat/internal/ent"

type UserCreateBody struct {
	Name     string `json:"name,omitempty"     binding:"required"`
	Email    string `json:"email,omitempty"    binding:"required,email"`
	Password string `json:"password,omitempty" binding:"required,min=6"`
}

type LoginBody struct {
	Email    string `json:"email,omitempty"    binding:"required,email"`
	Password string `json:"password,omitempty" binding:"required,min=6"`
}

type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

type LoginData struct {
	Token string
	User  *ent.User
}

type UserResponse struct {
	BaseModel
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AccountResponse struct {
	Provider     string `json:"provider"`
	MailVerified bool   `json:"mail_verified"`
}

type GetSessionResponse struct {
	User    *UserResponse    `json:"user"`
	Account *AccountResponse `json:"account"`
}
