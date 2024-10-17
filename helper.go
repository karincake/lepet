package lepet

import "fmt"

func processMsg(msg string, opt ...string) string {
	if len(opt) > 0 {
		var optAny []any = []any{}
		for key := range opt {
			optAny = append(optAny, opt[key])
		}
		return fmt.Sprintf(msg, optAny...)
	} else {
		return msg
	}
}
