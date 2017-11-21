package main

import (
	"fmt"
	"flag"
	"github.com/liyehaha/hostmonitor/config"
)

const (
	namespace	          = "monitor"
)

func main() {
	promserver = config.PromPushGW.Server
	fmt.Println(promserver)
}

