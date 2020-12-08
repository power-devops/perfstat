#include <libperfstat.h>

extern perfstat_cpu_t *get_cpu_stat(perfstat_cpu_t *, int);
extern perfstat_diskadapter_t *get_diskadapter_stat(perfstat_diskadapter_t *, int);
extern perfstat_disk_t *get_disk_stat(perfstat_disk_t *, int);
extern perfstat_diskpath_t *get_diskpath_stat(perfstat_diskpath_t *, int);
extern perfstat_fcstat_t *get_fcadapter_stat(perfstat_fcstat_t *, int);
extern perfstat_logicalvolume_t *get_logicalvolume_stat(perfstat_logicalvolume_t *, int);
extern perfstat_volumegroup_t *get_volumegroup_stat(perfstat_volumegroup_t *, int);
extern double get_partition_mhz(perfstat_partition_config_t);
