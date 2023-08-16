package service

import (
	"encoding/json"
	"time"
	"toys/domain/toys/repository/po"
	"toys/infrastructure/util/goredis"
	"toys/infrastructure/util/util/highperf"
)

const (
	ListStrKey = "blog:friend-link:list"
)

type ToysCache struct {
}

func NewToysCache() ToysCache {
	return ToysCache{}
}

func (*ToysCache) GetFriendLinks() ([]po.FriendLink, error) {
	list := make([]po.FriendLink, 0)
	b, err := goredis.Client.Get(ListStrKey).Bytes()
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &list); err != nil {
		return nil, err
	}

	return list, nil
}

func (*ToysCache) SetFriendLinks(list []po.FriendLink) error {
	b, err := json.Marshal(list)
	if err != nil {
		return err
	}
	return goredis.Client.Set(ListStrKey, highperf.Bytes2str(b), time.Hour*24*7).Err()
}
