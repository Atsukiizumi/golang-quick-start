package handler

import (
	"net/http"

	"GolangProject/cloud-disk/core/internal/logic"
	"GolangProject/cloud-disk/core/internal/svc"
	"GolangProject/cloud-disk/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func MailCodeSendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MailCodeSendRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewMailCodeSendLogic(r.Context(), svcCtx)
		resp, err := l.MailCodeSend(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
