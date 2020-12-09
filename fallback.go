// +build !aix

package perfstat

import (
	"fmt"
	"time"
)

func EnableLVMStat()  {}
func DisableLVMStat() {}

func CpuStat() ([]CPU, error) {
	return nil, fmt.Errorf("not implemented")
}

func CpuTotalStat() (*CPUTotal, error) {
	return nil, fmt.Errorf("not implemented")
}

func CpuUtilStat(intvl time.Duration) (*CPUUtil, error) {
	return nil, fmt.Errorf("not implemented")
}

func DiskTotalStat() (*DiskTotal, error) {
	return nil, fmt.Errorf("not implemented")
}

func DiskAdapterStat() ([]DiskAdapter, error) {
	return nil, fmt.Errorf("not implemented")
}

func DiskStat() ([]Disk, error) {
	return nil, fmt.Errorf("not implemented")
}

func DiskPathStat() ([]DiskPath, error) {
	return nil, fmt.Errorf("not implemented")
}

func FCAdapterStat() ([]FCAdapter, error) {
	return nil, fmt.Errorf("not implemented")
}

func PartitionStat() (*PartitionConfig, error) {
	return nil, fmt.Errorf("not implemented")
}

func LogicalVolumeStat() ([]LogicalVolume, error) {
	return nil, fmt.Errorf("not implemented")
}

func VolumeGroupStat() ([]VolumeGroup, error) {
	return nil, fmt.Errorf("not implemented")
}

func MemoryTotalStat() (*MemoryTotal, error) {
	return nil, fmt.Errorf("not implemented")
}

func MemoryPageStat() ([]MemoryPage, error) {
	return nil, fmt.Errorf("not implemented")
}

func PagingSpaceStat() ([]PagingSpace, error) {
	return nil, fmt.Errorf("not implemented")
}

func NetIfaceTotalStat() (*NetIfaceTotal, error) {
	return nil, fmt.Errorf("not implemented")
}

func NetBufferStat() ([]NetBuffer, error) {
	return nil, fmt.Errorf("not implemented")
}

func NetIfaceStat() ([]NetIface, error) {
	return nil, fmt.Errorf("not implemented")
}

func NetAdapterStat() ([]NetAdapter, error) {
	return nil, fmt.Errorf("not implemented")
}

func ProcessStat() ([]Process, error) {
	return nil, fmt.Errorf("not implemented")
}

func ThreadStat() ([]Thread, error) {
	return nil, fmt.Errorf("not implemented")
}

func Sysconf(name int32) (int64, error) {
	return 0, fmt.Errorf("not implemented")
}

func GetCPUImplementation() string {
	return ""
}

func POWER9OrNewer() bool {
	return false
}

func POWER9() bool {
	return false
}

func POWER8OrNewer() bool {
	return false
}

func POWER8() bool {
	return false
}

func POWER7OrNewer() bool {
	return false
}

func POWER7() bool {
	return false
}

func HasTransactionalMemory() bool {
	return false
}

func Is64Bit() bool {
	return false
}

func IsSMP() bool {
	return false
}

func HasVMX() bool {
	return false
}

func HasVSX() bool {
	return false
}

func HasDFP() bool {
	return false
}

func HasNxGzip() bool {
	return false
}

func PksCapable() bool {
	return false
}

func PksEnabled() bool {
	return false
}

func CPUMode() string {
	return ""
}

func KernelBits() int {
	return 0
}

func IsLPAR() bool {
	return false
}

func CpuAddCapable() bool {
	return false
}

func CpuRemoveCapable() bool {
	return false
}

func MemoryAddCapable() bool {
	return false
}

func MemoryRemoveCapable() bool {
	return false
}

func DLparCapable() bool {
	return false
}

func IsNUMA() bool {
	return false
}

func KernelKeys() bool {
	return false
}

func RecoveryMode() bool {
	return false
}

func EnhancedAffinity() bool {
	return false
}

func VTpmEnabled() bool {
	return false
}

func IsVIOS() bool {
	return false
}

func MLSEnabled() bool {
	return false
}

func SPLparCapable() bool {
	return false
}

func SPLparEnabled() bool {
	return false
}

func DedicatedLpar() bool {
	return false
}

func SPLparCapped() bool {
	return false
}

func SPLparDonating() bool {
	return false
}

func SmtCapable() bool {
	return false
}

func SmtEnabled() bool {
	return false
}

func VrmCapable() bool {
	return false
}

func VrmEnabled() bool {
	return false
}

func AmeEnabled() bool {
	return false
}

func EcoCapable() bool {
	return false
}

func EcoEnabled() bool {
	return false
}
