// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: relation.proto

package relationservice

import (
	"context"

	"fansX/services/relation/proto/relationRpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CancelFollowReq      = relationRpc.CancelFollowReq
	Empty                = relationRpc.Empty
	FollowReq            = relationRpc.FollowReq
	GetFollowerNumsReq   = relationRpc.GetFollowerNumsReq
	GetFollowerNumsResp  = relationRpc.GetFollowerNumsResp
	GetFollowingNumsReq  = relationRpc.GetFollowingNumsReq
	GetFollowingNumsResp = relationRpc.GetFollowingNumsResp
	IsFollowingReq       = relationRpc.IsFollowingReq
	IsFollowingResp      = relationRpc.IsFollowingResp
	ListFollowerReq      = relationRpc.ListFollowerReq
	ListFollowerResp     = relationRpc.ListFollowerResp
	ListFollowingReq     = relationRpc.ListFollowingReq
	ListFollowingResp    = relationRpc.ListFollowingResp

	RelationService interface {
		Follow(ctx context.Context, in *FollowReq, opts ...grpc.CallOption) (*Empty, error)
		CancelFollow(ctx context.Context, in *CancelFollowReq, opts ...grpc.CallOption) (*Empty, error)
		ListFollowing(ctx context.Context, in *ListFollowingReq, opts ...grpc.CallOption) (*ListFollowingResp, error)
		IsFollowing(ctx context.Context, in *IsFollowingReq, opts ...grpc.CallOption) (*IsFollowingResp, error)
		ListFollower(ctx context.Context, in *ListFollowerReq, opts ...grpc.CallOption) (*ListFollowerResp, error)
		GetFollowingNums(ctx context.Context, in *GetFollowingNumsReq, opts ...grpc.CallOption) (*GetFollowingNumsResp, error)
		GetFollowerNums(ctx context.Context, in *GetFollowerNumsReq, opts ...grpc.CallOption) (*GetFollowerNumsResp, error)
	}

	defaultRelationService struct {
		cli zrpc.Client
	}
)

func NewRelationService(cli zrpc.Client) RelationService {
	return &defaultRelationService{
		cli: cli,
	}
}

func (m *defaultRelationService) Follow(ctx context.Context, in *FollowReq, opts ...grpc.CallOption) (*Empty, error) {
	client := relationRpc.NewRelationServiceClient(m.cli.Conn())
	return client.Follow(ctx, in, opts...)
}

func (m *defaultRelationService) CancelFollow(ctx context.Context, in *CancelFollowReq, opts ...grpc.CallOption) (*Empty, error) {
	client := relationRpc.NewRelationServiceClient(m.cli.Conn())
	return client.CancelFollow(ctx, in, opts...)
}

func (m *defaultRelationService) ListFollowing(ctx context.Context, in *ListFollowingReq, opts ...grpc.CallOption) (*ListFollowingResp, error) {
	client := relationRpc.NewRelationServiceClient(m.cli.Conn())
	return client.ListFollowing(ctx, in, opts...)
}

func (m *defaultRelationService) IsFollowing(ctx context.Context, in *IsFollowingReq, opts ...grpc.CallOption) (*IsFollowingResp, error) {
	client := relationRpc.NewRelationServiceClient(m.cli.Conn())
	return client.IsFollowing(ctx, in, opts...)
}

func (m *defaultRelationService) ListFollower(ctx context.Context, in *ListFollowerReq, opts ...grpc.CallOption) (*ListFollowerResp, error) {
	client := relationRpc.NewRelationServiceClient(m.cli.Conn())
	return client.ListFollower(ctx, in, opts...)
}

func (m *defaultRelationService) GetFollowingNums(ctx context.Context, in *GetFollowingNumsReq, opts ...grpc.CallOption) (*GetFollowingNumsResp, error) {
	client := relationRpc.NewRelationServiceClient(m.cli.Conn())
	return client.GetFollowingNums(ctx, in, opts...)
}

func (m *defaultRelationService) GetFollowerNums(ctx context.Context, in *GetFollowerNumsReq, opts ...grpc.CallOption) (*GetFollowerNumsResp, error) {
	client := relationRpc.NewRelationServiceClient(m.cli.Conn())
	return client.GetFollowerNums(ctx, in, opts...)
}
