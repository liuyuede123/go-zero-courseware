package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"go-zero-courseware/courseware/rpc/courseware"
	"go-zero-courseware/courseware/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLogic) Get(in *courseware.GetRequest) (*courseware.GetResponse, error) {
	cw, err := l.svcCtx.CoursewareModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, status.Error(5000, "课件不存在")
	}

	return &courseware.GetResponse{
		Id:   cw.Id,
		Code: cw.Code,
		Name: cw.Name,
		Type: cw.Type,
	}, nil
}
