package perfstat

/*
#cgo LDFLAGS: -lperfstat

#include <libperfstat.h>
#include <string.h>
#include "c_helpers.h"
*/
import "C"

import "fmt"

func perfstatdisktotal2disktotal(n C.perfstat_disk_total_t) DiskTotal {
	var d DiskTotal

	d.Number = int32(n.number)
	d.Size = int64(n.size)
	d.Free = int64(n.free)
	d.XRate = int64(n.xrate)
	d.Xfers = int64(n.xfers)
	d.Wblks = int64(n.wblks)
	d.Rblks = int64(n.rblks)
	d.Time = int64(n.time)
	d.Version = int64(n.version)
	d.Rserv = int64(n.rserv)
	d.MinRserv = int64(n.min_rserv)
	d.MaxRserv = int64(n.max_rserv)
	d.RTimeOut = int64(n.rtimeout)
	d.RFailed = int64(n.rfailed)
	d.Wserv = int64(n.wserv)
	d.MinWserv = int64(n.min_wserv)
	d.MaxWserv = int64(n.max_wserv)
	d.WTimeOut = int64(n.wtimeout)
	d.WFailed = int64(n.wfailed)
	d.WqDepth = int64(n.wq_depth)
	d.WqTime = int64(n.wq_time)
	d.WqMinTime = int64(n.wq_min_time)
	d.WqMaxTime = int64(n.wq_max_time)

	return d
}

func perfstatdiskadapter2diskadapter(n *C.perfstat_diskadapter_t) DiskAdapter {
	var d DiskAdapter

	d.Name = C.GoString(&n.name[0])
	d.Description = C.GoString(&n.description[0])
	d.Number = int32(n.number)
	d.Size = int64(n.size)
	d.Free = int64(n.free)
	d.XRate = int64(n.xrate)
	d.Xfers = int64(n.xfers)
	d.Rblks = int64(n.rblks)
	d.Wblks = int64(n.wblks)
	d.Time = int64(n.time)
	d.Version = int64(n.version)
	d.AdapterType = int64(n.adapter_type)
	d.DkBSize = int64(n.dk_bsize)
	d.DkRserv = int64(n.dk_rserv)
	d.DkWserv = int64(n.dk_wserv)
	d.MinRserv = int64(n.min_rserv)
	d.MaxRserv = int64(n.max_rserv)
	d.MinWserv = int64(n.min_wserv)
	d.MaxWserv = int64(n.max_wserv)
	d.WqDepth = int64(n.wq_depth)
	d.WqSampled = int64(n.wq_sampled)
	d.WqTime = int64(n.wq_time)
	d.WqMinTime = int64(n.wq_min_time)
	d.WqMaxTime = int64(n.wq_max_time)
	d.QFull = int64(n.q_full)
	d.QSampled = int64(n.q_sampled)

	return d
}

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
