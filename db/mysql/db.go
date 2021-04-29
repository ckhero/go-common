package mysql

import (

	"errors"
	"fmt"
	"github.com/ckhero/go-common/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type DBMap map[string]*gorm.DB

type TransFunc func(tx *gorm.DB) error

var dbMap = make(DBMap)

/**
 * 获取数据库连接
 */
func getDB(name ...string) *gorm.DB {
	key := "default"
	if len(name) > 0 {
		key = name[0]
	}
	return dbMap[key]
}

/**
 * 连接数据库
 */
func ConnectDB(dbConfigMap map[string]config.Database) {
	for name, dbConfig := range dbConfigMap {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&loc=Local", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database, dbConfig.Charset)
		db, err := gorm.Open(mysql.Dialector{
			Config: &mysql.Config{
				DriverName:                dbConfig.Dialect,
				DSN:                       dsn,
				SkipInitializeWithVersion: false,
				DefaultStringSize:         256,
				DisableDatetimePrecision:  true,
				DontSupportRenameIndex:    true,
				DontSupportRenameColumn:   true,
			},
		}, &gorm.Config{
			Logger: logger.Default.LogMode(dbConfig.LogMode),
		})

		if err != nil {
			panic(fmt.Errorf("fatal error: connect database: %s\n", err))
		}
		sqlDB, err := db.DB()
		if err != nil {
			panic(fmt.Errorf("fatal error: connect database: %s\n", err))
		}
		//db.Logger = logrus.StandardLogger()
		//db.DB().SetLogger(logrus.StandardLogger())
		fmt.Println(dbConfig)
		sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConnNum)
		sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConnNum)
		sqlDB.SetConnMaxLifetime(time.Hour)

		db.AllowGlobalUpdate = false
		//db.InstantSet("gorm:save_associations", false)
		//db.InstantSet("gorm:association_save_reference", false)

		AddGormCallbacks(db)

		dbMap[name] = db
	}
}

/**
 * 关闭数据库连接
 */
func CloseDB() {
	if len(dbMap) > 0 {
		for _, db := range dbMap {
			sqlDB, err := db.DB()
			if err != nil {
				panic(fmt.Errorf("fatal error: connect database: %s\n", err))
			}
			_ = sqlDB.Close()
		}
	}
}

/**
 * 执行事物操作
 */
func Transaction(db *gorm.DB, closures ...TransFunc) (err error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			switch r.(type) {
			case error:
				err = r.(error)
			case string:
				err = errors.New(r.(string))
			default:
				err = errors.New("system internal error")
			}
		}
	}()

	if tx.Error != nil {
		return tx.Error
	}

	for _, closure := range closures {
		if err := closure(tx); err != nil {
			tx.Rollback()
			return err
		}
		if tx.Error != nil {
			tx.Rollback()
			return tx.Error
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
