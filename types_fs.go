package perfstat

type FileSystem struct {
	Device      string /* name of the mounted device */
	MountPoint  string /* where the device is mounted */
	TotalBlocks int64  /* number of 512 bytes blocks in the filesystem */
	FreeBlocks  int64  /* number of free 512 bytes block in the filesystem */
	TotalInodes int64  /* total number of inodes in the filesystem */
	FreeInodes  int64  /* number of free inodes in the filesystem */
}
