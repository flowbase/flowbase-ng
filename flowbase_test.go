package flowbase

import (
	"testing"
)

func TestNetworkName(t *testing.T) {
	wantName := "hey-ho"
	n := NewBaseNetwork(wantName)
	haveName := n.Name()
	if haveName != wantName {
		t.Fatalf("Have name %s but wanted %s", haveName, wantName)
	}
}
