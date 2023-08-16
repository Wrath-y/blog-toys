package assembler

import (
	"toys/domain/toys/entity"
	"toys/interfaces/dto"
)

func ToFriends(list []entity.Friend) []dto.Friend {
	res := make([]dto.Friend, 0, len(list))
	for _, v := range list {
		res = append(res, dto.Friend{
			Name: v.Name,
			Url:  v.Url,
		})
	}

	return res
}
