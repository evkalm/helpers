package helpers

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

type PerformanceData struct {
	Time  time.Time
	Heap  uint64
	Stack uint64
}

func GetStartMetrics() PerformanceData {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	return PerformanceData{
		Time:  time.Now(),
		Heap:  memStats.HeapAlloc,
		Stack: memStats.StackInuse,
	}
}

func PrintPerformance(startMetrics PerformanceData) {
	endMetrics := GetStartMetrics()

	executionTime := endMetrics.Time.Sub(startMetrics.Time)
	heapUsage := endMetrics.Heap - startMetrics.Heap
	stackUsage := endMetrics.Stack - startMetrics.Stack

	var mess strings.Builder

	mess.WriteString(colorPerf("*** Performance Data ***\n"))
	mess.WriteString(fmt.Sprintf(colorMapKey("  Execution Time: %s\n"), colorValue(executionTime)))
	mess.WriteString(fmt.Sprintf(colorMapKey("  Heap Usage: %s\n"), colorValue(humanReadableSize(heapUsage))))
	mess.WriteString(fmt.Sprintf(colorMapKey("  Stack Usage: %s\n"), colorValue(humanReadableSize(stackUsage))))

	fmt.Println(mess.String())
}

func humanReadableSize(size uint64) string {
	sizes := []string{"B", "KB", "MB", "GB", "TB"}
	var i int
	s := float64(size)
	for i = 0; s >= 1024 && i < len(sizes)-1; i++ {
		s /= 1024
	}
	return fmt.Sprintf("%.3f %s", s, sizes[i])
}
