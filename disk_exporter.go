package main

import (
	"fmt"
	"flag"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/node_exporter/collector"

	config "github.com/liyehaha/hostmonitor/config"
)

const (
	namespace	          = "monitor"
)

func diskmonitor() {

	disk, _ := collector.NewDiskstatsCollector()

	sink := make(chan prometheus.Metric)
	disk.Update(sink)
}

func main() {

	var config_path = flag.String("conf", "config.yaml", "config file")
	config.Load(*config_path)

	promserver := config.Monitor.PromPushGW
	fmt.Println(promserver)
	
	
}

