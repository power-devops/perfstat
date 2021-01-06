package main

import (
	"fmt"

	"github.com/power-devops/perfstat"
)

func main() {
	fss, _ := perfstat.FileSystemStat()

	for _, fs := range fss {
		fmt.Printf("===== Filesystem %s =====\n", fs.Device)
		fmt.Printf("Mounted on       %s\n", fs.MountPoint)
		fmt.Printf("Filesystem type  %s\n", fs.TypeString())
		fmt.Printf("Filesystem flags %v\n", fs.Flags)
		fmt.Printf("Filesystem flags %s\n", fs.FlagsString())
		fmt.Printf("Total 512 blocks %v\n", fs.TotalBlocks)
		fmt.Printf("Total Megabytes  %v\n", fs.TotalBlocks*512/1048576)
		fmt.Printf("Free 512 blocks  %v\n", fs.FreeBlocks)
		fmt.Printf("Free Megabytes   %v\n", fs.FreeBlocks*512/1048576)
		fmt.Printf("Total I-Nodes    %v\n", fs.TotalInodes)
		fmt.Printf("Free I-Nodes     %v\n", fs.FreeInodes)
	}
}
