package perfstat

type NetIfaceTotal struct {
	Number     int32 /* number of network interfaces  */
	IPackets   int64 /* number of packets received on interface */
	IBytes     int64 /* number of bytes received on interface */
	IErrors    int64 /* number of input errors on interface */
	OPackets   int64 /* number of packets sent on interface */
	OBytes     int64 /* number of bytes sent on interface */
	OErrors    int64 /* number of output errors on interface */
	Collisions int64 /* number of collisions on csma interface */
	XmitDrops  int64 /* number of packets not transmitted */
	Version    int64 /* version number (1, 2, etc.,) */
}
