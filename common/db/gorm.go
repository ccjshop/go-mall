// Package db grom初始化
package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type DBOption func(*gorm.DB) *gorm.DB
type DBOrderOption func(*gorm.DB) *gorm.DB

// GetConn 创建gorm
func GetConn(username string, password string, host string, port uint32, timeout string, dbName string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s",
		username, password, host, port, dbName, timeout)

	// 全局日志设置
	mysqlLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // 日志输出的目标，前缀和日志包含的内容
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // 日志级别
			Colorful:      true,        // 使用彩色打印
		},
	)

	// 连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 为了确保数据一致性，GORM 会在事务里执行写入操作（创建、更新、删除）。如果没有这方面的要求，您可以在初始化时禁用它，这样可以获得60%的性能提升
		SkipDefaultTransaction: true,
		// gorm采用的命名策略是，表名是蛇形复数，字段名是蛇形单数
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",    // 表名前缀
			SingularTable: true,  // 单数表名
			NoLowerCase:   false, // 关闭小写转换
		},
		// 自定义日志
		Logger: mysqlLogger,
	})
	if err != nil {
		return nil, err
	}

	return conn, nil
}
