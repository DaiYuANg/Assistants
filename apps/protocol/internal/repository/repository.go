package repository

import "gorm.io/gorm"

type GormModeler interface {
	gorm.Model
}

type Repository[T GormModeler] interface {
	FindById(id uint) T
}

type BaseRepository[T gorm.Model] struct {
	Database *gorm.DB
}

func (b *BaseRepository[T]) FindById(id uint) *T {
	//var t = T{
	//	ID: id,
	//}
	//b.Database.Find(t)
	//return t
	return nil
}

func NewBaseRepository[T GormModeler](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{Database: db}
}
