package main

import (
	"fmt"

	"github.com/power-devops/perfstat"
)

func main() {
	dinfo, _ := perfstat.DiskTotalStat()
	fmt.Printf("Total disk statistics\n")
	fmt.Printf("---------------------\n")
	fmt.Printf("number of  disks         : %v\n", dinfo.Number)
	fmt.Printf("total disk space         : %v\n", dinfo.Size)
	fmt.Printf("total free space         : %v\n", dinfo.Free)
	fmt.Printf("number of transfers      : %v\n", dinfo.Xfers)
	fmt.Printf("number of blocks written : %v\n", dinfo.Wblks)
	fmt.Printf("number of blocks read    : %v\n", dinfo.Rblks)

	fmt.Println()
	disks, _ := perfstat.DiskStat()
	for i := range disks {
		fmt.Printf("Statistics for disk : %s\n", disks[i].Name)
		fmt.Printf("---------------------\n")
		fmt.Printf("description             : %v\n", disks[i].Description)
		fmt.Printf("volume group name       : %v\n", disks[i].VGName)
		fmt.Printf("adapter name            : %v\n", disks[i].Adapter)
		fmt.Printf("size                    : %v MB\n", disks[i].Size)
		fmt.Printf("free space              : %v MB\n", disks[i].Free)
		fmt.Printf("number of blocks read   : %v blocks of %v bytes\n", disks[i].Rblks, disks[i].BSize)
		fmt.Printf("number of block written : %v blocks of %v bytes\n", disks[i].Wblks, disks[i].BSize)
		fmt.Println()
	}

	fmt.Println()
	paths, _ := perfstat.DiskPathStat()
	for i := range paths {
		fmt.Printf("Statistics for path : %s\n", paths[i].Name)
		fmt.Printf("---------------------\n")
		fmt.Printf("number of blocks read    : %v\n", paths[i].Rblks)
		fmt.Printf("number of blocks written : %v\n", paths[i].Wblks)
		fmt.Printf("adapter name             : %v\n", paths[i].Adapter)
		fmt.Println()
	}
}
