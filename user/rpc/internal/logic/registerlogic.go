package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-courseware/user/common/xerr"
	"go-zero-courseware/user/rpc/internal/svc"
	"go-zero-courseware/user/rpc/model"
	"go-zero-courseware/user/rpc/user"
	"golang.org/x/crypto/bcrypt"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	_, err := l.svcCtx.UserModel.FindOneByLoginName(l.ctx, in.LoginName)
	if err == nil {
		return nil, errors.Wrapf(xerr.NewErrCodeMsg(5000, "用户名已存在"), "loginName: %s,err:%v", in.LoginName, err)
	}

	if err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCodeMsg(500, "用户注册失败"), "loginName: %s,err:%v", in.LoginName, err)
	}

	pwd, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCodeMsg(500, "服务器异常"), "loginName: %s,err:%v", in.LoginName, err)
	}
	newUser := model.User{
		LoginName: in.LoginName,
		Username:  in.Username,
		Sex:       in.Sex,
		Password:  string(pwd),
	}
	_, err = l.svcCtx.UserModel.Insert(l.ctx, &newUser)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCodeMsg(500, "写入数据失败"), "loginName: %s,err:%v", in.LoginName, err)
	}

	return &user.RegisterResponse{}, nil
}
