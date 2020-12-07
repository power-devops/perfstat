package perfstat

/*
#cgo LDFLAGS: -lperfstat

#include <libperfstat.h>
*/
import "C"

import (
	"fmt"
)

func MemoryTotalStat() (*MemoryTotal, error) {
	var memory C.perfstat_memory_total_t

	rc := C.perfstat_memory_total(nil, &memory, C.sizeof_perfstat_memory_total_t, 1)
	if rc != 1 {
		return nil, fmt.Errorf("perfstat_memory_total() error")
	}
	m := perfstatmemorytotal2memorytotal(memory)
	return &m, nil
}
