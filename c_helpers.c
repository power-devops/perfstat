#include "c_helpers.h"

GETFUNC(cpu)
GETFUNC(disk)
GETFUNC(diskadapter)
GETFUNC(diskpath)
GETFUNC(fcstat)
GETFUNC(logicalvolume)
GETFUNC(memory_page)
GETFUNC(netadapter)
GETFUNC(netbuffer)
GETFUNC(netinterface)
GETFUNC(pagingspace)
GETFUNC(process)
GETFUNC(thread)
GETFUNC(volumegroup)

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

time_t boottime()
{
        register struct utmpx *utmp;

	setutxent();
        while ( (utmp = getutxent()) != NULL ) {
                if (utmp->ut_type == BOOT_TIME) {
                        return utmp->ut_tv.tv_sec;
                }
        }
	endutxent();
        return -1;
}
