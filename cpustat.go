package perfstat

/*
#cgo LDFLAGS: -lperfstat

#include <libperfstat.h>
#include <string.h>

#include "c_helpers.h"
*/
import "C"

import (
	"fmt"
	"runtime"
)

func Cpu() ([]CPU, error) {
	var cpustat *C.perfstat_cpu_t
	var cpu C.perfstat_id_t

	ncpu := runtime.NumCPU()

	cpustat_len := C.sizeof_perfstat_cpu_t * C.ulong(ncpu)
	cpustat = (*C.perfstat_cpu_t)(C.malloc(cpustat_len))
	C.strcpy(&cpu.name[0], C.CString(C.FIRST_CPU))
	r := C.perfstat_cpu(&cpu, cpustat, C.sizeof_perfstat_cpu_t, C.int(ncpu))
	if r <= 0 {
		return nil, fmt.Errorf("error perfstat_cpu()")
	}
	c := make([]CPU, r)
	for i := 0; i < int(r); i++ {
		n := C.get_cpu_stat(cpustat, C.int(i))
		if n != nil {
			c[i] = perfstatcpu2cpu(n)
		}
	}
	return c, nil
}

func CpuTotal() (*CPUTotal, error) {
	var cpustat *C.perfstat_cpu_total_t

	cpustat = (*C.perfstat_cpu_total_t)(C.malloc(C.sizeof_perfstat_cpu_total_t))
	r := C.perfstat_cpu_total(nil, cpustat, C.sizeof_perfstat_cpu_total_t, 1)
	if r <= 0 {
		return nil, fmt.Errorf("error perfstat_cpu_total()")
	}
	c := perfstatcputotal2cputotal(cpustat)
	return &c, nil
}
