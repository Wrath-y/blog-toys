package persistence

import (
	"toys/domain/toys/repository/facade"
	"toys/domain/toys/repository/po"
	"toys/infrastructure/util/db"
)

type FriendLinkRepository struct{}

func NewFriendLinkRepository() facade.ToysRepositoryI {
	return &FriendLinkRepository{}
}

func (*FriendLinkRepository) FindFriends() ([]po.FriendLink, error) {
	var res []po.FriendLink
	return res, db.Orm.Raw("select * from friend_link order by id desc").Find(&res).Error
}
