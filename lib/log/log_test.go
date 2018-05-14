package log

import (
	"testing"
	"os"
)

func TestLog(t *testing.T) {
	SetDebug(false)
	Error("hello")
	Print("hello")
	Printf("t %+v", os.Args)
	WithField("a", 1).Error("hello")

	SetDebug(true)
	Error("hello")
	Print("hello")
	Printf("t %+v", os.Args)
	WithField("a", 1).Error("hello")
}
