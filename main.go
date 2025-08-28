package main

import (
	"fmt"
	"loxiflow/options"
	"loxiflow/pkg/dpebpf"
	"loxiflow/pkg/ipfix"
	"loxiflow/util"
	"os"
	"time"

	"github.com/jessevdk/go-flags"
)

var version string = "0.9.6-beta"

func main() {
	logFile := fmt.Sprintf("%s%s.log", "/var/log/loxiflow", os.Getenv("HOSTNAME"))
	logLevel := util.LogString2Level("debug")
	logger := util.LogItInit(logFile, logLevel, true)
	logger.LogItInfo.Println("loxiflow start")
	// Parse command-line arguments
	_, err := flags.Parse(&options.Opts)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if options.Opts.Version {
		fmt.Printf("loxiflow version: %s\n", version)
		os.Exit(0)
	}

	maps := dpebpf.InitMaps()
	// ebpf mapping and collect
	dpebpf.Start(maps, options.Opts.EbpfSetTime)

	// IPfix part
	// Map waiting time. --> TODO chan
	time.Sleep(2 * time.Second)
	ipfix.Start(maps, options.Opts.CollectorAddress, options.Opts.CollectorPort, options.Opts.IpFixSetTime, options.Opts.CollectorProtocol)
}
