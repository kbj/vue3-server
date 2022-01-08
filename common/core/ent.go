package core

import (
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"vue3-server/common/global"
	"vue3-server/ent"
)

// InitializeEntInstance 创建Ent对象
func InitializeEntInstance() *ent.Client {
	driver, err := getDbDriver()
	if err != nil {
		global.Logger.Panic("获取数据库连接失败，失败原因：" + err.Error())
		return nil
	}

	// 获取数据库驱动中的sql.DB对象。
	db := driver.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	return ent.NewClient(ent.Driver(driver))
}

// 得到数据库的连接驱动
func getDbDriver() (*sql.Driver, error) {
	var driver *sql.Driver
	var err error

	switch global.Config.System.DbType {
	case dialect.MySQL:
		global.Logger.Info("当前数据库的类型配置为：MySQL")
		dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
			global.Config.System.Db.UserName,
			global.Config.System.Db.Password,
			global.Config.System.Db.Host,
			global.Config.System.Db.Port,
			global.Config.System.Db.DbName,
			global.Config.System.Db.Config,
		)
		driver, err = sql.Open(dialect.MySQL, dbUrl)
		break
	default:
		// 默认是sqlite
		global.Logger.Info("数据库类型未配置，默认配置为：SQLite3")
		dbUrl := fmt.Sprintf("file:%s?mode=memory&cache=shared&_fk=1", global.Config.System.Db.DbName)
		driver, err = sql.Open(dialect.SQLite, dbUrl)
	}
	return driver, err
}
