// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: in.proto

package pushmessageservice

import (
	"context"

	"puhser/getway/in/proto/InnerGetWay"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	PushMessageReq  = InnerGetWay.PushMessageReq
	PushMessageResp = InnerGetWay.PushMessageResp

	PushMessageService interface {
		PushMessage(ctx context.Context, in *PushMessageReq, opts ...grpc.CallOption) (*PushMessageResp, error)
	}

	defaultPushMessageService struct {
		cli zrpc.Client
	}
)

func NewPushMessageService(cli zrpc.Client) PushMessageService {
	return &defaultPushMessageService{
		cli: cli,
	}
}

func (m *defaultPushMessageService) PushMessage(ctx context.Context, in *PushMessageReq, opts ...grpc.CallOption) (*PushMessageResp, error) {
	client := InnerGetWay.NewPushMessageServiceClient(m.cli.Conn())
	return client.PushMessage(ctx, in, opts...)
}
