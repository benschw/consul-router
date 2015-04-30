package main

import (
	"fmt"
	"testing"

	"github.com/benschw/dns-clb-go/dns"
)

type StubAddressGetter struct {
	Val dns.Address
}

func (lb *StubAddressGetter) GetAddress(address string) (dns.Address, error) {
	if address == "foo.service.consul" {
		return lb.Val, nil
	}
	return dns.Address{}, fmt.Errorf("%s not found", address)
}

func TestSvcKeyParsing(t *testing.T) {
	// given
	requestHost := "foobarbaz.edge"

	address := dns.Address{Address: "10.0.0.1", Port: 5678}

	lb := &StubAddressGetter{Val: address}

	mapper := &SrvRecordRequestMapper{
		Lb:     lb,
		Target: "service.consul",
		Domain: "edge",
	}

	// when
	key, err := mapper.getServiceKey(requestHost)

	// then
	if err != nil {
		t.Error(err)
	}
	if key != "foobarbaz" {
		t.Errorf("expected foobarbaz, got '%s'", key)
	}
}
