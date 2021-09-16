package scanner

import (
	"fmt"
	"net"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

var ModeScanFuncMap map[string]ConnectFunc

func init() {
	ModeScanFuncMap = make(map[string]ConnectFunc)
	ModeScanFuncMap["syn"] = SynConnect
	ModeScanFuncMap["full"] = FullConnect
}

//ConnectFunc TCP连接函数
type ConnectFunc func(dstIP string, dstPort int) bool

//FullConnect TCP全连接
func FullConnect(dstIP string, dstPort int) bool {
	con, err := net.DialTimeout("tcp4", fmt.Sprintf(`%s:%d`, dstIP, dstPort), 1*time.Second)
	if err != nil {
		return false
	}
	con.Close()
	return true
}

//SynConnect TCP半连接
func SynConnect(host string, dstPort int) bool {
	srcIP, srcPort, err := localIPPort(net.ParseIP(host))
	dstAddrs, err := net.LookupIP(host)
	if err != nil {
		return false
	}
	dstIP := dstAddrs[0].To4()
	var dstport layers.TCPPort
	dstport = layers.TCPPort(dstPort)
	srcport := layers.TCPPort(srcPort)

	// Our IP header... not used, but necessary for TCP checksumming.
	ip := &layers.IPv4{
		SrcIP:    srcIP,
		DstIP:    dstIP,
		Protocol: layers.IPProtocolTCP,
	}
	// Our TCP header
	tcp := &layers.TCP{
		SrcPort: srcport,
		DstPort: dstport,
		SYN:     true,
	}
	err = tcp.SetNetworkLayerForChecksum(ip)

	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{
		ComputeChecksums: true,
		FixLengths:       true,
	}

	if err := gopacket.SerializeLayers(buf, opts, tcp); err != nil {
		return false
	}
	conn, err := net.ListenPacket("ip4:tcp", "0.0.0.0")
	if err != nil {
		return false
	}
	defer conn.Close()

	if _, err := conn.WriteTo(buf.Bytes(), &net.IPAddr{IP: dstIP}); err != nil {
		return false
	}

	// Set deadline so we don't wait forever.
	if err := conn.SetDeadline(time.Now().Add(time.Second)); err != nil {
		return false
	}

	for {
		b := make([]byte, 4096)
		n, addr, err := conn.ReadFrom(b)
		if err != nil {
			return false
		} else if addr.String() == dstIP.String() {
			// Decode a packet
			packet := gopacket.NewPacket(b[:n], layers.LayerTypeTCP, gopacket.Default)
			// Get the TCP layer from this packet
			if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
				tcp, _ := tcpLayer.(*layers.TCP)

				if tcp.DstPort == srcport {
					if tcp.SYN && tcp.ACK {
						// log.Printf("%v:%d is OPEN\n", dstIP, dstport)
						return true
					} else {
						return false
					}
				}
			}
		}
	}
}

// get the local ip and port based on our destination ip
func localIPPort(dstIP net.IP) (net.IP, int, error) {
	serverAddr, err := net.ResolveUDPAddr("udp", dstIP.String()+":54321")
	if err != nil {
		return nil, 0, err
	}
	// We don't actually connect to anything, but we can determine
	// based on our destination ip what source ip we should use.
	if con, err := net.DialUDP("udp", nil, serverAddr); err == nil {
		if udpaddr, ok := con.LocalAddr().(*net.UDPAddr); ok {
			return udpaddr.IP, udpaddr.Port, nil
		}
	}
	return nil, -1, err
}
