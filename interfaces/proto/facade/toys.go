package facade

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"toys/application/service"
	grpcCtx "toys/infrastructure/common/context"
	"toys/infrastructure/common/errcode"
	"toys/interfaces/assembler"
	"toys/interfaces/proto"
	"toys/launch/grpc/resp"
)

type Toys struct{}

func (*Toys) GetPixivs(ctx context.Context, req *proto.GetPixivsReq) (*proto.Response, error) {
	list, err := service.NewArticleApplicationService(grpcCtx.NewContext(ctx)).GetPixivs(req.NextMarker, req.Page)
	if err != nil {
		return resp.FailWithErrCode(errcode.GetPixivsFailed)
	}
	return resp.Success(list)
}

func (*Toys) GetFriends(ctx context.Context, req *emptypb.Empty) (*proto.Response, error) {
	list, err := service.NewArticleApplicationService(grpcCtx.NewContext(ctx)).GetFindFriendsById()
	if err != nil {
		return resp.FailWithErrCode(errcode.FindFriendsFailed)
	}

	return resp.Success(assembler.ToFriends(list))
}
