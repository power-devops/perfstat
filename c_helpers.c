#include <libperfstat.h>

perfstat_cpu_t *get_cpu_stat(perfstat_cpu_t *cpustat, int n) {
	if (!cpustat) return NULL;

	return &(cpustat[n]);
}

double get_partition_mhz(perfstat_partition_config_t pinfo) {
	return pinfo.processorMHz;
}

char *get_ps_hostname(perfstat_pagingspace_t *ps) {
	return ps->u.nfs_paging.hostname;
}

char *get_ps_filename(perfstat_pagingspace_t *ps) {
	return ps->u.nfs_paging.filename;
}

char *get_ps_vgname(perfstat_pagingspace_t *ps) {
	return ps->u.lv_paging.vgname;
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

perfstat_fcstat_t *get_fcadapter_stat(perfstat_fcstat_t *fca, int n) {
	if (!fca) return NULL;

	return &(fca[n]);
}

perfstat_logicalvolume_t *get_logicalvolume_stat(perfstat_logicalvolume_t *lv, int n) {
	if (!lv) return NULL;

	return &(lv[n]);
}

perfstat_volumegroup_t *get_volumegroup_stat(perfstat_volumegroup_t *vg, int n) {
	if (!vg) return NULL;

	return &(vg[n]);
}

perfstat_memory_page_t *get_memorypage_stat(perfstat_memory_page_t *m, int n) {
	if (!m) return NULL;

	return &(m[n]);
}

perfstat_netbuffer_t *get_netbuffer_stat(perfstat_netbuffer_t *b, int n) {
	if (!b) return NULL;

	return &(b[n]);
}

perfstat_netinterface_t *get_netinterface_stat(perfstat_netinterface_t *b, int n) {
	if (!b) return NULL;

	return &(b[n]);
}

perfstat_netadapter_t *get_netadapter_stat(perfstat_netadapter_t *b, int n) {
	if (!b) return NULL;

	return &(b[n]);
}

perfstat_pagingspace_t *get_pagingspace_stat(perfstat_pagingspace_t *b, int n) {
	if (!b) return NULL;

	return &(b[n]);
}

perfstat_process_t *get_process_stat(perfstat_process_t *b, int n) {
	if (!b) return NULL;

	return &(b[n]);
}

perfstat_thread_t *get_thread_stat(perfstat_thread_t *b, int n) {
	if (!b) return NULL;

	return &(b[n]);
}
