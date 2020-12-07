package perfstat

/*
#cgo LDFLAGS: -lperfstat

#include <libperfstat.h>
*/
import "C"

import "fmt"

func perfstatnetinterfacetotal2netifacetotal(n C.perfstat_netinterface_total_t) NetIfaceTotal {
	var i NetIfaceTotal

	i.Number = int32(n.number)
	i.IPackets = int64(n.ipackets)
	i.IBytes = int64(n.ibytes)
	i.IErrors = int64(n.ierrors)
	i.OPackets = int64(n.opackets)
	i.OBytes = int64(n.obytes)
	i.OErrors = int64(n.oerrors)
	i.Collisions = int64(n.collisions)
	i.XmitDrops = int64(n.xmitdrops)
	i.Version = int64(n.version)

	return i
}

func NetIfaceTotalStat() (*NetIfaceTotal, error) {
	var nif C.perfstat_netinterface_total_t

	rc := C.perfstat_netinterface_total(nil, &nif, C.sizeof_perfstat_netinterface_total_t, 1)
	if rc != 1 {
		return nil, fmt.Errorf("perfstat_netinterface_total() error")
	}
	n := perfstatnetinterfacetotal2netifacetotal(nif)
	return &n, nil
}
