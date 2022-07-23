package dao

import (
	"time"

	"github.com/serialt/cli/config"
	"github.com/serialt/cli/sugar"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

// func NewDatabase() *gorm.DB {
// 	return NewDBConnect(&t.Config.Database)
// }

// 设置gorm日志使用zap
var GormLogger zapgorm2.Logger
var logSugar = sugar.Log

func init() {
	GormLogger = zapgorm2.New(sugar.Logger)
	GormLogger.SetAsDefault()
}

// 创建mydb
func NewDBConnect(db *config.Database) *gorm.DB {

	var GormDB *gorm.DB
	switch db.Type {
	// case "sqlite":
	// 	mySqlite := NewMySqlite()
	// 	GormDB = SqliteInit(mySqlite)
	// 	fmt.Println("使用的数据库是 sqlite")
	case "mysql":
		GormDB = GetMysqlGormDB(db)
		sugar.Log.Info("使用的数据库是 mysql")
	case "postgresql":
		GormDB = GetPostgreSQLGormDB(db)
		sugar.Log.Info("使用的数据库是 postgresql")
	default:
		sugar.Log.Info("The database is not supported, please choice [sqlite] or [mysql]")
	}
	if db.MaxOpenConns != 0 {

		// Gorm 使用database/sql 维护连接池
		sqlDB, _ := GormDB.DB()

		// 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxIdleConns(db.MaxIdleConns)

		// 设置打开数据库连接的最大数量
		sqlDB.SetMaxOpenConns(db.MaxOpenConns)

		// 设置了连接可复用的最大时间
		sqlDB.SetConnMaxLifetime(db.ConnMaxLifetime)
	} else {
		sqlDB, _ := GormDB.DB()

		// SetMaxIdleConns 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxIdleConns(10)

		// SetMaxOpenConns 设置打开数据库连接的最大数量。
		sqlDB.SetMaxOpenConns(100)

		// SetConnMaxLifetime 设置了连接可复用的最大时间。
		sqlDB.SetConnMaxLifetime(time.Hour)

	}

	return GormDB

}