package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"go-zero-courseware/courseware/rpc/courseware"
	"go-zero-courseware/courseware/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *courseware.UpdateRequest) (*courseware.UpdateResponse, error) {
	cw, err := l.svcCtx.CoursewareModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, status.Error(5000, "课件不存在")
	}

	codeCw, err := l.svcCtx.CoursewareModel.FindOneByCode(l.ctx, in.Code)
	if err == nil && codeCw.Id != cw.Id {
		return nil, status.Error(5000, "课件编号已存在")
	}

	if in.Code != "" {
		cw.Code = in.Code
	}
	if in.Name != "" {
		cw.Name = in.Name
	}
	if in.Type > 0 {
		cw.Type = in.Type
	}
	err = l.svcCtx.CoursewareModel.Update(l.ctx, cw)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &courseware.UpdateResponse{}, nil
}
