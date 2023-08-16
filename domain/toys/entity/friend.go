package entity

import (
	"time"
)

type Friend struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Url        string    `json:"url"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (a *Friend) Create() {
	a.CreateTime = time.Now()
	a.UpdateTime = time.Now()
}

func (a *Friend) Update() {
	a.UpdateTime = time.Now()
}
