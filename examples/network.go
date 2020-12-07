package main

import (
	"fmt"

	"github.com/power-devops/perfstat"
)

func main() {
	ninfo, _ := perfstat.NetIfaceTotalStat()

	fmt.Printf("Total network statistics\n")
	fmt.Printf("========================\n")
	fmt.Printf("number of network interfaces            : %v\n", ninfo.Number)
	fmt.Printf("number of packets received on interface : %v\n", ninfo.IPackets)
	fmt.Printf("number of bytes received on interface   : %v\n", ninfo.IBytes)
	fmt.Printf("number of input errors on interface     : %v\n", ninfo.IErrors)
	fmt.Printf("number of packets sent on interface     : %v\n", ninfo.OPackets)
	fmt.Printf("number of bytes sent on interface       : %v\n", ninfo.OBytes)
	fmt.Printf("number of output errors on interface    : %v\n", ninfo.OErrors)
	fmt.Printf("number of collisions on csma interface  : %v\n", ninfo.Collisions)
	fmt.Printf("number of packets not transmitted       : %v\n", ninfo.XmitDrops)

}
