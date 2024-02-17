package repository

import (
	"asssistant/protocol/internal/model"
	"gorm.io/gorm"
)

func NewConnectionRepository(db *gorm.DB) *ConnectionRepository {
	repository := NewBaseRepository[model.Test](db)
	//return &ConnectionRepository{BaseRepository: repository}
	return nil
}

type ConnectionRepository struct {
}

func (t ConnectionRepository) FindByConnectionName(name string) {

}
