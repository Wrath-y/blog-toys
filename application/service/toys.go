package service

import (
	osssdk "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"toys/domain/toys/entity"
	"toys/domain/toys/service"
	"toys/infrastructure/common/context"
)

type ToysApplicationService struct {
	*context.Context
	articleDomainService service.ToysDomainService
}

func NewArticleApplicationService(ctx *context.Context) *ToysApplicationService {
	return &ToysApplicationService{
		ctx,
		service.NewToysDomainService(ctx),
	}
}

func (a *ToysApplicationService) GetPixivs(nextMarker string, page int32) (*osssdk.ListObjectsResult, error) {
	return a.articleDomainService.GetPixivs(nextMarker, page)
}

func (a *ToysApplicationService) GetFindFriendsById() ([]entity.Friend, error) {
	return a.articleDomainService.FindFriends()
}
