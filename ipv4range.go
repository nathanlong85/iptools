package ipv4range

import "net"

// IPv4Range represents an IPv4 subnet
type IPv4Range struct {
	IPs       []net.IP
	Mask      net.IPMask
	Network   net.IP
	Broadcast net.IP
}

// New IPv4 network range. Results include network and broadcase IPs.
// cidrNet format should be "192.168.0.0/23"
func New(cidrNet string) (IPv4Range, error) {
	var r IPv4Range

	ip, ipnet, err := net.ParseCIDR(cidrNet)
	if err != nil {
		return r, err
	}

	var ips []net.IP
	for ip.Mask(ipnet.Mask); ipnet.Contains(ip); increment(ip) {
		copiedIP := make([]byte, len(ip))
		copy(copiedIP, ip)
		ips = append(ips, copiedIP)
	}

	r.IPs = ips
	r.Mask = ipnet.Mask
	r.Network = ipnet.IP
	r.Broadcast = broadcastAddress(ipnet.IP, ipnet.Mask)

	return r, nil
}

// UsableIPs returns only usable IPs by filtering out
// the network and broadcast addresses
func (r *IPv4Range) UsableIPs() []net.IP {
	ips := make([]net.IP, len(r.IPs))
	copy(ips, r.IPs)

	// Filter out Network address
	ips = append(ips[:0], ips[1:]...)
	// Filter out Broadcast address
	ips = append(ips[:len(ips)-1], ips[len(ips):]...)

	return ips
}

// RemoveIP manually removes an IP from an IPv4Range
func (r *IPv4Range) RemoveIP(ip string) bool {
	remIP := net.ParseIP(ip)
	if remIP == nil {
		return false
	}

	for idx, val := range r.IPs {
		if val.Equal(remIP) {
			r.IPs = append(r.IPs[:idx], r.IPs[idx+1:]...)
		}
	}

	return true
}

func broadcastAddress(ip net.IP, mask net.IPMask) net.IP {
	return net.IPv4(
		ip[0]|^mask[0],
		ip[1]|^mask[1],
		ip[2]|^mask[2],
		ip[3]|^mask[3])
}

func increment(ip net.IP) {
	for i := len(ip) - 1; i >= 0; i-- {
		ip[i]++
		if ip[i] > 0 {
			break
		}
	}
}
