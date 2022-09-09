package logic

import (
	"context"
	"github.com/pkg/errors"
	"go-zero-courseware/user/common/xerr"
	"go-zero-courseware/user/rpc/internal/svc"
	"go-zero-courseware/user/rpc/model"
	"go-zero-courseware/user/rpc/user"
	"golang.org/x/crypto/bcrypt"

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
		return nil, errors.Wrapf(xerr.NewErrCodeMsg(500, "数据不存在"), "loginName: %s,err:%v", in.LoginName, err)
	}
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCodeMsg(500, "查询用户失败"), "loginName: %s,err:%v", in.LoginName, err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(in.Password))
	if err != nil {
		return nil, xerr.NewErrCodeMsg(5000, "密码错误")
	}

	return &user.LoginResponse{
		Id:    userInfo.Id,
		Token: "a.b.c",
	}, nil
}
