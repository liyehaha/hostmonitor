package main

import (
	//"./config"
	"./disk"

	"fmt"
)

func main() {
	d := disk.DiskMonitor()

	fmt.Println(d)
}