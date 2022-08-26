package logic

import (
	"context"
	"go-zero-courseware/user/rpc/model"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/status"

	"go-zero-courseware/user/rpc/internal/svc"
	"go-zero-courseware/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	userInfo, err := l.svcCtx.UserModel.FindOneByLoginName(l.ctx, in.LoginName)
	if err == model.ErrNotFound {
		return nil, status.Error(5000, "用户不存在")
	}
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	err = bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(in.Password))
	if err != nil {
		return nil, status.Error(5000, "密码错误")
	}

	return &user.LoginResponse{
		Id:    userInfo.Id,
		Token: "a.b.c",
	}, nil
}
