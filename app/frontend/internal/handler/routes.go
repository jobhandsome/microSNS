// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/jobhandsome/microSNS/app/frontend/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: LoginHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/cate/lists",
				Handler: CateListHandler(serverCtx),
			},
		},
	)
}