package logic

import (
	"context"
	"go-zero-courseware/user/rpc/userclient"
	"google.golang.org/grpc/status"

	"go-zero-courseware/user/api/internal/svc"
	"go-zero-courseware/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	info, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &userclient.UserInfoRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &types.UserInfoResponse{
		Id:        info.Id,
		Username:  info.Username,
		LoginName: info.LoginName,
		Sex:       info.Sex,
	}, nil
}
