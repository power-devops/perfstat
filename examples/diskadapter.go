package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/power-devops/perfstat"
)

func main() {
	interval := flag.Int("interval", 1, "interval time between each report")
	count := flag.Int("count", 1, "number of reports")
	flag.Parse()

	da_prev, _ := perfstat.DiskAdapterStat()

	fmt.Printf("\n%-8s %7s %8s %8s %8s %8s\n", " Name ", " Disks ", " Size ", " Free ", " ARS ", " AWS ")
	fmt.Printf("%-8s %7s %8s %8s %8s %8s\n", "======", "======", "======", "======", "=====", "=====")
	for *count > 0 {
		time.Sleep(time.Duration(*interval) * time.Second)

		da, _ := perfstat.DiskAdapterStat()
		for i := range da {
			delta_write := da[i].Wblks - da_prev[i].Wblks
			delta_read := da[i].Rblks - da_prev[i].Rblks
			delta_xfers := da[i].Xfers - da_prev[i].Xfers
			delta_xrate := da[i].XRate - da_prev[i].XRate
			avg_rps := int64(0)
			if da[i].XRate-da_prev[i].XRate != 0 {
				avg_rps = delta_read / (da[i].XRate - da_prev[i].XRate)
			}
			avg_wps := int64(0)
			if delta_xfers-delta_xrate != 0 {
				avg_wps = delta_write / (delta_xfers - delta_xrate)
			}
			fmt.Printf("%-8s %7d %8v %8v %8v %8v\n",
				da[i].Name,
				da[i].Number,
				da[i].Size,
				da[i].Free,
				avg_rps,
				avg_wps)
		}
		da_prev = da
		*count--
		fmt.Println()
	}
}
