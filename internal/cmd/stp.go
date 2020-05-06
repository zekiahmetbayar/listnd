package cmd

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

// parseStp parses stp packets
func parseStp(packet gopacket.Packet) {
	stpLayer := packet.Layer(layers.LayerTypeSTP)
	if stpLayer != nil {
		debug("STP packet")
		linkSrc, _ := getMacs(packet)

		// add device and mark this device as a bridge
		devices.add(linkSrc)
		devices[linkSrc].bridge.enable()
		devices[linkSrc].bridge.setTimestamp(
			packet.Metadata().Timestamp)
	}
}