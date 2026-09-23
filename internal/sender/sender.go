package sender

import (
	"fmt"
	"net"
	"time"

	"Network-Atlas/internal/cli"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func Run(cfg cli.Config) error {
	srcIP := net.ParseIP(cfg.SourceIP)
	dstIP := net.ParseIP(cfg.DestIP)

	if srcIP == nil || dstIP == nil {
		return fmt.Errorf("invalid IP address")
	}

	handle, err := pcap.OpenLive(
		"lo",
		65535,
		false,
		pcap.BlockForever,
	)
	if err != nil {
		return err
	}
	defer handle.Close()

	for i := 1; i <= cfg.Count; i++ {
		if err := sendPacket(handle, srcIP, dstIP, cfg.Port); err != nil {
			fmt.Printf("[%02d] FAILED: %v\n", i, err)
		} else {
			fmt.Printf(
				"[%02d] %s -> %s:%d  SUCCESS\n",
				i,
				cfg.SourceIP,
				cfg.DestIP,
				cfg.Port,
			)
		}

		time.Sleep(time.Duration(cfg.Interval) * time.Millisecond)
	}

	fmt.Printf("\nDone! %d packets processed.\n", cfg.Count)

	return nil
}

func sendPacket(
	handle *pcap.Handle,
	srcIP net.IP,
	dstIP net.IP,
	port int,
) error {

	ip := &layers.IPv4{
		Version:  4,
		TTL:      64,
		Protocol: layers.IPProtocolUDP,
		SrcIP:    srcIP,
		DstIP:    dstIP,
	}

	udp := &layers.UDP{
		SrcPort: 40000,
		DstPort: layers.UDPPort(port),
	}

	if err := udp.SetNetworkLayerForChecksum(ip); err != nil {
		return err
	}

	payload := gopacket.Payload(
		[]byte("hello from Network-Atlas"),
	)

	buffer := gopacket.NewSerializeBuffer()

	options := gopacket.SerializeOptions{
		ComputeChecksums: true,
		FixLengths:       true,
	}

	if err := gopacket.SerializeLayers(
		buffer,
		options,
		ip,
		udp,
		payload,
	); err != nil {
		return err
	}

	return handle.WritePacketData(buffer.Bytes())
}
