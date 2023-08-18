package entity

import (
	"time"

	"gorm.io/gorm"
)

type BasicModel struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type BasicAccessModel struct {
	UUID   string `json:"uuid"`
	UserID uint   `json:"userId"`
}

type PageInfo struct {
	Page     int `json:"page" binding:"required"`
	PageSize int `json:"pageSize" binding:"required"`
}

type PageResult struct {
	List  interface{} `json:"list"`
	Total int64       `json:"total"`
}
