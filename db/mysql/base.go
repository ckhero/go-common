package mysql

import (
	"context"
	"errors"
	sysConst "github.com/ckhero/go-common/constant/sys"
	"gorm.io/gorm"
	"time"
)

type Base struct {
	db *gorm.DB
}

type Time struct {
	BasicTimeFields
	DeletedAt *time.Time `gorm:"column:deletedAt;null"`
}

type BasicTimeFields struct {
	CreatedAt time.Time `gorm:"column:createdAt;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"column:updatedAt;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

type BaseEntity struct {
	CreatedAt	string	`gorm:"<-:false;column:created_at;type:timestamp;not null"`
	UpdatedAt	string	`gorm:"<-:false;column:updated_at;type:timestamp;not null"`
}

type BaseEntityOld struct {
	CreatedAt string `gorm:"<-:false;column:createdAt;type:timestamp;not null"`
	UpdatedAt string `gorm:"<-:false;column:updatedAt;type:timestamp;not null"`
	DeletedAt string `gorm:"column:deletedAt;type:timestamp;not null"`
}

type BaseEntityWithoutDeleteAt struct {
	CreatedAt string `gorm:"<-:false;column:createdAt;type:timestamp;not null"`
	UpdatedAt string `gorm:"<-:false;column:updatedAt;type:timestamp;not null"`
}

type BaseDao struct {
	Base
}

/**
 * 获取数据库连接
 */
func (b *Base) GetDB(ctx context.Context, name ...string) *gorm.DB {
	if b.db == nil {
		b.SetDB(ctx, getDB(name...))
	}
	return b.db
}

func (b *Base) GetDefaultDB(ctx context.Context) *gorm.DB {
	return b.GetDB(ctx, sysConst.SysMysqlCfgKeyDefault)
}

/**
 * 设置数据库连接（仅当使用事物需要变更操作句柄时使用）
 */
func (b *Base) SetDB(ctx context.Context, db *gorm.DB) {
	b.db = SetSpanToGorm(ctx, db)
}


func (_ *BaseDao) IsNotFound(db *gorm.DB) bool {
	return errors.Is(db.Error, gorm.ErrRecordNotFound)
}