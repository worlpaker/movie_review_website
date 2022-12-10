package ErrHandler

import (
	"log"
	"runtime"
)
//Log to error in runtime
func Log(err error) (ok bool) {
	if err != nil {
		pc, filename, line, _ := runtime.Caller(1)
		log.Printf("[Error] in %s[%s:%d] || Error: %v \n", runtime.FuncForPC(pc).Name(), filename, line, err)
		ok = true
		return
	}
	return
}
