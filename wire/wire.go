//go:build wireinject
// +build wireinject

package wire

import (
	"zhiqu/provider"
)
import "github.com/google/wire"

// wire.go 初始化模块
func NewApp() (*provider.BeanFactory, error) {
	panic(wire.Build(provider.Set))
}
