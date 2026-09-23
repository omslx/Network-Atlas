package receiver

import (
	"fmt"

	"Network-Atlas/internal/cli"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func Run(cfg cli.Config) error {
	handle, err := pcap.OpenLive(
		"lo",
		65535,
		true,
		pcap.BlockForever,
	)
	if err != nil {
		return err
	}
	defer handle.Close()

	filter := fmt.Sprintf("udp port %d", cfg.Port)

	if err := handle.SetBPFFilter(filter); err != nil {
		return err
	}

	fmt.Printf(
		"Listening on 127.0.0.0/8:%d...\n\n",
		cfg.Port,
	)

	packetSource := gopacket.NewPacketSource(
		handle,
		handle.LinkType(),
	)

	for packet := range packetSource.Packets() {
		printPacket(packet)
	}

	return nil
}

func printPacket(packet gopacket.Packet) {
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	udpLayer := packet.Layer(layers.LayerTypeUDP)

	if ipLayer == nil || udpLayer == nil {
		return
	}

	ip := ipLayer.(*layers.IPv4)
	udp := udpLayer.(*layers.UDP)

	fmt.Printf(
		"[RECV] %s:%d -> %s:%d  SUCCESS\n",
		ip.SrcIP,
		udp.SrcPort,
		ip.DstIP,
		udp.DstPort,
	)

	if application := packet.ApplicationLayer(); application != nil {
		fmt.Printf(
			"       Payload: %s\n\n",
			string(application.Payload()),
		)
	}
}
