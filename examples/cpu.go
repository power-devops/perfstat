package main

import (
	"fmt"

	"github.com/power-devops/perfstat"
)

func main() {
	c, _ := perfstat.CpuTotal()
	fmt.Printf("CPU: \t%s\n", c.Description)
	fmt.Printf("MHz: \t%d\n", c.ProcessorHz)
	fmt.Printf("Count:\t%d\n", c.NCpusCfg)
	fmt.Printf("LA: \t%f %f %f\n", c.LoadAvg1, c.LoadAvg5, c.LoadAvg15)
	mins := float32(c.Lbolt) / 100.0 / 60.0
	fmt.Printf("Uptime:\t%f min\n", mins)
	ticks := float32(c.User + c.Sys + c.Idle + c.Wait)
	fmt.Printf("Load: %f(user) %f(sys) %f(idle) %f(wait)\n",
		float32(c.User)/ticks*100,
		float32(c.Sys)/ticks*100,
		float32(c.Idle)/ticks*100,
		float32(c.Wait)/ticks*100)
	cpus, _ := perfstat.Cpu()
	for _, c := range cpus {
		fmt.Println("CPU Name", c.Name)
		ticks := float32(c.User + c.Sys + c.Idle + c.Wait)
		fmt.Printf("Load: %f(user) %f(sys) %f(idle) %f(wait)\n",
			float32(c.User)/ticks*100,
			float32(c.Sys)/ticks*100,
			float32(c.Idle)/ticks*100,
			float32(c.Wait)/ticks*100)
	}
}
