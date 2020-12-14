package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part1/user-service/basic/config"
	log "github.com/micro/go-micro/v2/logger"
	"time"
)

func initMysql() {
	var err error

	// 创建连接
	mysqlDB, err := sql.Open("mysql", config.GetMysqlConfig().GetURL())
	if err != nil {
		log.Fatal(err) //有异常时打印信息，调用os.exit(1)退出应用程序,不执行defer方法
		panic(err)//有异常时，停止执行当前函数，并返回（如果有defer，就会执行defer方法）
	}

	mysqlDB.SetMaxOpenConns(config.GetMysqlConfig().GetMaxOpenConnection())
	mysqlDB.SetMaxIdleConns(config.GetMysqlConfig().GetMaxIdleConnection())

	mysqlDB.SetConnMaxLifetime(time.Second * time.Duration(config.GetMysqlConfig().GetConnMaxLifetime()))
	if err = mysqlDB.Ping(); err != nil {
		log.Fatal(err)
	}
}
