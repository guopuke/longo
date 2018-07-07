package sd

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

// HealthCheck shows `OK` as the ping-pong result.
func HealthCheck(c *gin.Context) {
	message := "OK"
	c.String(http.StatusOK, "\n"+message)
}

// DiskCheck checks the disk usage.
func DiskCheck(c *gin.Context) {
	u, _ := disk.Usage("/")

	// fmt.Println(u)  ↓↓↓
	// {"path":"/","fstype":"apfs","total":250685575168,"free":152769146880,"used":94015078400,"usedPercent":37.50318634687881,"inodesTotal":9223372036854775807,"inodesUsed":1707934,"inodesFree":9223372036853067873,"inodesUsedPercent":1.8517457532618575e-11}

	totalMB := int(u.Total) / MB
	totalGB := int(u.Total) / GB
	_ = int(u.Free) / MB
	_ = int(u.Free) / GB
	usedMB := int(u.Used) / MB
	usedGB := int(u.Used) / GB
	usedPercent := int(u.UsedPercent)

	status := http.StatusOK
	text := "OK"

	if usedPercent >= 95 {
		text = "CRITICAL"
	} else if usedPercent >= 90 {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}

	message := fmt.Sprintf("%s - Free space: %dMB (%dGB) / %dMB (%dGB) | Used: %d%%", text, usedMB, usedGB, totalMB, totalGB, usedPercent)

	c.String(status, "\n"+message)
}

// CPUCheck checks the cpu usage.
func CPUCheck(c *gin.Context) {
	cores, _ := cpu.Counts(false)

	a, _ := load.Avg()
	l1 := a.Load1
	l5 := a.Load5
	l15 := a.Load15

	status := http.StatusOK
	text := "OK"

	if l5 >= float64(cores-1) {
		status = http.StatusInternalServerError
		text = "CRITICAL"
	} else if l5 >= float64(cores-2) {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}

	message := fmt.Sprintf("%s - Load average: %.2f, %.2f, %.2f | Cores: %d", text, l1, l5, l15, cores)

	c.String(status, "\n"+message)
}

// RAMCheck checks the disk usage.
func RAMCheck(c *gin.Context) {
	m, _ := mem.VirtualMemory()

	usedMB := int(m.Used) / MB
	usedGB := int(m.Used) / GB
	totalMB := int(m.Total) / MB
	totalGB := int(m.Total) / GB
	usedPercent := int(m.UsedPercent)

	status := http.StatusOK
	text := "OK"

	if usedPercent >= 95 {
		status = http.StatusInternalServerError
		text = "CRITICAL"
	} else if usedPercent >= 90 {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}

	message := fmt.Sprintf("%s - Free space:%dMB (%dGB) / %dMB (%dGB) | Used: %d%%", text, usedMB, usedGB, totalMB, totalGB, usedPercent)
	c.String(status, "\n"+message)
}
