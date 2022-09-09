package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"go-zero-courseware/user/common/xerr"
	"go-zero-courseware/user/rpc/userclient"
	"time"

	"go-zero-courseware/user/api/internal/svc"
	"go-zero-courseware/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	login, err := l.svcCtx.UserRpc.Login(l.ctx, &userclient.LoginRequest{
		LoginName: req.LoginName,
		Password:  req.Password,
	})
	if err != nil {
		return nil, xerr.NewErrCodeMsg(500, "用户登录失败")
	}

	now := time.Now().Unix()
	login.Token, err = l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, int64(login.Id))
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCodeMsg(5000, "token生成失败"), "loginName: %s,err:%v", req, err)
	}
	return &types.LoginResponse{
		Id:    login.Id,
		Token: login.Token,
	}, nil
}

func (l *UserLoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
