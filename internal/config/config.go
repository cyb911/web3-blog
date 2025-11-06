package config

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv   string
	AppPort  string
	MySQLDSN string
}

var (
	cfg  *Config
	once sync.Once
)

func MustLoad() *Config {
	once.Do(func() {
		envPath := findEnvFile()
		if envPath != "" {
			err := godotenv.Load(envPath)
			if err != nil {
				log.Printf("加载 .env 文件失败: %v", err)
			} else {
				log.Printf("加载 .env 文件成功:%s", envPath)
			}
		} else {
			log.Println("未找到 .env 文件，使用系统环境变量")
		}
	})

	cfg = &Config{
		AppEnv:   getEnvDefault("APP_ENV", "dev"),
		AppPort:  getEnvDefault("APP_PORT", "8080"),
		MySQLDSN: os.Getenv("MYSQL_DSN"),
	}

	if cfg.MySQLDSN == "" {
		log.Fatal("配置错误：缺少 MYSQL_DSN")
	}

	return cfg
}

func Get() *Config {
	if cfg == nil {
		return MustLoad()
	}

	return cfg
}

// findEnvFile 从当前目录向上查找 .env 文件
func findEnvFile() string {
	dir, _ := os.Getwd()

	for i := 0; i < 6; i++ {
		encPath := filepath.Join(dir, ".env")
		if _, err := os.Stat(encPath); err == nil {
			return encPath
		}

		dir = filepath.Dir(dir)
	}
	return ""
}

func getEnvDefault(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
