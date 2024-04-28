package provider

import (
	"sync"
	"zhiqu/app/category"
	"zhiqu/app/login"
	"zhiqu/app/user"
)

var (
	AppBeanFactory *BeanFactory
	once           sync.Once
)

type BeanFactory struct {
	CategoryController *category.Controller
	LoginController    *login.Controller
	UserController     *user.Controller
}

func init() {
}
func NewBeanFactory(loginController *login.Controller,
	categoryController *category.Controller, userController *user.Controller) *BeanFactory {
	AppBeanFactory = &BeanFactory{
		LoginController:    loginController,
		CategoryController: categoryController,
		UserController:     userController,
	}
	return AppBeanFactory
}
