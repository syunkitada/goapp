package base_db_model

import "time"

type Model struct {
	CreatedAt time.Time `gorm:"not null;"`
	UpdatedAt time.Time `gorm:"not null;"`
}
