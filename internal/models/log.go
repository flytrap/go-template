package models

import "gorm.io/gorm"

type Log struct {
	gorm.Model

	UserId  int64  `json:"user_id" gorm:"index;comment:用户id"`
	Service string `json:"service" gorm:"size:32;comment:服务代号"`
	Type    string `json:"type" gorm:"size:32;comment:日志类型"`
	Content string `json:"content" gorm:"type:text;comment:内容"`
}
