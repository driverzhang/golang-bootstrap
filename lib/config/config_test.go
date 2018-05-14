package config

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	fmt.Println(C.Debug, C.Mysql.DsnRead, C.Mysql.DsnWrite, C.Mysql)
	t.Log(C.Debug)
}
