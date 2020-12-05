#include <libperfstat.h>

perfstat_cpu_t *get_cpu_stat(perfstat_cpu_t *cpustat, int n) {
	perfstat_cpu_t pcpu;

	if (!cpustat) return NULL;

	return &(cpustat[n]);
}

double get_partition_mhz(perfstat_partition_config_t pinfo) {
	return pinfo.processorMHz;
}
