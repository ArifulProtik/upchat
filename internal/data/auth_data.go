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
	Token string    `json:"token"`
	User  *ent.User `json:"user"`
}
