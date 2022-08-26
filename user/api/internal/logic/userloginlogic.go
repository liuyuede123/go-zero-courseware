package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"go-zero-courseware/user/rpc/userclient"
	"google.golang.org/grpc/status"
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
		return nil, err
	}

	now := time.Now().Unix()
	login.Token, err = l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, int64(login.Id))
	if err != nil {
		return nil, status.Error(5000, err.Error())
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
