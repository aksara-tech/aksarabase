package s

import "fmt"

//F is short function for fmt.Sprintf
func F(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}
