package model

import (
	"asssistant/protocol/internal/protocol"
	"gorm.io/gorm"
)

type Connection struct {
	gorm.Model
	ConnectionName string            `gorm:"index:idx_name,unique"`
	Protocol       protocol.Protocol `gorm:"index"`
	Address        string            `gorm:"index"`
	Port           int               `gorm:"index"`
}
