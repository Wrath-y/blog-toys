package po

import (
	"time"
)

type FriendLink struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Url        string    `json:"url"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (*FriendLink) TableName() string {
	return "friend_link"
}
