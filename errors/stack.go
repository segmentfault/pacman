package errors

import (
	"bytes"
	"fmt"
	"runtime"
)

// LogStack return call function stack info from start stack to end stack.
// if end is a positive number, return all call function stack.
func LogStack(start, end int) string {
	stack := bytes.Buffer{}
	for i := start; i < end || end <= 0; i++ {
		pc, str, line, _ := runtime.Caller(i)
		if line == 0 {
			break
		}
		stack.WriteString(fmt.Sprintf("%s:%d %s\n", str, line, runtime.FuncForPC(pc).Name()))
	}
	return stack.String()
}
