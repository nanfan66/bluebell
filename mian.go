package main

import (
	"bluebell01/controller"
	"bluebell01/dao/redis"
	"bluebell01/logger"
	"bluebell01/pkg/snowflake"
	"bluebell01/router"
	"bluebell01/setting"
	"bluebell01/dao/mysql"

	"fmt"
)

func main() {
	// 加载配置
	if err := setting.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	// 初始化日志
	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

	// 初始化数据库mysql
	if err := mysql.InitMysql(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.CloseMysql()

	// 初始化redis
	if err := redis.InitRedis(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.CloseRedis()


	// 初始化雪花id库
	if err := snowflake.InitSnowflake(setting.Conf.StartTime, setting.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}

	// 初始化gin框架内置的校验器使用的翻译器
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed, err:%v\n", err)
		return
	}
	// 注册路由
	r := router.SetUpRouter(setting.Conf.Mode)
	err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
