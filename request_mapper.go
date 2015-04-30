package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/benschw/dns-clb-go/dns"
)

type AddressGetter interface {
	GetAddress(string) (dns.Address, error)
}

type SrvRecordRequestMapper struct {
	Lb     AddressGetter
	Target string
	Domain string
}

func (m *SrvRecordRequestMapper) getServiceKey(fullHost string) (string, error) {
	parts := strings.Split(fullHost, ":")
	host := parts[0]

	// "Domain" should be a substring of "host"
	if len(host) <= len(m.Domain) {
		return "", fmt.Errorf("Request for '%s' doesn't match Domain '%s'", host, m.Domain)
	}

	// tail end of "host" should match "Domain"
	if host[len(host)-len(m.Domain):] != m.Domain {
		return "", fmt.Errorf("Request for '%s' doesn't match Domain '%s'", host, m.Domain)
	}

	// return "host" up to "Domain"
	return host[0 : len(host)-len(m.Domain)-1], nil
}

func (m *SrvRecordRequestMapper) getAddress(host string) (dns.Address, error) {
	svcKey, err := m.getServiceKey(host)
	if err != nil {
		return dns.Address{}, err
	}
	address := fmt.Sprintf("%s.%s", svcKey, m.Target)
	log.Printf("Lookup: %s", address)
	return m.Lb.GetAddress(address)
}

func (m *SrvRecordRequestMapper) MapRequest(req *http.Request) (string, error) {
	address, err := m.getAddress(req.Host)
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("Not Found: %s %s %s", req.Method, req.RemoteAddr, req.URL)
	}

	log.Printf("%s %s %s --> http://%s:%d", req.Method, req.RemoteAddr, req.URL, address.Address, address.Port)
	return fmt.Sprintf("http://%s:%d", address.Address, address.Port), nil
}
