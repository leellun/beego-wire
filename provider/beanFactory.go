package provider

import (
	"sync"
	"zhiqu/app/category"
	"zhiqu/app/login"
	"zhiqu/controllers"
)

var (
	appBeanFactory *BeanFactory
	once           sync.Once
)

type BeanFactory struct {
	CategoryController     *category.Controller
	CollectVideoController *controllers.CollectVideoController
	LoginController        *login.Controller
}

func NewBeanFactory(loginController *login.Controller,
	categoryController *category.Controller) *BeanFactory {
	once.Do(func() {
		appBeanFactory = &BeanFactory{
			LoginController:    loginController,
			CategoryController: categoryController,
		}
	})
	return appBeanFactory
}
