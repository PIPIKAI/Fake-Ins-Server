package common

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/PIPIKAI/Ins-gin-vue/server/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	host := viper.GetString("datasource.url")
	port := viper.GetString("datasource.port")
	dbname := viper.GetString("datasource.dbname")
	charset := viper.GetString("datasource.charset")
	local := viper.GetString("datasource.loc")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s", username, password, host, port, dbname, charset, url.QueryEscape(local))

	log.Println("dsn:", dsn)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn, // DSN data source name
		DefaultStringSize: 171, // string 类型字段的默认长度
	}), &gorm.Config{
		// 连接时候的配置
		SkipDefaultTransaction: false, // 跳过默认事物
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",    // 表的前缀
			SingularTable: false, // 单数表名
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 逻辑外键
	})

	// 数据库连接池的设置
	mysqlDB, _ := db.DB()

	mysqlDB.SetMaxIdleConns(10)  //设置空闲连接池中连接的最大数
	mysqlDB.SetMaxOpenConns(100) // 数据库连接最大数量

	mysqlDB.SetConnMaxLifetime(time.Hour) // 设置连接最大可复用时间

	if err != nil {
		panic("链接数据库失败:" + err.Error())
	}
	log.Println("链接数据库成功！")
	db.AutoMigrate(
		&model.User{},
		&model.Post{},
		&model.ImgUrl{},
		&model.Comment{},
		&model.Category{},
		&model.Collection{},
		&model.Like{},
		&model.DisLike{},
	)
	DB = db
	return db
}
func GetDB() *gorm.DB {
	return DB
}
