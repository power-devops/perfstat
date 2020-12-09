// +build aix

package perfstat

/*
#cgo LDFLAGS: -lperfstat

#include <libperfstat.h>
*/
import "C"

// EnableLVMStat() switches on LVM (logical volumes and volume groups) performance statistics.
// With this enabled you can use fields KBReads, KBWrites, and IOCnt
// in LogicalVolume and VolumeGroup data types.
func EnableLVMStat() {
	C.perfstat_config(C.PERFSTAT_ENABLE|C.PERFSTAT_LV|C.PERFSTAT_VG, nil)
}

// DisableLVMStat() switchess of LVM (logical volumes and volume groups) performance statistics.
// This is the default state. In this case LogicalVolume and VolumeGroup data types are
// populated with informations about LVM structures, but performance statistics fields
// (KBReads, KBWrites, IOCnt) are empty.
func DisableLVMStat() {
	C.perfstat_config(C.PERFSTAT_DISABLE|C.PERFSTAT_LV|C.PERFSTAT_VG, nil)
}
