package logic

import (
	"context"
	"go-zero-courseware/courseware/common/xerr"
	"go-zero-courseware/courseware/rpc/coursewareclient"

	"go-zero-courseware/courseware/api/internal/svc"
	"go-zero-courseware/courseware/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CoursewareGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCoursewareGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CoursewareGetLogic {
	return &CoursewareGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CoursewareGetLogic) CoursewareGet(req *types.GetRequest) (resp *types.GetResponse, err error) {
	cw, err := l.svcCtx.CoursewareRpc.Get(l.ctx, &coursewareclient.GetRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, xerr.NewErrCodeMsg(5000, "获取课件失败")
	}

	return &types.GetResponse{
		Id:   cw.Id,
		Code: cw.Code,
		Name: cw.Name,
		Type: cw.Type,
	}, nil
}
