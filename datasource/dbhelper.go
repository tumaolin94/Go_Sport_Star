package datasource

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"log"
	"superstar/conf"
	"sync"
)

var (
	masterEngine *xorm.Engine
	slaveEngine *xorm.Engine
	lock sync.Mutex
)

func InstanceMaster() *xorm.Engine {
	// TODO:
	if masterEngine != nil {
		return masterEngine
	}
	lock.Lock()
	defer lock.Unlock()

	if masterEngine != nil {
		return masterEngine
	}

	c := conf.MasterDbConfig
	driveSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		c.User, c.Pwd, c.Host, c.Port, c.DbName)

	engine, err := xorm.NewEngine(conf.DriverName, driveSource)

	if err != nil {
		log.Fatal("dbhelper.DbInstanceMaster error=", err)
		return nil
	} else {
		masterEngine = engine
		return masterEngine
	}

 }

func InstanceSlave() *xorm.Engine {
	// TODO:
	if slaveEngine != nil {
		return slaveEngine
	}
	lock.Lock()
	defer lock.Unlock()

	if slaveEngine != nil {
		return slaveEngine
	}

	c := conf.SlaveDbConfig
	driveSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		c.User, c.Pwd, c.Host, c.Port, c.DbName)

	engine, err := xorm.NewEngine(conf.DriverName, driveSource)

	if err != nil {
		log.Fatal("dbhelper.DbInstanceMaster error=", err)
		return nil
	} else {
		slaveEngine = engine
		return slaveEngine
	}
}