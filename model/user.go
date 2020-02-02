package model

import (
	"time"
)

type User struct {
	ID         string    `db:"id" json:"UserId `
	Name       string    `db:"nickname" json:"name"`
	Password   string    `db:"password"`
	CreateTime time.Time `db:"create_at" json:"createTime" `
	UpdateTime time.Time `db:"update_at" json:"updateTime" `
}
