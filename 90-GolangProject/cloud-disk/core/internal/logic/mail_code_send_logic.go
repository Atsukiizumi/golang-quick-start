package logic

import (
	"GolangProject/cloud-disk/core/helper"
	"context"

	"GolangProject/cloud-disk/core/internal/svc"
	"GolangProject/cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MailCodeSendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailCodeSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MailCodeSendLogic {
	return &MailCodeSendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MailCodeSendLogic) MailCodeSend(req *types.MailCodeSendRequest) (resp *types.MailCodeSendResponse, err error) {
	err = helper.SendCodeMail(req.Email, "123456")
	if err != nil {
		return nil, err
	}
	return resp, nil
}
