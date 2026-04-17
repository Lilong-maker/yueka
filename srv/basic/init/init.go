package init

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"yueka/srv/basic/config"
	"yueka/srv/handler/model"
)

func init() {
	ViperInit()
	MysqlInit()
	RedisInit()

}

var err error

func MysqlInit() {
	MysqlConfig := config.Gen.Mysql
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		MysqlConfig.User,
		MysqlConfig.Password,
		MysqlConfig.Host,
		MysqlConfig.Port,
		MysqlConfig.Database,
	)
	config.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("数据库连接成功")

	err := config.DB.AutoMigrate(&model.User{},
		&model.Miao{},
		&model.Goods{},
		&model.Order{},
		&model.OrderItem{},
	)
	if err != nil {
		return
	}
	fmt.Println("表迁移成功")

	sqlDB, _ := config.DB.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了可以重新使用连接的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func ViperInit() {
	viper.SetConfigFile("C:\\Users\\Lenovo\\Desktop\\yueka\\config.yml")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config.Gen)
	if err != nil {
		return
	}
	fmt.Println("配置文件加载成功")
}

var Ctx = context.Background()
var Rdb *redis.Client

func RedisInit() {
	RedisConfig := config.Gen.Redis
	Addr := fmt.Sprintf("%s:%d",
		RedisConfig.Host,
		RedisConfig.Port,
	)
	Rdb = redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: RedisConfig.Password, // no password set
		DB:       RedisConfig.Database, // use default DB
	})
	err := Rdb.Ping(Ctx).Err()
	if err != nil {
		return
	}
	fmt.Println("redis连接成功")
}
