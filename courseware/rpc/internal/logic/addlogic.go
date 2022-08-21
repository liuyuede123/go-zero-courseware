package logic

import (
	"context"
	"go-zero-courseware/courseware/rpc/courseware"
	"go-zero-courseware/courseware/rpc/internal/svc"
	"go-zero-courseware/courseware/rpc/model"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
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
		return nil, status.Error(5000, "课件已存在")
	}

	if err == model.ErrNotFound {
		newCourseware := model.Courseware{
			Code: in.Code,
			Name: in.Name,
			Type: in.Type,
		}

		_, err = l.svcCtx.CoursewareModel.Insert(l.ctx, &newCourseware)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		return &courseware.AddResponse{}, nil
	}

	return nil, status.Error(500, err.Error())
}
