// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package handler

import (
	"net/http"

	"GolangProject/cloud-disk/core/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/mail/send",
				Handler: MailCodeSendHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/detail",
				Handler: UserDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: UserHandler(serverCtx),
			},
		},
	)
}
