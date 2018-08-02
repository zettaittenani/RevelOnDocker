package models

import (
    "github.com/jinzhu/gorm"
)

type Test struct {
    gorm.Model
    Name string
    Age  uint
}
