package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"request-example/config"
	"time"
)

var db *gorm.DB

type Model struct {
	ID        int `json:"id" gorm:"primary_key"`
	CreatedAt int `json:"created_at"`
	UpdatedAt int `json:"updated_at"`
}

func InitDb() {
	var (
		err error
	)

	db, err = gorm.Open(sqlite.Open(config.DbConf.Name), &gorm.Config{
		SkipDefaultTransaction: true, //为了确保数据一致性，GORM 会在事务里执行写入操作（创建、更新、删除）。如果没有这方面的要求，您可以在初始化时禁用它，这将获得大约 30%+ 性能提升。
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.DbConf.TablePrefix, //表名前缀，`User` 的表名应该是`d_users`
			SingularTable: true,                      //使用单数表名，启用该选项，此时，`User` 的表名应该是`d_user`
			NameReplacer:  nil,
			NoLowerCase:   false,
		},
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: false, //执行任何 SQL 时都创建 prepared statement 并缓存，可以提高后续的调用速度
	})
	if err != nil {
		log.Println(err)
	}

	sqlDB, err2 := db.DB()
	if err2 != nil {
		log.Println(err2)
	}
	sqlDB.SetMaxIdleConns(config.DbConf.SetMaxIdleConns) //最大空闲连接数
	sqlDB.SetMaxOpenConns(config.DbConf.SetMaxOpenConns) //最大连接数
	sqlDB.SetConnMaxLifetime(time.Hour)                  //设置了连接可复用的最大时间

	db.AutoMigrate(&New{})
}
