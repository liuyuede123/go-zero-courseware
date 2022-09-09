package logic

import (
	"context"
	"go-zero-courseware/courseware/common/xerr"
	"go-zero-courseware/courseware/rpc/coursewareclient"

	"go-zero-courseware/courseware/api/internal/svc"
	"go-zero-courseware/courseware/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CoursewareAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCoursewareAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CoursewareAddLogic {
	return &CoursewareAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CoursewareAddLogic) CoursewareAdd(req *types.AddRequest) (resp *types.AddResponse, err error) {
	_, err = l.svcCtx.CoursewareRpc.Add(l.ctx, &coursewareclient.AddRequest{
		Code: req.Code,
		Name: req.Name,
		Type: req.Type,
	})
	if err != nil {
		return nil, xerr.NewErrCodeMsg(5000, "新增课件失败")
	}

	return &types.AddResponse{}, nil
}
