// File: internal/sysinfo/sysinfo.go
package sysinfo

import (
    "fmt"
    "strings"

    "github.com/shirou/gopsutil/v3/cpu"
    "github.com/shirou/gopsutil/v3/host"
    "github.com/shirou/gopsutil/v3/mem"
    "github.com/shirou/gopsutil/v3/net"
    "github.com/shirou/gopsutil/v3/process"
)

// Data holds collected metrics.
type Data struct {
    Timestamp  string
    IPs        []string
    Hostname   string
    OS         string
    CPUPercent []float64
    MemTotal   uint64
    MemUsed    uint64
    Users      []string
}

// Collect gathers requested modules.
func Collect(mods []string) *Data {
    d := &Data{Timestamp: time.Now().Format(time.RFC3339)}
    for _, m := range mods {
        switch strings.ToLower(m) {
        case "ip":
            ifs, _ := net.Interfaces()
            for _, iface := range ifs {
                for _, addr := range iface.Addrs {
                    d.IPs = append(d.IPs, addr.Addr)
                }
            }
        case "os":
            hi, _ := host.Info()
            d.Hostname = hi.Hostname
            d.OS = fmt.Sprintf("%s %s", hi.Platform, hi.PlatformVersion)
        case "cpu":
            pct, _ := cpu.Percent(0, false)
            d.CPUPercent = pct
        case "mem":
            mi, _ := mem.VirtualMemory()
            d.MemTotal = mi.Total
            d.MemUsed = mi.Used
        case "users":
            seen := map[string]struct{}{}
            procs, _ := process.Processes()
            for _, p := range procs {
                if u, err := p.Username(); err == nil {
                    if _, ok := seen[u]; !ok {
                        seen[u] = struct{}{}
                        d.Users = append(d.Users, u)
                    }
                }
            }
        }
    }
    return d
}
