// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: auth.proto

package server

import (
	"context"

	"auth/internal/logic"
	"auth/internal/svc"
	"auth/proto/AuthRpc"
)

type AuthServiceServer struct {
	svcCtx *svc.ServiceContext
	AuthRpc.UnimplementedAuthServiceServer
}

func NewAuthServiceServer(svcCtx *svc.ServiceContext) *AuthServiceServer {
	return &AuthServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *AuthServiceServer) Authentication(ctx context.Context, in *AuthRpc.AuthenticationReq) (*AuthRpc.AuthenticationResp, error) {
	l := logic.NewAuthenticationLogic(ctx, s.svcCtx)
	return l.Authentication(in)
}

func (s *AuthServiceServer) RefreshSession(ctx context.Context, in *AuthRpc.RefreshSessionReq) (*AuthRpc.RefreshSessionResp, error) {
	l := logic.NewRefreshSessionLogic(ctx, s.svcCtx)
	return l.RefreshSession(in)
}

func (s *AuthServiceServer) DeleteSession(ctx context.Context, in *AuthRpc.DeleteSessionReq) (*AuthRpc.DeleteSessionResp, error) {
	l := logic.NewDeleteSessionLogic(ctx, s.svcCtx)
	return l.DeleteSession(in)
}

func (s *AuthServiceServer) CreateVoucher(ctx context.Context, in *AuthRpc.CreateVoucherReq) (*AuthRpc.CreateVoucherResp, error) {
	l := logic.NewCreateVoucherLogic(ctx, s.svcCtx)
	return l.CreateVoucher(in)
}
