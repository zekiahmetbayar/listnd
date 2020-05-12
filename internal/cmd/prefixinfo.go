package cmd

import (
	"fmt"
	"net"

	"github.com/google/gopacket/layers"
)

// PrefixInfo stores a router's prefix information
type PrefixInfo struct {
	timeInfo
	Prefix layers.ICMPv6Option
}

// String converts the prefix to a string
func (p *PrefixInfo) String() string {
	prefixFmt := "Prefix: %-34s (age: %.f)"
	pfLen := uint8(p.Prefix.Data[0])
	pf := net.IP(p.Prefix.Data[14:])
	ps := fmt.Sprintf("%v/%v", pf, pfLen)
	return fmt.Sprintf(prefixFmt, ps, p.getAge())
}
