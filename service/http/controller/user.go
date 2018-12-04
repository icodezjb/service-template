package controller

type UserController struct {
	baseController
}

func NewUserController() *UserController {
	return &UserController{}
}
