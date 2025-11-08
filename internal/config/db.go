package config

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := Get().MySQLDSN
	var err error
	// 注意:=形式，会导致全局变量DB未正确赋值
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, //禁止迁移期间创建表时使用外键约束
	})

	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	// 设置数据库连接池
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("获取底层连接失败: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("数据库连接成功")
}

// CloseDB 数据关闭
func CloseDB() {
	if DB == nil {
		return
	}

	sqlDB, err := DB.DB()

	if err != nil {
		log.Printf("获取底层连接失败: %v", err)
		return
	}

	if err := sqlDB.Close(); err != nil {
		log.Printf("关闭数据库连接失败: %v", err)
	} else {
		fmt.Println("数据库连接已关闭")
	}
}
