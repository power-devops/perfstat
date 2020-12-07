package perfstat

/*
#cgo LDFLAGS: -lperfstat

#include <libperfstat.h>
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

func DiskTotalStat() (*DiskTotal, error) {
	var disk C.perfstat_disk_total_t

	rc := C.perfstat_disk_total(nil, &disk, C.sizeof_perfstat_disk_total_t, 1)
	if rc != 1 {
		return nil, fmt.Errorf("perfstat_disk_total() error")
	}
	d := perfstatdisktotal2disktotal(disk)
	return &d, nil
}
