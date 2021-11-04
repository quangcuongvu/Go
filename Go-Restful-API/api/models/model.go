package models

import "time"

type Model struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	CreatAt   time.Time `gorm:"defaul:current_timestamp;notnull" json:"creatAt"`
	UpdatedAt time.Time `gorm:"defaul:current_timestamp;notnull" json:"updatedAt"`
}
