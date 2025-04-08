// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: content.proto

package server

import (
	"context"

	"bilibili/services/content/meta/internal/logic"
	"bilibili/services/content/meta/internal/svc"
	"bilibili/services/content/meta/proto/metaContentRpc"
)

type MetaContentServiceServer struct {
	svcCtx *svc.ServiceContext
	metaContentRpc.UnimplementedMetaContentServiceServer
}

func NewMetaContentServiceServer(svcCtx *svc.ServiceContext) *MetaContentServiceServer {
	return &MetaContentServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *MetaContentServiceServer) Publish(ctx context.Context, in *metaContentRpc.PublishReq) (*metaContentRpc.Empty, error) {
	l := logic.NewPublishLogic(ctx, s.svcCtx)
	return l.Publish(in)
}

func (s *MetaContentServiceServer) Update(ctx context.Context, in *metaContentRpc.UpdateReq) (*metaContentRpc.Empty, error) {
	l := logic.NewUpdateLogic(ctx, s.svcCtx)
	return l.Update(in)
}

func (s *MetaContentServiceServer) Delete(ctx context.Context, in *metaContentRpc.DeleteReq) (*metaContentRpc.Empty, error) {
	l := logic.NewDeleteLogic(ctx, s.svcCtx)
	return l.Delete(in)
}

func (s *MetaContentServiceServer) StatusSearch(ctx context.Context, in *metaContentRpc.StatusSearchReq) (*metaContentRpc.StatusSearchResp, error) {
	l := logic.NewStatusSearchLogic(ctx, s.svcCtx)
	return l.StatusSearch(in)
}
