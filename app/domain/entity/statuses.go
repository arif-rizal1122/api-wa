package entity


import "time"


type Status struct {
	ID        int       `gorm:"primaryKey;autoIncrement;column:id" db:"id"`
	Picture   string    `gorm:"type:varchar(100);not null;column:picture" db:"picture"`
	Caption   string    `gorm:"type:varchar(100);not null;column:caption" db:"caption"`
	UserId    int       `gorm:"column:user_id" db:"user_id"`
	CreatedAt time.Time `gorm:"column:created_at" db:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" db:"updated_at"`

	// Definisi relasi
	User User `gorm:"foreignKey:UserId"`
}

func (Status) TableName() string {
	return "statuses"
}

