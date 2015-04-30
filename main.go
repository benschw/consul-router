package main

import (
	"net/http"
	"os"

	"github.com/AaronO/gogo-proxy"
	"github.com/benschw/dns-clb-go/clb"
)

func getConfig() (string, string, string, string) {
	nsIp := os.Getenv("NS_IP")     // use /etc/resolv.conf if not set
	nsPort := os.Getenv("NS_PORT") // 53
	target := os.Getenv("TARGET")  // service.consul
	domain := os.Getenv("DOMAIN")  // edge

	if nsPort == "" {
		nsPort = "53"
	}

	return nsIp, nsPort, target, domain
}

func main() {
	nsIp, nsPort, target, domain := getConfig

	var lb AddressGetter
	if nsIp == "" {
		lb = clb.NewDefaultClb(clb.RoundRobin)
	} else {
		lb = clb.NewClb(nsIp, nsPort, clb.RoundRobin)
	}

	requestMapper := &SrvRecordRequestMapper{
		Lb:     lb,
		Target: target,
		Domain: domain,
	}

	p, _ := proxy.New(proxy.ProxyOptions{
		Balancer: requestMapper.MapRequest,
	})

	http.ListenAndServe(":8080", p)
}
