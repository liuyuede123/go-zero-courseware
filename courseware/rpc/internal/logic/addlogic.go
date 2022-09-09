package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-courseware/courseware/common/xerr"
	"go-zero-courseware/courseware/rpc/courseware"
	"go-zero-courseware/courseware/rpc/internal/svc"
	"go-zero-courseware/courseware/rpc/model"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *courseware.AddRequest) (*courseware.AddResponse, error) {
	_, err := l.svcCtx.CoursewareModel.FindOneByCode(l.ctx, in.Code)
	if err == nil {
		return nil, xerr.NewErrCodeMsg(5000, "课件编号已存在")
	}

	if err == model.ErrNotFound {
		newCourseware := model.Courseware{
			Code: in.Code,
			Name: in.Name,
			Type: in.Type,
		}

		_, err = l.svcCtx.CoursewareModel.Insert(l.ctx, &newCourseware)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCodeMsg(500, "新增课件失败"), "code: %s,err:%v", in.Code, err)
		}

		return &courseware.AddResponse{}, nil
	}

	return nil, errors.Wrapf(xerr.NewErrCodeMsg(500, "查询课件失败"), "code: %s,err:%v", in.Code, err)
}
