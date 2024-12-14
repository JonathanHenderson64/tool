package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var g_db *gorm.DB

func GetDb() *gorm.DB {
	return g_db
}
func ConnectMysql(dns string) (err error) {
	g_db, err = gorm.Open(mysql.Open(dns), &gorm.Config{Logger: logger.Default.LogMode(logger.Error)})
	if nil != err {
		return err
	}

	sqlDB, err := g_db.DB()
	if nil != err {
		return err
	}

	sqlDB.SetConnMaxIdleTime(time.Hour)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(-1)
	return nil
}

func CreateTables(tables ...any) error {
	for _, table := range tables {
		if !g_db.Migrator().HasTable(table) {
			return g_db.AutoMigrate(table)
		}
	}
	return nil
}

func CloseMysql() {
	if nil == g_db {
		return
	}
	sqlDB, err := g_db.DB()
	if nil != err || sqlDB == nil {
		return
	}
	sqlDB.Close()
}
