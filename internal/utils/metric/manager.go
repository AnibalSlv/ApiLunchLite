package metric

import (
	"github.com/shirou/gopsutil/v3/process"
)

func Manager(targetPID int32) (float64, float64) {
	// Crea una instancia del proceso
	p, _ := process.NewProcess(targetPID)

	rRam := ram(p)
	rCpu := cpu(p)

	return rCpu, rRam
}
