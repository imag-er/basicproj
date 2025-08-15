package dal

import (
	"backend/config"
	"backend/dal/models"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initDB() {

	dbconfig := config.Config.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbconfig.Username, dbconfig.Password, dbconfig.Host, dbconfig.Port, dbconfig.Name,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		hlog.Fatal("unable to connect to database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		hlog.Fatal("unable to connect to database")
	}

	sqlDB.SetMaxIdleConns(dbconfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(dbconfig.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(dbconfig.ConnMaxLifetime))

	// Register all models for auto-migration
	modelsToMigrate := []interface{}{
		&models.User{},
		&models.Article{},
	}
	if err := db.AutoMigrate(modelsToMigrate...); err != nil {
		hlog.Fatal("failed to migrate database")
	}

	DB = db
}
