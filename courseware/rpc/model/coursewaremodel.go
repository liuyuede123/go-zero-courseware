package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CoursewareModel = (*customCoursewareModel)(nil)

type (
	// CoursewareModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCoursewareModel.
	CoursewareModel interface {
		coursewareModel
	}

	customCoursewareModel struct {
		*defaultCoursewareModel
	}
)

// NewCoursewareModel returns a model for the database table.
func NewCoursewareModel(conn sqlx.SqlConn, c cache.CacheConf) CoursewareModel {
	return &customCoursewareModel{
		defaultCoursewareModel: newCoursewareModel(conn, c),
	}
}
