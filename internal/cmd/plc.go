package cmd

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

// parsePlc parses plc (power-line communication/homeplug) packets
func parsePlc(packet gopacket.Packet) {
	ethLayer := packet.Layer(layers.LayerTypeEthernet)
	if ethLayer != nil {
		eth, _ := ethLayer.(*layers.Ethernet)
		if eth.EthernetType == 0x88e1 || eth.EthernetType == 0x8912 {
			debug("PLC packet")
			linkSrc, _ := getMacs(packet)

			// add device and mark this device as a powerline
			devices.add(linkSrc)
			devices[linkSrc].powerline.enable()
			devices[linkSrc].powerline.setTimestamp(
				packet.Metadata().Timestamp)
		}
	}
}