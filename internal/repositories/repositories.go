package repositories

import (
	"strings"

	"github.com/flytrap/go-template/internal/config"
	"github.com/flytrap/go-template/internal/models"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var RepositorySet = wire.NewSet(NewLoggingRepository)

func AutoMigrate(db *gorm.DB) error {
	if dbType := config.C.Gorm.DBType; strings.ToLower(dbType) == "mysql" {
		db = db.Set("gorm:table_options", "ENGINE=InnoDB")
	}

	return db.AutoMigrate(
		new(models.Log),
	)
}
