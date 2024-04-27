package provider

import (
	"github.com/beego/beego/v2/core/logs"
	"sync"
	"zhiqu/app/category"
	"zhiqu/app/login"
	"zhiqu/app/user"
	"zhiqu/controllers"
)

var (
	appBeanFactory *BeanFactory
	once           sync.Once
	BeeLog         *logs.BeeLogger
)

type BeanFactory struct {
	CategoryController     *category.Controller
	CollectVideoController *controllers.CollectVideoController
	LoginController        *login.Controller
	UserController         *user.Controller
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
	once.Do(func() {
		appBeanFactory = &BeanFactory{
			LoginController:    loginController,
			CategoryController: categoryController,
			UserController:     userController,
		}
	})
	return appBeanFactory
}
