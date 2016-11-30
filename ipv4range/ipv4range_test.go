package ipv4range

import "testing"

func TestNew(t *testing.T) {
	r, err := New("192.168.0.0/24")
	if err != nil {
		t.Fatal(err)
	}

	if len(r.IPs) != 256 {
		t.Fatal("Incorrect number of IPs returned for /24 test range")
	}
}

func TestNextAvailableIP(t *testing.T) {
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

func TestAvailableIPs(t *testing.T) {
	r, err := New("192.168.0.0/24")
	if err != nil {
		t.Fatal(err)
	}

	if len(r.AvailableIPs()) != 254 {
		t.Fatal("Incorrect number of IP available IPs returned for /24 test range")
	}

	for _, ip := range r.AvailableIPs() {
		if ip.Equal(r.Broadcast) || ip.Equal(r.Network) {
			t.Fatal("Available IPs should not include broadcast or network IP")
		}
	}
}
