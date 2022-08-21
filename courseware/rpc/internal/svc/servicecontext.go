package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-courseware/courseware/rpc/internal/config"
	"go-zero-courseware/courseware/rpc/model"
)

type ServiceContext struct {
	Config config.Config

	CoursewareModel model.CoursewareModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:          c,
		CoursewareModel: model.NewCoursewareModel(conn, c.CacheRedis),
	}
}
