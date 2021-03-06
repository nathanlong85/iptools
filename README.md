# iptools
This package suppliments the existing net package with convenience methods and the ability to deal with additional types of network objects.

###ipv4range
ipv4range allows you to create and work with an ipv4 subnet.

Examples:

```
import "github.com/nathanlong85/iptools/ipv4range"

// Create new range
r, err := ipv4range.New("192.168.0.0/23")

// Returns []net.IP with all available IPs in subnet
// (excludes broadcast and network addresses by default)
avail := r.Available()

// Returns all unavailable IPs that have been removed
unavail := r.Unavailable()

// Returns next N availble IPs in a []net.IP
nextOne, err := r.NextAvailable(1)
nextFive, err := r.NextAvailable(5)

// Remove an available IP from a subnet's list of AvailableIPs
ok := r.Remove(net.ParseIP("192.168.0.1"))

// Returns all IPs in subnet, including unavailable ones, in a []net.IP
all := r.All()

// Returns broadcast address of network
bcast := r.Broadcast()

// Returns network address of network
network := r.Network()

// Returns netmask of network
netmask := r.Mask()
```
