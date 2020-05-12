package cmd

import (
	"fmt"
	"io"

	"github.com/google/gopacket/layers"
)

// prefixList stores router prefixes
type prefixList struct {
	prefixes []*PrefixInfo
}

// clear deletes all prefixes
func (p *prefixList) clear() {
	p.prefixes = nil
}

// add adds a prefix
func (p *prefixList) add(prefix layers.ICMPv6Option) *PrefixInfo {
	pf := PrefixInfo{}
	pf.Prefix = prefix
	p.prefixes = append(p.prefixes, &pf)
	return &pf
}

// get returns all prefixes
func (p *prefixList) get() []*PrefixInfo {
	return p.prefixes
}

// Print prints all prefixes
func (p *prefixList) Print(w io.Writer) {
	for _, prefix := range p.prefixes {
		fmt.Fprintf(w, "      %s\n", prefix)
	}
}
