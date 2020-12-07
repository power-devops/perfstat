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
)

func DiskTotalStat() (*DiskTotal, error) {
	var disk C.perfstat_disk_total_t

	rc := C.perfstat_disk_total(nil, &disk, C.sizeof_perfstat_disk_total_t, 1)
	if rc != 1 {
		return nil, fmt.Errorf("perfstat_disk_total() error")
	}
	d := perfstatdisktotal2disktotal(disk)
	return &d, nil
}

func DiskAdapterStat() ([]DiskAdapter, error) {
	var adapter *C.perfstat_diskadapter_t
	var adptname C.perfstat_id_t

	numadpt := C.perfstat_diskadapter(nil, nil, C.sizeof_perfstat_diskadapter_t, 0)
	if numadpt <= 0 {
		return nil, fmt.Errorf("perfstat_diskadapter() error")
	}

	adapter_len := C.sizeof_perfstat_diskadapter_t * C.ulong(numadpt)
	adapter = (*C.perfstat_diskadapter_t)(C.malloc(adapter_len))
	C.strcpy(&adptname.name[0], C.CString(C.FIRST_DISKADAPTER))
	r := C.perfstat_diskadapter(&adptname, adapter, C.sizeof_perfstat_diskadapter_t, numadpt)
	if r < 0 {
		return nil, fmt.Errorf("perfstat_diskadapter() error")
	}
	da := make([]DiskAdapter, r)
	for i := 0; i < int(r); i++ {
		d := C.get_diskadapter_stat(adapter, C.int(i))
		if d != nil {
			da[i] = perfstatdiskadapter2diskadapter(d)
		}
	}
	return da, nil
}

func DiskStat() ([]Disk, error) {
	var disk *C.perfstat_disk_t
	var diskname C.perfstat_id_t

	numdisk := C.perfstat_disk(nil, nil, C.sizeof_perfstat_disk_t, 0)
	if numdisk <= 0 {
		return nil, fmt.Errorf("perfstat_disk() error")
	}

	disk_len := C.sizeof_perfstat_disk_t * C.ulong(numdisk)
	disk = (*C.perfstat_disk_t)(C.malloc(disk_len))
	C.strcpy(&diskname.name[0], C.CString(C.FIRST_DISK))
	r := C.perfstat_disk(&diskname, disk, C.sizeof_perfstat_disk_t, numdisk)
	if r < 0 {
		return nil, fmt.Errorf("perfstat_disk() error")
	}
	d := make([]Disk, r)
	for i := 0; i < int(r); i++ {
		ds := C.get_disk_stat(disk, C.int(i))
		if ds != nil {
			d[i] = perfstatdisk2disk(ds)
		}
	}
	return d, nil
}
