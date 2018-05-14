package orm

import (
	"testing"
)

func TestRead(t *testing.T) {
	var job []struct {
		Id int `json:"id"`
	}
	err := engine.Read.Table("meizitu_list").Find(&job)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", job)
}

// 实际在编写model层的时候，应该在model包里init这个engine
var engine *RWEngine

func init() {
	dsn := "root:1992@tcp(localhost:3306)/nodezhang"
	var err error
	engine, err = New(dsn, dsn)
	if err != nil {
		panic(err)
	}
}
