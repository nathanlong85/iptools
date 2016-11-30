package ipv4range

import (
	"fmt"
	"net"
)

// TODO: Encapsulate IPV4Range values, make UnavailableIPs and Available fields

// IPv4Range represents an IPv4 subnet
type IPv4Range struct {
	ips       []net.IP
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

	r.ips = ips
	r.Mask = ipnet.Mask
	r.Network = ipnet.IP
	r.Broadcast = broadcast(ipnet.IP, ipnet.Mask)

	return r, nil
}

// All returns a slice of all IPs that were calculated in the subnet.
// This includes broadcast, network, and any removed addresses.
func (r *IPv4Range) All() []net.IP {
	return r.ips
}

// Available returns only usable IPs by filtering out
// the network and broadcast addresses
func (r *IPv4Range) Available() []net.IP {
	ips := make([]net.IP, len(r.ips))
	copy(ips, r.ips)

	// Filter out Network address
	ips = append(ips[:0], ips[1:]...)
	// Filter out Broadcast address
	ips = append(ips[:len(ips)-1], ips[len(ips):]...)

	return ips
}

// Remove manually removes an IP from an IPv4Range
func (r *IPv4Range) Remove(ip string) bool {
	remIP := net.ParseIP(ip)
	if remIP == nil {
		return false
	}

	for idx, val := range r.ips {
		if val.Equal(remIP) {
			r.ips = append(r.ips[:idx], r.ips[idx+1:]...)
		}
	}

	return true
}

// NextAvailable returns the next available IP(s) in the IP Range. The number of
// available IPs returned should be specified as a parameter.
func (r *IPv4Range) NextAvailable(num int) ([]net.IP, error) {
	if len(r.Available()) < num {
		return nil, fmt.Errorf("Requested %d IPs, only %d available", num, len(r.Available()))
	}

	return r.Available()[:num], nil
}

func broadcast(ip net.IP, mask net.IPMask) net.IP {
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
