package ipv4range

import "testing"

func TestNew(t *testing.T) {
	r, err := New("192.168.0.0/24")
	if err != nil {
		t.Fatal(err)
	}

	if len(r.IPs) != 256 {
		t.Fatal("Incorrect number of IP addresses returned for /24 test range")
	}
}
