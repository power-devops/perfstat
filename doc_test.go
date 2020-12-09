// +build aix
package perfstat_test

import (
	"fmt"

	"github.com/power-devops/perfstat"
)

// Example how to use the package
func Example_basic() {
	// more example you can find in the examples subdirectory
	c, err := perfstat.CpuTotalStat()
	if err != nil {
		fmt.Println("Can't get CPU statistics - exiting")
		fmt.Println(err)
		return
	}
	fmt.Printf("CPU: \t%s\n", c.Description)
	fmt.Printf("MHz: \t%d\n", c.ProcessorHz)
	fmt.Printf("Count:\t%d\n", c.NCpusCfg)
	fmt.Printf("LA: \t%f %f %f\n", c.LoadAvg1, c.LoadAvg5, c.LoadAvg15)
}
