package model

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model

	Payload  string `json:"payload"`
	Progress int    `json:"progress"`
	Status   string `json:"status"`
}
