package main

import (
        "fmt"

        "github.com/power-devops/perfstat"
)

func main() {
	dinfo, _ := perfstat.DisksStat()
	fmt.Printf("Total disk statistics\n");
	fmt.Printf("---------------------\n");
	fmt.Printf("number of  disks         : %v\n", dinfo.Number);
	fmt.Printf("total disk space         : %v\n", dinfo.Size);
	fmt.Printf("total free space         : %v\n", dinfo.Free);
	fmt.Printf("number of transfers      : %v\n", dinfo.Xfers);
	fmt.Printf("number of blocks written : %v\n", dinfo.Wblks);
	fmt.Printf("number of blocks read    : %v\n", dinfo.Rblks);
}
