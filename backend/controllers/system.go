package controllers

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"time"
)

type SystemUsage struct {
	CPU      float64 `json:"cpu"`
	Cores    int     `json:"cores"`
	RAM      float64 `json:"ram"`
	Disk     float64 `json:"disk"`
	Upload   float64 `json:"upload"`
	Download float64 `json:"download"`
	TS       int64   `json:"ts"`
}

func GetSystemInfo() (*SystemUsage, error) {

	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return nil, err
	}

	numCores, err := cpu.Counts(true)
	if err != nil {
		return nil, err
	}

	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	diskStat, err := disk.Usage("/")
	if err != nil {
		return nil, err
	}

	netStat, err := net.IOCounters(true)
	if err != nil {
		return nil, err
	}
	var upload, download float64
	for _, stat := range netStat {
		upload += float64(stat.BytesSent)
		download += float64(stat.BytesRecv)
	}

	interval := time.Second
	uploadSpeed := float64(upload) / interval.Seconds()
	downloadSpeed := float64(download) / interval.Seconds()
	uploadSpeedMbps := uploadSpeed * 8 / 1e6
	downloadSpeedMbps := downloadSpeed * 8 / 1e6

	return &SystemUsage{
		CPU:      percent[0],
		Cores:    numCores,
		RAM:      vmStat.UsedPercent,
		Disk:     diskStat.UsedPercent,
		Upload:   uploadSpeedMbps,
		Download: downloadSpeedMbps,
		TS:       time.Now().Unix(),
	}, nil

}
