package datasource

import (
	"fmt"
	"log"

	"test_gin/common/setting"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	//
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
)

//Db gormDB
type Db struct {
	Conn *gorm.DB
}

type IDb interface {
	//Connect 初始化数据库配置
	Connect() error
	//DB 返回DB
	DB() *gorm.DB
}

func NewDB() IDb {
	return &Db{}
}

//Connect 初始化数据库配置
func (d *Db) Connect() error {
	var (
		dbName, user, pwd, host string
	)

	conf := setting.Config.Database
	dbName = conf.Name
	user = conf.User
	pwd = conf.Password
	host = conf.Host

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pwd, host, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   conf.TablePrefix,
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("connecting mysql error: ", err)
		return err
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	d.Conn = db

	log.Println("Connect Mysql Success")

	return nil
}

//DB 返回DB
func (d *Db) DB() *gorm.DB {
	return d.Conn
}
