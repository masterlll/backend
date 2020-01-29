package model

import (
	"time"
)

type User struct {
	UserID     string    `db:"user_id" json:"UserId `
	Name       string    `db:"name" json:"name"`
	CreateTime time.Time `db:"create_at" json:"createTime" `
	UpdateTime time.Time `db:"update_at" json:"updateTime" `
}
