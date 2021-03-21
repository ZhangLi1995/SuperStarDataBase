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
