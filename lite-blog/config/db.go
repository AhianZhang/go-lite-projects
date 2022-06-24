package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	Id         int
	Name       string
	Gender     int
	Address    string
	CreateTime time.Time
}

func main() {
	// 配置日志
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别 Silent、Error、Warn、Info
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,       // 禁用彩色打印
		},
	)
	dsn := "root:123456@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	// 创建
	// user := User{Name: "ahian", Gender: 1, Address: "北京", CreateTime: time.Now()}
	// db.Select("name", "gender", "address", "create_time").Create(&user)
	// fmt.Println(user.Id)
	// 批量创建
	var userBatch []User

	for i := 0; i < 3; i++ {
		userBatch = append(userBatch, User{Name: "ahian" + strconv.Itoa(i), Gender: i % 2, Address: "北京", CreateTime: time.Now()})
	}
	// 一起创建
	db.Select("name", "gender", "address", "create_time").Create(&userBatch)
	// 分批创建
	db.Select("name", "gender", "address", "create_time").CreateInBatches(&userBatch, 1)

	var user1 User
	// 查询第一个
	db.First(&user1)
	fmt.Println(user1)
	// 查询指定个数
	var users []User
	db.Limit(10).Find(&users)
	fmt.Println(users)

}

//======= 配置表名
type tabler interface {
	TableName() string
}

func (User) TableName() string {
	return "t_user"
}

//======= hooks
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("before create hook say hi")
	return
}
