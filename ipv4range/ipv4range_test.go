package ipv4range

import (
	"net"
	"testing"
)

func TestNew(t *testing.T) {
	r, err := New("192.168.0.0/24")
	if err != nil {
		t.Fatal(err)
	}

	if len(r.All()) != 256 {
		t.Fatal("Incorrect number of IPs returned for /24 test range")
	}
}

func TestNextAvailable(t *testing.T) {
	r, err := New("192.168.0.0/24")
	if err != nil {
		t.Fatal(err)
	}

	next, err := r.NextAvailable(5)
	if err != nil {
		t.Fatal(err)
	}

	if len(next) != 5 {
		t.Fatal("Not enough available IPs returned")
	}
}

func TestAvailable(t *testing.T) {
	r, err := New("192.168.0.0/24")
	if err != nil {
		t.Fatal(err)
	}

	if len(r.Available()) != 254 {
		t.Fatal("Incorrect number of IP available IPs returned for /24 test range")
	}

	for _, ip := range r.Available() {
		if ip.Equal(r.Broadcast()) || ip.Equal(r.Network()) {
			t.Fatal("Available IPs should not include broadcast or network IP")
		}
	}
}

func TestAll(t *testing.T) {
	r, err := New("192.168.0.0/24")
	if err != nil {
		t.Fatal(err)
	}

	if len(r.All()) != 256 {
		t.Fatal("Incorrect number of IPs returned for /24 test range")
	}
}

func TestRemove(t *testing.T) {
	r, err := New("192.168.0.0/24")
	if err != nil {
		t.Fatal(err)
	}

	r.Remove(net.ParseIP("192.168.0.1"))
	if len(r.Available()) != 253 {
		t.Fatalf("%d available IPs returned. Should be 253", len(r.Available()))
	}
}

func TestUnavailable(t *testing.T) {
	r, err := New("192.168.0.0/24")
	if err != nil {
		t.Fatal(err)
	}

	if len(r.Unavailable()) != 2 {
		t.Fatalf("%d unavailable IPs returned. Should be 2", len(r.Unavailable()))
	}

	r.Remove(net.ParseIP("192.168.0.1"))
	if len(r.Unavailable()) != 3 {
		t.Fatalf("%d unavailable IPs returned. Should be 3", len(r.Unavailable()))
	}
}

func TestBroadcast(t *testing.T) {
	_, ipnet, _ := net.ParseCIDR("192.168.0.0/24")
	bcast := Broadcast(ipnet.IP, ipnet.Mask)
	if !bcast.Equal(net.ParseIP("192.168.0.255")) {
		t.Fatalf("Broadcast %s not correct for 192.168.0.0/24", bcast)
	}
}
