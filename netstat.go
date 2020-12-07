package perfstat

/*
#cgo LDFLAGS: -lperfstat

#include <libperfstat.h>
*/
import "C"

import (
	"fmt"
)

func NetIfaceTotalStat() (*NetIfaceTotal, error) {
	var nif C.perfstat_netinterface_total_t

	rc := C.perfstat_netinterface_total(nil, &nif, C.sizeof_perfstat_netinterface_total_t, 1)
	if rc != 1 {
		return nil, fmt.Errorf("perfstat_netinterface_total() error")
	}
	n := perfstatnetinterfacetotal2netifacetotal(nif)
	return &n, nil
}
