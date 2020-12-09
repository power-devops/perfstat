package main

import (
	"fmt"
	"time"

	"github.com/power-devops/perfstat"
)

func main() {
	c, _ := perfstat.CpuTotalStat()
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
	cpus, _ := perfstat.CpuStat()
	for _, c := range cpus {
		fmt.Println("CPU Name", c.Name)
		ticks := float32(c.User + c.Sys + c.Idle + c.Wait)
		fmt.Printf("Load: %f(user) %f(sys) %f(idle) %f(wait)\n",
			float32(c.User)/ticks*100,
			float32(c.Sys)/ticks*100,
			float32(c.Idle)/ticks*100,
			float32(c.Wait)/ticks*100)
	}
	util, _ := perfstat.CpuUtilStat(5 * time.Second)
	fmt.Printf("=======Overall CPU Utilization Metrics=======\n")
	fmt.Printf("User Percentage =                %v\n", util.UserPct)
	fmt.Printf("System Percentage =              %v\n", util.KernPct)
	fmt.Printf("Idle Percentage =                %v\n", util.IdlePct)
	fmt.Printf("Wait Percentage =                %v\n", util.WaitPct)
	fmt.Printf("Physical Busy =                  %v\n", util.PhysicalBusy)
	fmt.Printf("Physical Consumed =              %v\n", util.PhysicalConsumed)
	fmt.Printf("Freq Percentage =                %v\n", util.FreqPct)
	fmt.Printf("Entitlement Used Percentage =    %v\n", util.EntitlementPct)
	fmt.Printf("Entitlement Busy Percentage =    %v\n", util.BusyPct)
	fmt.Printf("Idle Cycles Donated Percentage = %v\n", util.IdleDonatedPct)
	fmt.Printf("Busy Cycles Donated Percentage = %v\n", util.BusyDonatedPct)
	fmt.Printf("Idle Cycles Stolen Percentage =  %v\n", util.IdleStolenPct)
	fmt.Printf("Busy Cycles Stolen Percentage =  %v\n", util.BusyStolenPct)
	fmt.Printf("User percentage for logical cpu in ticks = %v\n", util.LUserPct)
	fmt.Printf("Sytem percentage for logical cpu in ticks= %v\n", util.LKernPct)
	fmt.Printf("Idle percentage for logical cpu in ticks=  %v\n", util.LIdlePct)
	fmt.Printf("Wait percentage for logical cpu in ticks=  %v\n", util.LWaitPct)
	fmt.Printf("delta time in milliseconds =  %v \n", util.DeltaTime)
	fmt.Printf("=============================================\n")
}
