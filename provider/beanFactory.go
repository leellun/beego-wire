package provider

import (
	"github.com/beego/beego/v2/core/logs"
	"sync"
	"zhiqu/app/category"
	"zhiqu/app/login"
	"zhiqu/app/user"
)

var (
	AppBeanFactory *BeanFactory
	once           sync.Once
	BeeLog         *logs.BeeLogger
)

type BeanFactory struct {
	CategoryController *category.Controller
	LoginController    *login.Controller
	UserController     *user.Controller
}

func init() {
	BeeLog = logs.NewLogger()
	err := BeeLog.SetLogger(logs.AdapterConsole)
	if err != nil {
		return
	}
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
