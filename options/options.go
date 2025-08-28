package options

var Opts struct {
	Version           bool   `short:"v" long:"version" description:"Show loxiflow version"`
	IpFixSetTime      int    `short:"t" long:"ipfix-set-time" description:"Set time for the ip-fix send time" default:"1"`
	EbpfSetTime       int    `short:"e" long:"ebpf-set-time" description:"Set time for the ebpf collecting time" default:"1"`
	CollectorAddress  string `short:"a" long:"collector-address" description:"Set the collector address" default:"127.0.0.1"`
	CollectorPort     int    `short:"p" long:"collector-port" description:"Set the collector port" default:"4739"`
	CollectorProtocol string `long:"collector-protocol" description:"Set the collector protocol" default:"udp"`
}
