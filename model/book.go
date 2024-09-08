package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	Id          string `gorm:"type:uuid;primaryKey" json:"id"`
	Title       string `gorm:"type:varchar" json:"title"`
	Author      string `gorm:"type:varchar" json:"author"`
	Description string `gorm:"type:varchar" json:"description"`
	CreatedAt   int64  `gorm:"type:bigint" json:"created_at"`
	CreatedBy   string `gorm:"type:varchar" json:"created_by"`
	UpdatedAt   int64  `gorm:"type:bigint" json:"updated_at"`
	UpdatedBy   string `gorm:"type:varchar" json:"updated_by"`
}

func (base *Book) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New()
	time := time.Now().UnixMilli()
	tx.Statement.SetColumn("id", uuid)
	tx.Statement.SetColumn("CreatedAt", time)
	tx.Statement.SetColumn("UpdatedAt", time)
	return nil
}

func (base *Book) BeforeUpdate(tx *gorm.DB) error {
	time := time.Now().UnixMilli()
	tx.Statement.SetColumn("UpdatedAt", time)
	return nil
}