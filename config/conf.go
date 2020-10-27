package config

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

// Init 初始化配置项
func init() {
	// 从本地读取环境变量
	godotenv.Load()

	// 连接数据库
	Database()

}
type Product struct {
	gorm.Model
	Code  string
	Price uint
}
// DB 数据库链接单例
var DB *gorm.DB

// Database 在中间件中初始化mysql链接
func Database() {
	dsn  := os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_DATABASE")+"?charset=utf8mb4"
	db, err := gorm.Open(mysql.Open(dsn ), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}
