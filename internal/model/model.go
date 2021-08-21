package model

import (
	"time"

	"github.com/smallpaes/go-blog-backend/global"
	"github.com/smallpaes/go-blog-backend/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/soft_delete"

	"fmt"
)

type Model struct {
	ID         uint32                `gorm:"primary_key" json:"id"`
	CreatedBy  string                `json:"created_by"`
	ModifiedBy string                `json:"modified_by"`
	CreatedOn  uint32                `json:"created_on"`
	ModifiedOn uint32                `json:"modified_on"`
	DeletedAt  int                   `json:"deleted_at"`
	IsDel      soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt"`
}

// create db instance with config
func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(loggerLevel(global.ServerSetting.RunMode)),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name
		},
		NowFunc: func() time.Time {
			return time.Unix(time.Now().Unix(), 0)
		},
	})

	if err != nil {
		return nil, err
	}
	fmt.Println("init")
	db.Callback().Create().Before("gorm:create").Register("update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Before("gorm:update").Register("update_time_stamp", updateTimeStampForUpdateCallback)

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	sqlDB.SetMaxOpenConns(databaseSetting.MaxOpenConns)
	return db, nil
}

// Scope contain current operation's information when performing
// any operation on the database
func updateTimeStampForCreateCallback(db *gorm.DB) {
	if db.Error == nil {
		nowTime := time.Now().Unix()
		createdTimeField := db.Statement.Schema.LookUpField("CreatedOn")
		_, isZero := createdTimeField.ValueOf(db.Statement.ReflectValue)
		if isZero {
			_ = createdTimeField.Set(db.Statement.ReflectValue, nowTime)
		}

		modifiedTimeField := db.Statement.Schema.LookUpField("ModifiedOn")
		_, isZero = modifiedTimeField.ValueOf(db.Statement.ReflectValue)
		if isZero {
			fmt.Println(modifiedTimeField, db.Statement.ReflectValue)
			_ = modifiedTimeField.Set(db.Statement.ReflectValue, nowTime)
		}
	}
}

func updateTimeStampForUpdateCallback(db *gorm.DB) {
	if db.Error == nil {
		nowTime := time.Now().Unix()
		modifiedTimeField := db.Statement.Schema.LookUpField("ModifiedOn")
		_ = modifiedTimeField.Set(db.Statement.ReflectValue, nowTime)
	}
}

func loggerLevel(m string) logger.LogLevel {
	if m == "debug" {
		return logger.Warn
	} else {
		return logger.Silent
	}
}
