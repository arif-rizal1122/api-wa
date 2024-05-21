package entity

import "time"

type User struct {
	ID        int       `gorm:"primaryKey;autoIncrement;column:id" db:"id"`
	Name      string    `gorm:"type:varchar(100);not null;column:name" db:"name"`
	Username  string    `gorm:"type:varchar(100);not null;unique;column:username" db:"username"`
	Email     string    `gorm:"type:varchar(200);not null;unique;column:email" db:"email"`
	Password  string    `gorm:"type:varchar(100);not null;column:password" db:"password"`
	Phone     string    `gorm:"type:varchar(20);not null;unique;column:phone" db:"phone"`
	CreatedAt time.Time `gorm:"column:created_at" db:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" db:"updated_at"`
}

func (User) TableName() string {
	return "users"
}


