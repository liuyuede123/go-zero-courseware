package logic

import (
	"context"
	"github.com/pkg/errors"
	"go-zero-courseware/courseware/common/xerr"
	"go-zero-courseware/courseware/rpc/model"
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
		if err == model.ErrNotFound {
			return nil, xerr.NewErrCodeMsg(5000, "数据不存在")
		}
		return nil, errors.Wrapf(status.Errorf(500, "查询课件失败"), "id: %d,err:%v", in.Id, err)
	}
	return &courseware.GetResponse{
		Id:   cw.Id,
		Code: cw.Code,
		Name: cw.Name,
		Type: cw.Type,
	}, nil

}
