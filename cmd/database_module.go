package cmd

import (
	"github.com/samber/lo"
	"go.uber.org/fx"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"path/filepath"
	"protocol/assistant/internal/model"
)

var databaseModule = fx.Module("databaseModule", fx.Provide(newDatabaseModule), fx.Invoke(autoMigrate))

func newDatabaseModule() *gorm.DB {
	cache, err := os.UserCacheDir()
	if err != nil {
		panic(err)
	}
	println(cache)
	path := filepath.Join(cache, "protocol.db")
	println(path)
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	return db
}

func autoMigrate(db *gorm.DB) {
	lo.Must0(db.AutoMigrate(model.Connection{}))
}
