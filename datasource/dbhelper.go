package datasource

import (
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/wangjian890523/superstar/conf"
)

var (
	masterEngine *xorm.Engine
	slaveEngine  *xorm.Engine
	lock         sync.Mutex
)

func InstanceMaster() *xorm.Engine {
	if masterEngine != nil {
		return masterEngine
	}

	lock.Lock()
	defer lock.Unlock()

	if masterEngine != nil {
		return masterEngine
	}

	c := conf.MasterDbConf
	driveSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", 
	c.User, c.Pwd, c.Host, c.Port, c.DbName)

	fmt.Println("db:", driveSource)
	engine, err := xorm.NewEngine(conf.DriverName, driveSource)
	if err != nil {
		log.Fatal("dbhelper.DbInstanceMaster error=", err)
		return nil
	}

	engine.ShowSQL(false)
	engine.SetTZLocation(conf.SysTimeLocation)

	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	engine.SetDefaultCacher(cacher)

	masterEngine = engine
	return masterEngine

}

func InstanceSlave() *xorm.Engine {

	if slaveEngine != nil {
		return slaveEngine
	}

	lock.Lock()
	defer lock.Unlock()

	if slaveEngine != nil {
		return slaveEngine
	}

	c := conf.SlaveDbConf
	driveSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", c.User,
		c.Pwd, c.Host, c.Port, c.DbName)

	engine, err := xorm.NewEngine(conf.DriverName, driveSource)
	if err != nil {
		log.Fatal("dbhelper.DbInstanceSlave error=", err)
		return nil
	}

	engine.SetTZLocation(conf.SysTimeLocation)
	slaveEngine = engine
	return slaveEngine

}
