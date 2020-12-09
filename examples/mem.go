package main

import (
	"fmt"

	"github.com/power-devops/perfstat"
)

func main() {
	minfo, _ := perfstat.MemoryTotalStat()
	fmt.Printf("Memory statistics\n")
	fmt.Printf("-----------------\n")
	fmt.Printf("real memory size                 : %v MB\n", minfo.RealTotal*4096/1024/1024)
	fmt.Printf("reserved paging space            : %v MB\n", minfo.PgSpRsvd)
	fmt.Printf("virtual memory size              : %v MB\n", minfo.VirtualTotal*4096/1024/1024)
	fmt.Printf("number of free pages             : %v\n", minfo.RealFree)
	fmt.Printf("number of pinned pages           : %v\n", minfo.RealPinned)
	fmt.Printf("number of pages in file cache    : %v\n", minfo.NumPerm)
	fmt.Printf("total paging space pages         : %v\n", minfo.PgSpTotal)
	fmt.Printf("free paging space pages          : %v\n", minfo.PgSpFree)
	fmt.Printf("used paging space                : %v%%\n", float32(minfo.PgSpTotal-minfo.PgSpFree)*100.0/float32(minfo.PgSpTotal))
	fmt.Printf("number of paging space page ins  : %v\n", minfo.PgSpIn)
	fmt.Printf("number of paging space page outs : %v\n", minfo.PgSpOut)
	fmt.Printf("number of page ins               : %v\n", minfo.PageIn)
	fmt.Printf("number of page outs              : %v\n", minfo.PageOut)
	fmt.Println()
	psinfo, _ := perfstat.PagingSpaceStat()
	for i := range psinfo {
		fmt.Printf("Statistics for paging space : %s\n", psinfo[i].Name)
		fmt.Printf("------------------------------\n")
		fmt.Printf("Type                       : %v\n", psinfo[i].Type)
		fmt.Printf("Volume group               : %v\n", psinfo[i].VGName)
		fmt.Printf("Size in logical partitions : %v\n", psinfo[i].LPSize)
		fmt.Printf("Size in MB                 : %v\n", psinfo[i].MBSize)
		fmt.Printf("Used in MB                 : %v\n", psinfo[i].MBUsed)
	}
}
