package impl

import (
	"context"
	"github.com/Andras5014/restful-api-demo/apps/host"
)

func (h HostServiceImpl) CreateHost(ctx context.Context, ins *host.Host) (*host.Host, error) {
	h.l.Debugf("create host: %s", ins.Name)
	return nil, nil
}

func (h HostServiceImpl) QueryHost(ctx context.Context, request *host.QueryHostRequest) (*host.Host, error) {
	//TODO implement me
	panic("implement me")
}

func (h HostServiceImpl) DescribeHost(ctx context.Context, request *host.QueryHostRequest) (*host.Host, error) {
	//TODO implement me
	panic("implement me")
}

func (h HostServiceImpl) UpdateHost(ctx context.Context, request *host.UpdateHostRequest) (*host.Host, error) {
	//TODO implement me
	panic("implement me")
}

func (h HostServiceImpl) DeleteHost(ctx context.Context, request *host.DeleteHostRequest) (*host.Host, error) {
	//TODO implement me
	panic("implement me")
}
