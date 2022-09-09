package logic

import (
	"context"
	"go-zero-courseware/courseware/common/xerr"
	"go-zero-courseware/courseware/rpc/coursewareclient"

	"go-zero-courseware/courseware/api/internal/svc"
	"go-zero-courseware/courseware/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CoursewareDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCoursewareDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CoursewareDeleteLogic {
	return &CoursewareDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CoursewareDeleteLogic) CoursewareDelete(req *types.DeleteRequest) (resp *types.DeleteResponse, err error) {
	_, err = l.svcCtx.CoursewareRpc.Delete(l.ctx, &coursewareclient.DeleteRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, xerr.NewErrCodeMsg(5000, "删除课件失败")
	}

	return &types.DeleteResponse{}, nil
}
