package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
	"usersvr/config"
	"usersvr/log"
)

//返回db对象

var (
	db     *gorm.DB
	dbOnce sync.Once
)

func initMysql() {
	cfg := config.GetGlobalConfig().DBConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.UserName, cfg.PassWord, cfg.Host, cfg.Port, cfg.DataBase)
	log.Infof("db conn %s:%d", cfg.Host, cfg.Port)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("open mysql err")
	}
	s, err := db.DB()
	if err != nil {
		panic("get db conn err")
	}
	s.SetMaxIdleConns(cfg.MaxIdleConn)                                 // 设置最大空闲连接
	s.SetMaxOpenConns(cfg.MaxOpenConn)                                 // 设置最大打开的连接
	s.SetConnMaxLifetime(time.Duration(cfg.MaxIdleTime) * time.Second) // 设置空闲时间为(s)
}

func GetDb() *gorm.DB {
	dbOnce.Do(initMysql)
	return db
}

func CloseDb() {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			panic("fetch db connection err:" + err.Error())
		}
		sqlDB.Close()
	}
}
