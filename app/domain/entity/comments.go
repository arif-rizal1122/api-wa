package entity

import "time"

type Comments struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement;column:id" db:"id"`
	IsiComment string    `gorm:"type:text;column:isi_comment" db:"isi_comment"`
	Picture    string    `gorm:"type:varchar(100);not null;column:picture" db:"picture"`
	UserId     uint64    `gorm:"column:user_id" db:"user_id"`
	StatusId   uint64    `gorm:"column:status_id" db:"status_id"`
	CreatedAt  time.Time `gorm:"column:created_at" db:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" db:"updated_at"`
}

func (Comments) TableName() string {
	return "comments"
}
