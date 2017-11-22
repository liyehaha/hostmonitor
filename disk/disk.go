package disk

import (
	"os/exec"
	"strings"
	"encoding/json"
)

type DiskUsage struct {
	Name	string
	Total	string
	Used    string
	Percent string
}

func getDiskStats() ([]byte, error) {
	cmd := "df -h --output=source,size,used,pcent | grep '^/dev' | column -t -o '|' | sed 's/ //g'"
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return nil, err
	}
	return out, err
}

func formatStrResult(b []byte) map[string]*DiskUsage {
	map1 := strings.Split(string(b), "\n")
	map2 := make(map[string]*DiskUsage)
	for _, v := range map1 {
		if v != "" {
			s := strings.Split(v, "|")
			map2[s[0]] = &DiskUsage{Total: s[1], Used: s[2], Percent: s[3]}
		}
	}
	return map2
}

func DiskMonitor() map[string]*DiskUsage {
	disk_usage, _ := getDiskStats()
	result := formatStrResult(disk_usage)
	return result
}