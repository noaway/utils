package convert

import (
	"testing"
	"time"
)

func TestToFormat(t *testing.T) {
	t.Log(StrTo("2016/07/05").MustTime("2006/01/02"))
}

func TestToStr(t *testing.T) {
	t.Log(ToStr(time.Now()))
}

func TestToStrMd5(t *testing.T) {
	t.Log(StrTo("王阳").Md5())
}
