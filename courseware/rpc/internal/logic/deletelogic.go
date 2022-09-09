package logic

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"

	"go-zero-courseware/courseware/rpc/courseware"
	"go-zero-courseware/courseware/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteLogic) Delete(in *courseware.DeleteRequest) (*courseware.DeleteResponse, error) {
	err := l.svcCtx.CoursewareModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrapf(status.Errorf(500, "删除课件失败"), "id: %d,err:%v", in.Id, err)
	}

	return &courseware.DeleteResponse{}, nil
}
