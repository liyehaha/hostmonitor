package disk

import (
	"io"
	"bufio"
	"strings"
	"os"
	"strconv"
	"fmt"
)

const (
	diskSectorSize uint64 = 512
)

func convertDiskSectorsToBytes(sectorCount string) (string, error) {
	sectors, err := strconv.ParseUint(sectorCount, 10, 64)
	if err != nil {
		return "", err
	}

	return strconv.FormatUint(sectors*diskSectorSize, 10), nil
}

func parseDiskStats(r io.Reader, f string) (map[string]map[int]string, error) {
	var (
		diskStats = map[string]map[int]string{}
		scanner   = bufio.NewScanner(r)
	)

	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		if len(parts) < 4 { // we strip major, minor and dev
			return nil, fmt.Errorf("invalid line in %s: %s", f, scanner.Text())
		}
		dev := parts[2]
		diskStats[dev] = map[int]string{}
		for i, v := range parts[3:] {
			diskStats[dev][i] = v
		}
		bytesRead, err := convertDiskSectorsToBytes(diskStats[dev][2])
		if err != nil {
			return nil, fmt.Errorf("invalid value for sectors read in %s: %s", f, scanner.Text())
		}
		diskStats[dev][11] = bytesRead

		bytesWritten, err := convertDiskSectorsToBytes(diskStats[dev][6])
		if err != nil {
			return nil, fmt.Errorf("invalid value for sectors written in %s: %s", f, scanner.Text())
		}
		diskStats[dev][12] = bytesWritten
	}

	return diskStats, scanner.Err()
}

func getDiskStats(f string) (map[string]map[int]string, error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return parseDiskStats(file, f)
}