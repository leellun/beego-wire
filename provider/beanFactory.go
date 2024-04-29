package provider

import (
	"github.com/beego/beego/v2/core/logs"
	"sync"
	"zhiqu/app"
)

var (
	AppBeanFactory *BeanFactory
	once           sync.Once
	BeeLog         *logs.BeeLogger
)

type BeanFactory struct {
	AppFxRouter *app.FxRouter
}

func init() {
	BeeLog = logs.NewLogger()
	err := BeeLog.SetLogger(logs.AdapterConsole)
	if err != nil {
		return
	}
}
func NewBeanFactory(appFxRouter *app.FxRouter) *BeanFactory {
	once.Do(func() {
		AppBeanFactory = &BeanFactory{
			AppFxRouter: appFxRouter,
		}
	})
	return AppBeanFactory
}
