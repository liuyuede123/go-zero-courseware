package logic

import (
	"context"
	"go-zero-courseware/courseware/common/xerr"
	"go-zero-courseware/courseware/rpc/coursewareclient"

	"go-zero-courseware/courseware/api/internal/svc"
	"go-zero-courseware/courseware/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CoursewareUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCoursewareUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CoursewareUpdateLogic {
	return &CoursewareUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CoursewareUpdateLogic) CoursewareUpdate(req *types.UpdateRequest) (resp *types.UpdateResponse, err error) {
	_, err = l.svcCtx.CoursewareRpc.Update(l.ctx, &coursewareclient.UpdateRequest{
		Id:   req.Id,
		Code: req.Code,
		Name: req.Name,
		Type: req.Type,
	})
	if err != nil {
		return nil, xerr.NewErrCodeMsg(5000, "更新课件失败")
	}

	return &types.UpdateResponse{}, nil
}
