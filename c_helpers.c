#include <libperfstat.h>

perfstat_cpu_t *get_cpu_stat(perfstat_cpu_t *cpustat, int n) {
	if (!cpustat) return NULL;

	return &(cpustat[n]);
}

double get_partition_mhz(perfstat_partition_config_t pinfo) {
	return pinfo.processorMHz;
}

perfstat_diskadapter_t *get_diskadapter_stat(perfstat_diskadapter_t *adapter, int n) {
	if (!adapter) return NULL;

	return &(adapter[n]);
}

perfstat_disk_t *get_disk_stat(perfstat_disk_t *disk, int n) {
	if (!disk) return NULL;

	return &(disk[n]);
}

perfstat_diskpath_t *get_diskpath_stat(perfstat_diskpath_t *disk, int n) {
	if (!disk) return NULL;

	return &(disk[n]);
}
