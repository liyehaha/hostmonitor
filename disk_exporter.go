package main

import (
	"fmt"
	"flag"

	//"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/node_exporter/collector"

	config "github.com/liyehaha/hostmonitor/config"
)

const (
	namespace	          = "monitor"
)

func diskmonitor() {

	disk, _ := collector.NewDiskstatsCollector()

	fmt.Println(disk)
	fmt.Println(1)
	//return err
}

func main() {

	var config_path = flag.String("conf", "config.yaml", "config file")
	config.Load(*config_path)
	fmt.Println(2)
	promserver := config.Monitor.PromPushGW
	fmt.Println(promserver)
	
	diskmonitor()
}

