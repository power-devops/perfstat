package main

import (
	"fmt"

	"github.com/power-devops/perfstat"
)

func main() {
	perfstat.EnableLVMStat()
	vgs, _ := perfstat.VolumeGroupStat()
	lvs, _ := perfstat.LogicalVolumeStat()

	for i := range vgs {
		fmt.Printf("Volume Group : %v\n", vgs[i].Name)
		fmt.Printf("==============\n")
		fmt.Printf("Number of physical volumes in the volume group        : %v\n", vgs[i].TotalDisks)
		fmt.Printf("Number of active physical volumes in the volume group : %v\n", vgs[i].ActiveDisks)
		fmt.Printf("Number of logical volumes in the volume group         : %v\n", vgs[i].TotalLogicalVolumes)
		fmt.Printf("Number of logical volumes opened in the volume group  : %v\n", vgs[i].OpenedLogicalVolumes)
		fmt.Printf("Number of read and write requests                     : %v\n", vgs[i].IOCnt)
		fmt.Printf("Number of Kilobytes read                              : %v\n", vgs[i].KBReads)
		fmt.Printf("Number of Kilobytes written                           : %v\n", vgs[i].KBWrites)
		fmt.Println()
	}

	for i := range lvs {
		fmt.Printf("Logical Volume: %s\n", lvs[i].Name)
		fmt.Printf("================\n")
		fmt.Printf("Volume group name                                                    : %v\n", lvs[i].VGName)
		fmt.Printf("Physical partition size                                              : %v MB\n", lvs[i].PPsize)
		fmt.Printf("Total number of logical paritions configured for this logical volume : %v\n", lvs[i].LogicalPartitions)
		fmt.Printf("Number of physical mirrors for each logical partition                : %v\n", lvs[i].Mirrors)
		fmt.Printf("Number of read and write requests                                    : %v\n", lvs[i].IOCnt)
		fmt.Printf("Number of Kilobytes read                                             : %v\n", lvs[i].KBReads)
		fmt.Printf("Number of Kilobytes written                                          : %v\n", lvs[i].KBWrites)
		fmt.Println()
	}
	perfstat.DisableLVMStat()
}
