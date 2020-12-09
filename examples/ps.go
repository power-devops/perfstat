package main

import (
	"fmt"

	"github.com/power-devops/perfstat"
)

func main() {
	psinfo, _ := perfstat.ProcessStat()
	for i := range psinfo {
		fmt.Printf("Process name                     : %v\n", psinfo[i].ProcessName)
		fmt.Printf("Process ID                       : %v\n", psinfo[i].PID)
		fmt.Printf("Process priority                 : %v\n", psinfo[i].Priority)
		fmt.Printf("Thread count                     : %v\n", psinfo[i].NumThreads)
		fmt.Printf("Owner UID                        : %v\n", psinfo[i].UID)
		fmt.Printf("WLM class ID                     : %v\n", psinfo[i].ClassID)
		fmt.Printf("Process virtual size             : %v\n", psinfo[i].Size)
		fmt.Printf("Real memory used for data        : %v\n", psinfo[i].RealMemData)
		fmt.Printf("Real memory used for text        : %v\n", psinfo[i].RealMemText)
		fmt.Printf("Virtual memory used for data     : %v\n", psinfo[i].VirtMemData)
		fmt.Printf("Virtual memory used for text     : %v\n", psinfo[i].VirtMemText)
		fmt.Printf("Data size from shared library    : %v\n", psinfo[i].SharedLibDataSize)
		fmt.Printf("Heap size                        : %v\n", psinfo[i].HeapSize)
		fmt.Printf("Real memory in use by process    : %v\n", psinfo[i].RealInUse)
		fmt.Printf("Virtual memory in use by process : %v\n", psinfo[i].VirtInUse)
		fmt.Printf("Pinned memory for this process   : %v\n", psinfo[i].Pinned)
		fmt.Printf("Paging space in use              : %v\n", psinfo[i].PgSpInUse)
		fmt.Printf("File pages used                  : %v\n", psinfo[i].FilePages)
		fmt.Printf("User mode CPU time               : %v\n", psinfo[i].UCpuTime)
		fmt.Printf("System mode CPU time             : %v\n", psinfo[i].SCpuTime)
		fmt.Printf("Bytes written to disk            : %v\n", psinfo[i].InBytes)
		fmt.Printf("Bytes read from disk             : %v\n", psinfo[i].OutBytes)
		fmt.Println()
	}
}
