package logic

import (
	"GolangProject/cloud-disk/core/models"
	"context"
	"errors"

	"GolangProject/cloud-disk/core/internal/svc"
	"GolangProject/cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *types.UserDetailRequest) (resp *types.UserDetailResponse, err error) {
	user := new(models.UserBasic)
	// 查询用户信息
	get, err := models.Engine.Where("identity = ?", req.Identity).Get(user)
	if err != nil {
		return nil, err
	}
	if !get {
		return nil, errors.New("用户不存在")
	}

	resp = new(types.UserDetailResponse)
	resp.Name = user.Name
	resp.Email = user.Email

	return resp, nil
}
