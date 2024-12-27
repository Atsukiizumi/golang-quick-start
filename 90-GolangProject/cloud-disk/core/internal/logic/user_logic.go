package logic

import (
	"GolangProject/cloud-disk/core/helper"
	"GolangProject/cloud-disk/core/models"
	"context"
	"errors"

	"GolangProject/cloud-disk/core/internal/svc"
	"GolangProject/cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// 1. 从数据库中查询当前用户s
	user := new(models.UserBasic)
	get, err := models.Engine.Where("name = ? AND password = ?", req.Name, helper.Md5(req.Password)).Get(user)
	if err != nil {
		return nil, err
	}
	if !get {
		return nil, errors.New("用户名或密码错误")
	}
	// 2. 生成token
	token, err := helper.GeneratorToken(user.Id, user.Identity, user.Name)
	if err != nil {
		return nil, err
	}
	resp = new(types.LoginResponse)
	resp.Token = token

	return
}
