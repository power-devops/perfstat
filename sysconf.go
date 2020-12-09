// +build aix

package perfstat

/*
#include <unistd.h>
*/
import "C"

import "fmt"

const (
	SC_ARG_MAX   = 0
	SC_CHILD_MAX = 1
	SC_CLK_TCK   = 2
)

func Sysconf(name int32) (int64, error) {
	r := C.sysconf(C.int(name))
	if r == -1 {
		return 0, fmt.Errorf("sysconf error")
	} else {
		return int64(r), nil
	}
}
