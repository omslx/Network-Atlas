package cli

import "flag"

type Config struct {
	Send     bool
	Receive  bool
	SourceIP string
	DestIP   string
	Port     int
	Count    int
	Interval int
}

func Parse() Config {
	send := flag.Bool("c", false, "send packets")
	receive := flag.Bool("s", false, "receive packets")

	sourceIP := flag.String(
		"src",
		"127.0.0.50",
		"source IP",
	)

	destIP := flag.String(
		"dst",
		"127.0.0.1",
		"destination IP",
	)

	port := flag.Int(
		"port",
		9999,
		"UDP port",
	)

	count := flag.Int(
		"count",
		10,
		"number of packets",
	)

	interval := flag.Int(
		"interval",
		100,
		"interval between packets in milliseconds",
	)

	flag.Parse()

	return Config{
		Send:     *send,
		Receive:  *receive,
		SourceIP: *sourceIP,
		DestIP:   *destIP,
		Port:     *port,
		Count:    *count,
		Interval: *interval,
	}
}
