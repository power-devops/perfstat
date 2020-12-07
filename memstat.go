package perfstat

/*
#cgo LDFLAGS: -lperfstat

#include <libperfstat.h>
*/
import "C"

import "fmt"

func perfstatmemorytotal2memorytotal(n C.perfstat_memory_total_t) MemoryTotal {
	var m MemoryTotal
	m.VirtualTotal = int64(n.virt_total)
	m.RealTotal = int64(n.real_total)
	m.RealFree = int64(n.real_free)
	m.RealPinned = int64(n.real_pinned)
	m.RealInUse = int64(n.real_inuse)
	m.BadPages = int64(n.pgbad)
	m.PageFaults = int64(n.pgexct)
	m.PageIn = int64(n.pgins)
	m.PageOut = int64(n.pgouts)
	m.PgSpIn = int64(n.pgspins)
	m.PgSpOut = int64(n.pgspouts)
	m.Scans = int64(n.scans)
	m.Cycles = int64(n.cycles)
	m.PgSteals = int64(n.pgsteals)
	m.NumPerm = int64(n.numperm)
	m.PgSpTotal = int64(n.pgsp_total)
	m.PgSpFree = int64(n.pgsp_free)
	m.PgSpRsvd = int64(n.pgsp_rsvd)
	m.RealSystem = int64(n.real_system)
	m.RealUser = int64(n.real_user)
	m.RealProcess = int64(n.real_process)
	m.VirtualActive = int64(n.virt_active)
	m.IOME = int64(n.iome)
	m.IOMU = int64(n.iomu)
	m.IOHWM = int64(n.iohwm)
	m.PMem = int64(n.pmem)
	m.CompressedTotal = int64(n.comprsd_total)
	m.CompressedWSegPg = int64(n.comprsd_wseg_pgs)
	m.CPgIn = int64(n.cpgins)
	m.CPgOut = int64(n.cpgouts)
	m.TrueSize = int64(n.true_size)
	m.ExpandedMemory = int64(n.expanded_memory)
	m.CompressedWSegSize = int64(n.comprsd_wseg_size)
	m.TargetCPoolSize = int64(n.target_cpool_size)
	m.MaxCPoolSize = int64(n.max_cpool_size)
	m.MinUCPoolSize = int64(n.min_ucpool_size)
	m.CPoolSize = int64(n.cpool_size)
	m.UCPoolSize = int64(n.ucpool_size)
	m.CPoolInUse = int64(n.cpool_inuse)
	m.UCPoolInUse = int64(n.ucpool_inuse)
	m.Version = int64(n.version)
	m.RealAvailable = int64(n.real_avail)
	m.BytesCoalesced = int64(n.bytes_coalesced)
	m.BytesCoalescedMemPool = int64(n.bytes_coalesced_mempool)

	return m
}

func Memory() (*MemoryTotal, error) {
	var memory C.perfstat_memory_total_t

	rc := C.perfstat_memory_total(nil, &memory, C.sizeof_perfstat_memory_total_t, 1)
	if rc != 1 {
		return nil, fmt.Errorf("perfstat_memory_total() error")
	}
	m := perfstatmemorytotal2memorytotal(memory)
	return &m, nil
}
