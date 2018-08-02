package models

import (
    "time"
)

type User struct {
    ID        uint64     `gorm:"primary_key" json:"id"`
    Name      string     `sql:"size:64" json:"name" validate:"max=64"`
    CreatedAt time.Time  `json:"created_at"`
    UpdatedAt time.Time  `json:"updated_at"`
    DeletedAt *time.Time `json:"deleted_at"`
}
