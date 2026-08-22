package metric

import (
	"time"

	"github.com/shirou/gopsutil/v3/process"
)

func cpu(p *process.Process) float64 {

	_, _ = p.CPUPercent()

	time.Sleep(200 * time.Millisecond)

	cpuPercent, err := p.CPUPercent()
	if err != nil {
		cpuPercent = 0.0
	}

	return cpuPercent
}
