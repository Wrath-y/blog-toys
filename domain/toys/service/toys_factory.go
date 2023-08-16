package service

import (
	"toys/domain/toys/entity"
	"toys/domain/toys/repository/po"
)

type ToysFactory struct {
}

func NewToysFactory() ToysFactory {
	return ToysFactory{}
}

func (a *ToysFactory) CreateFriendEntities(poList []po.FriendLink) []entity.Friend {
	res := make([]entity.Friend, 0, len(poList))
	for _, v := range poList {
		tmp := a.CreateArticleEntity(v)
		res = append(res, tmp)
	}

	return res
}

func (*ToysFactory) CreateArticleEntity(po po.FriendLink) entity.Friend {
	return entity.Friend{
		Id:         po.Id,
		Name:       po.Name,
		Email:      po.Email,
		Url:        po.Url,
		CreateTime: po.CreateTime,
		UpdateTime: po.UpdateTime,
	}
}
