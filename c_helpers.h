#include <libperfstat.h>

extern perfstat_cpu_t *get_cpu_stat(perfstat_cpu_t *, int);
extern perfstat_diskadapter_t *get_diskadapter_stat(perfstat_diskadapter_t *, int);
extern perfstat_disk_t *get_disk_stat(perfstat_disk_t *, int);
extern perfstat_diskpath_t *get_diskpath_stat(perfstat_diskpath_t *, int);
extern perfstat_fcstat_t *get_fcadapter_stat(perfstat_fcstat_t *, int);
extern perfstat_logicalvolume_t *get_logicalvolume_stat(perfstat_logicalvolume_t *, int);
extern perfstat_volumegroup_t *get_volumegroup_stat(perfstat_volumegroup_t *, int);
extern perfstat_memory_page_t *get_memorypage_stat(perfstat_memory_page_t *, int);
extern perfstat_netbuffer_t *get_netbuffer_stat(perfstat_netbuffer_t *, int);
extern perfstat_netinterface_t *get_netinterface_stat(perfstat_netinterface_t *, int);
extern perfstat_netadapter_t *get_netadapter_stat(perfstat_netadapter_t *, int);
extern perfstat_pagingspace_t *get_pagingspace_stat(perfstat_pagingspace_t *, int);
extern perfstat_process_t *get_process_stat(perfstat_process_t *, int);
extern perfstat_thread_t *get_thread_stat(perfstat_thread_t *, int);
extern double get_partition_mhz(perfstat_partition_config_t);
extern char *get_ps_hostname(perfstat_pagingspace_t *);
extern char *get_ps_filename(perfstat_pagingspace_t *);
extern char *get_ps_vgname(perfstat_pagingspace_t *);
