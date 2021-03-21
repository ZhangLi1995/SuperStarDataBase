package datasource

import (
	"fmt"
	"log"
	"superstar/conf"
	"sync"

	"github.com/go-xorm/xorm"
)

var (
	masterEngine *xorm.Engine
	slaveEngine  *xorm.Engine
	once         sync.Once
)

func InstanceMaster() *xorm.Engine {
	once.Do(func() {
		dbConfig := conf.MasterDbConfig
		driveSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4",
			dbConfig.User, dbConfig.Pwd, dbConfig.Host, dbConfig.Port, dbConfig.DbName)
		engine, err := xorm.NewEngine(conf.DriverName, driveSource)
		if err != nil {
			log.Fatalf("[InstanceMaster] failed. err: %v", err)
			return
		}
		// Debug模式，打印全部的SQL语句，帮助对比，看ORM与SQL执行的对照关系
		engine.ShowSQL(false)
		engine.SetTZLocation(conf.SysTimeLocation)

		// 性能优化的时候才考虑，加上本机的SQL缓存
		// 增加缓存后，QPS 达到：9320
		cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
		engine.SetDefaultCacher(cacher)

		masterEngine = engine
	})
	return masterEngine
}

func InstanceSlave() *xorm.Engine {
	once.Do(func() {
		dbConfig := conf.SlaveDbConfig
		driveSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4",
			dbConfig.User, dbConfig.Pwd, dbConfig.Host, dbConfig.Port, dbConfig.DbName)
		engine, err := xorm.NewEngine(conf.DriverName, driveSource)
		if err != nil {
			log.Fatalf("[InstanceSlave] failed. err: %v", err)
			return
		}
		slaveEngine = engine
	})
	return slaveEngine
}
