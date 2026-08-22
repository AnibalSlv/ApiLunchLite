package metric

import "github.com/shirou/gopsutil/v3/process"

func ram(p *process.Process) float64 {
	memInfo, err := p.MemoryInfo()
	var memUsageMB float64
	if err == nil {
		// RSS (Resident Set Size) es la memoria física real usada
		memUsageMB = float64(memInfo.RSS) / (1024 * 1024)
	}

	return memUsageMB
}
