package main

import (
	//"./config"
	disk "./disk"

	"fmt"
)

func main() {
	var f = "/proc/diskstats"

	fmt.Println(disk.GetDiskStats(f))
}