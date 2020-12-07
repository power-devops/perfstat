package perfstat

type DiskTotal struct {
	Number    int32 /* total number of disks */
	Size      int64 /* total size of all disks (in MB) */
	Free      int64 /* free portion of all disks (in MB) */
	XRate     int64 /* __rxfers: total number of transfers from disk */
	Xfers     int64 /* total number of transfers to/from disk */
	Wblks     int64 /* 512 bytes blocks written to all disks */
	Rblks     int64 /* 512 bytes blocks read from all disks */
	Time      int64 /* amount of time disks are active */
	Version   int64 /* version number (1, 2, etc.,) */
	Rserv     int64 /* Average read or receive service time */
	MinRserv  int64 /* min read or receive service time */
	MaxRserv  int64 /* max read or receive service time */
	RTimeOut  int64 /* number of read request timeouts */
	RFailed   int64 /* number of failed read requests */
	Wserv     int64 /* Average write or send service time */
	MinWserv  int64 /* min write or send service time */
	MaxWserv  int64 /* max write or send service time */
	WTimeOut  int64 /* number of write request timeouts */
	WFailed   int64 /* number of failed write requests */
	WqDepth   int64 /* instantaneous wait queue depth (number of requests waiting to be sent to disk) */
	WqTime    int64 /* accumulated wait queueing time */
	WqMinTime int64 /* min wait queueing time */
	WqMaxTime int64 /* max wait queueing time */
}
