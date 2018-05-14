package orm

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type RWEngine struct {
	Read  *xorm.Engine
	Write *xorm.Engine
}

// 支持读写分离，需要输入读写两个数据源
// 具体
func New(dsnRead, dsnWrite string) (engine *RWEngine, err error) {
	engine = &RWEngine{}
	engine.Read, err = xorm.NewEngine("mysql", dsnRead)
	if err != nil {
		return
	}
	engine.Write, err = xorm.NewEngine("mysql", dsnWrite)
	if err != nil {
		return
	}
	return
}
