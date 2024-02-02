package cmd

import (
	"asssistant/protocol/internal/model"
	"github.com/samber/lo"
	"go.uber.org/fx"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"path/filepath"
)

var databaseModule = fx.Module("databaseModule", fx.Provide(newDatabaseModule), fx.Invoke(autoMigrate, queryHistory))

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

var tables = []interface{}{model.Connection{}}

func autoMigrate(db *gorm.DB) {
	go func() {
		migrator := db.Migrator()
		var notExistsTable []interface{}
		for _, table := range tables {
			if migrator.HasTable(table) {
				continue
			}
			notExistsTable = append(notExistsTable, table)
		}
		lo.Must0(db.AutoMigrate(notExistsTable...))
	}()
}

func queryHistory(db *gorm.DB) {
	var connections []model.Connection
	db.Find(&connections)
}
