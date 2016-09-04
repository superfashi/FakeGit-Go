package fakegit

// #include <stdlib.h>
import "C"
import (
	"fmt"
	"os"
)

func IsIn(s string, n []string) bool {
	for _, k := range n {
		if s == k {
			return true
		}
	}
	return false
}

func IndexOf(s string, n []string) int {
	for i, k := range n {
		if s == k {
			return i
		}
	}
	return -1
}

func Pop(s int, n *[]string) string {
	ret := (*n)[s]
	(*n) = append((*n)[:s], (*n)[s+1:]...)
	return ret
}

func Fatal(v ...interface{}) {
	fmt.Println(v...)
	os.Exit(1)
}

func RunCommand(argx string) {
	C.system(C.CString(argx))
}
