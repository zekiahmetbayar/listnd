package dev

import (
	"log"
	"net"
	"testing"

	"github.com/google/gopacket/layers"
)

func TestAddrMapAdd(t *testing.T) {
	var a AddrMap
	var want, notWant, got *AddrInfo

	// add zero
	want = nil
	got = a.Add(addrZero)
	if got != want {
		t.Errorf("got = %p; want %p", got, want)
	}

	// add unspec mac
	want = nil
	got = a.Add(addrUnspecMAC)
	if got != want {
		t.Errorf("got = %p; want %p", got, want)
	}

	// add mac
	notWant = nil
	mac, err := net.ParseMAC("00:00:5e:00:53:01")
	if err != nil {
		log.Fatal(err)
	}
	got = a.Add(layers.NewMACEndpoint(mac))
	if got == notWant {
		t.Errorf("got = %p, notWant %p", got, notWant)
	}

	// add unspec ipv4
	want = nil
	got = a.Add(addrUnspecIPv4)
	if got != want {
		t.Errorf("got = %p; want %p", got, want)
	}

	// add ipv4
	notWant = nil
	ipv4 := net.ParseIP("127.0.0.1")
	got = a.Add(layers.NewIPEndpoint(ipv4))
	if got == notWant {
		t.Errorf("got = %p, notWant %p", got, notWant)
	}

	// add unspec ipv6
	want = nil
	got = a.Add(addrUnspecIPv6)
	if got != want {
		t.Errorf("got = %p; want %p", got, want)
	}

	// add ipv6
	notWant = nil
	ipv6 := net.ParseIP("::1")
	got = a.Add(layers.NewIPEndpoint(ipv6))
	if got == notWant {
		t.Errorf("got = %p, notWant %p", got, notWant)
	}
}

func TestAddrMapDel(t *testing.T) {
	var a AddrMap
	var want, got *AddrInfo

	// check invalid addresses
	a.Del(addrZero)
	a.Del(addrUnspecMAC)
	a.Del(addrUnspecIPv4)
	a.Del(addrUnspecIPv6)

	// check valid address, empty map
	ipv4 := layers.NewIPEndpoint(net.ParseIP("127.0.0.1"))
	a.Del(ipv4)

	// check valid address, filled map
	a.Add(ipv4)
	a.Del(ipv4)
	want = nil
	got = a.Get(ipv4)
	if got != want {
		t.Errorf("got = %p; want %p", got, want)
	}
}
