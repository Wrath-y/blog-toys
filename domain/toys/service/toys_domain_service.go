package service

import (
	osssdk "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-redis/redis/v7"
	"toys/domain/toys/entity"
	"toys/domain/toys/repository/facade"
	"toys/domain/toys/repository/persistence"
	"toys/domain/toys/repository/po"
	"toys/infrastructure/common/context"
	baseEvent "toys/infrastructure/common/event"
	"toys/infrastructure/util/oss"
)

type ToysDomainService struct {
	*context.Context
	toysFactory     ToysFactory
	toysCache       ToysCache
	toysRepositoryI facade.ToysRepositoryI
	publisherI      baseEvent.PublisherI
}

func NewToysDomainService(ctx *context.Context) ToysDomainService {
	return ToysDomainService{
		Context:         ctx,
		toysFactory:     NewToysFactory(),
		toysCache:       NewToysCache(),
		toysRepositoryI: persistence.NewFriendLinkRepository(),
		publisherI:      baseEvent.NewBasePublisher(),
	}
}

func (td *ToysDomainService) FindFriends() ([]entity.Friend, error) {
	var err error
	list := make([]po.FriendLink, 0)

	list, err = td.toysCache.GetFriendLinks()
	if err != nil && err != redis.Nil {
		td.Logger.ErrorL("获取友链缓存失败", "", err.Error())
		return nil, err
	}
	if err == nil {
		return td.toysFactory.CreateFriendEntities(list), nil
	}

	list, err = td.toysRepositoryI.FindFriends()
	if err != nil {
		td.Logger.ErrorL("获取友链失败", "", err.Error())
		return nil, err
	}

	if err := td.toysCache.SetFriendLinks(list); err != nil {
		td.Logger.ErrorL("缓存友链失败", "", err.Error())
	}

	return td.toysFactory.CreateFriendEntities(list), nil
}

func (td *ToysDomainService) GetPixivs(nextMarker string, page int32) (*osssdk.ListObjectsResult, error) {
	oss.Setup()

	list, err := oss.List(nextMarker, int(page))
	if err != nil {
		td.Logger.ErrorL("获取pixiv失败", page, err.Error())
		return nil, err
	}

	return list, nil
}
