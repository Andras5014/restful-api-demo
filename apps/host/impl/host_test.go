package impl

import (
	"context"
	"github.com/Andras5014/restful-api-demo/apps/host"
	"github.com/infraboard/mcube/logger/zap"
	"testing"
)

var (
	service host.Service
)

func init() {
	// 初始化全局logger
	zap.DevelopmentSetup()
	service = NewHostServiceImpl()

}
func TestHostServiceImpl_CreateHost(t *testing.T) {
	ins := host.NewHost()
	service.CreateHost(context.Background(), ins)
}
