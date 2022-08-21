package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-courseware/courseware/api/internal/config"
	"go-zero-courseware/courseware/rpc/coursewareclient"
)

type ServiceContext struct {
	Config config.Config

	CoursewareRpc coursewareclient.Courseware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		CoursewareRpc: coursewareclient.NewCourseware(zrpc.MustNewClient(c.CoursewareRpc)),
	}
}
