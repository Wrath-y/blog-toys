package facade

import "toys/domain/toys/repository/po"

type ToysRepositoryI interface {
	FindFriends() ([]po.FriendLink, error)
}
