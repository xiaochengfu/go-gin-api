package remind_server

import (
	"strings"

	"github.com/spf13/cast"
)

func (*server) ConvSecond(t string) int64 {
	strArr := strings.Split(t, ":")
	var num, s int64
	for i := 0; i < len(strArr); i++ {
		num = cast.ToInt64(strArr[i])
		if i == 0 {
			s += num * 60 * 60
		}
		if i == 1 {
			s += num * 60
		}
		if i == 2 {
			s += num
		}
	}
	return s
}
