package perfstat

/*
#cgo LDFLAGS: -lperfstat

#include <libperfstat.h>

#include "c_helpers.h"
*/
import "C"

import "fmt"

func perfstatpartitionconfig2partitionconfig(n C.perfstat_partition_config_t) PartitionConfig {
	var p PartitionConfig
	p.Version = int64(n.version)
	p.Name = C.GoString(&n.partitionname[0])
	p.Node = C.GoString(&n.nodename[0])
	p.Conf.SmtCapable = (n.conf[0] & (1 << 7)) > 0
	p.Conf.SmtEnabled = (n.conf[0] & (1 << 6)) > 0
	p.Conf.LparCapable = (n.conf[0] & (1 << 5)) > 0
	p.Conf.LparEnabled = (n.conf[0] & (1 << 4)) > 0
	p.Conf.SharedCapable = (n.conf[0] & (1 << 3)) > 0
	p.Conf.SharedEnabled = (n.conf[0] & (1 << 2)) > 0
	p.Conf.DLparCapable = (n.conf[0] & (1 << 1)) > 0
	p.Conf.Capped = (n.conf[0] & (1 << 0)) > 0
	p.Conf.Kernel64bit = (n.conf[1] & (1 << 7)) > 0
	p.Conf.PoolUtilAuthority = (n.conf[1] & (1 << 6)) > 0
	p.Conf.DonateCapable = (n.conf[1] & (1 << 5)) > 0
	p.Conf.DonateEnabled = (n.conf[1] & (1 << 4)) > 0
	p.Conf.AmsCapable = (n.conf[1] & (1 << 3)) > 0
	p.Conf.AmsEnabled = (n.conf[1] & (1 << 2)) > 0
	p.Conf.PowerSave = (n.conf[1] & (1 << 1)) > 0
	p.Conf.AmeEnabled = (n.conf[1] & (1 << 0)) > 0
	p.Conf.SharedExtended = (n.conf[2] & (1 << 7)) > 0
	p.Number = int32(n.partitionnum)
	p.GroupID = int32(n.groupid)
	p.ProcessorFamily = C.GoString(&n.processorFamily[0])
	p.ProcessorModel = C.GoString(&n.processorModel[0])
	p.MachineID = C.GoString(&n.machineID[0])
	p.ProcessorMhz = float64(C.get_partition_mhz(n))
	p.NumProcessors.Online = int64(n.numProcessors.online)
	p.NumProcessors.Max = int64(n.numProcessors.max)
	p.NumProcessors.Min = int64(n.numProcessors.min)
	p.NumProcessors.Desired = int64(n.numProcessors.desired)
	p.OSName = C.GoString(&n.OSName[0])
	p.OSVersion = C.GoString(&n.OSVersion[0])
	p.OSBuild = C.GoString(&n.OSBuild[0])
	p.LCpus = int32(n.lcpus)
	p.SmtThreads = int32(n.smtthreads)
	p.Drives = int32(n.drives)
	p.NetworkAdapters = int32(n.nw_adapters)
	p.CpuCap.Online = int64(n.cpucap.online)
	p.CpuCap.Max = int64(n.cpucap.max)
	p.CpuCap.Min = int64(n.cpucap.min)
	p.CpuCap.Desired = int64(n.cpucap.desired)
	p.Weightage = int32(n.cpucap_weightage)
	p.EntCapacity = int32(n.entitled_proc_capacity)
	p.VCpus.Online = int64(n.vcpus.online)
	p.VCpus.Max = int64(n.vcpus.max)
	p.VCpus.Min = int64(n.vcpus.min)
	p.VCpus.Desired = int64(n.vcpus.desired)
	p.PoolID = int32(n.processor_poolid)
	p.ActiveCpusInPool = int32(n.activecpusinpool)
	p.PoolWeightage = int32(n.cpupool_weightage)
	p.SharedPCpu = int32(n.sharedpcpu)
	p.MaxPoolCap = int32(n.maxpoolcap)
	p.EntPoolCap = int32(n.entpoolcap)
	p.Mem.Online = int64(n.mem.online)
	p.Mem.Max = int64(n.mem.max)
	p.Mem.Min = int64(n.mem.min)
	p.Mem.Desired = int64(n.mem.desired)
	p.MemWeightage = int32(n.mem_weightage)
	p.TotalIOMemoryEntitlement = int64(n.totiomement)
	p.MemPoolID = int32(n.mempoolid)
	p.HyperPgSize = int64(n.hyperpgsize)
	p.ExpMem.Online = int64(n.exp_mem.online)
	p.ExpMem.Max = int64(n.exp_mem.max)
	p.ExpMem.Min = int64(n.exp_mem.min)
	p.ExpMem.Desired = int64(n.exp_mem.desired)
	p.TargetMemExpFactor = int64(n.targetmemexpfactor)
	p.TargetMemExpSize = int64(n.targetmemexpsize)
	p.SubProcessorMode = int32(n.subprocessor_mode)
	return p
}

func Partition() (*PartitionConfig, error) {
	var part C.perfstat_partition_config_t

	rc := C.perfstat_partition_config(nil, &part, C.sizeof_perfstat_partition_config_t, 1)
	if rc != 1 {
		return nil, fmt.Errorf("perfstat_partition_config() error")
	}
	p := perfstatpartitionconfig2partitionconfig(part)
	return &p, nil

}
