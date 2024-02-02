package model

import (
	"gorm.io/gorm"
	"protocol/assistant/internal/protocol"
)

type Connection struct {
	gorm.Model
	Protocol protocol.Protocol
}
