package flowbase

import (
	"fmt"
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

func TestRun(t *testing.T) {
	n := NewBaseNetwork("testrun-network")

	hw := NewHelloWorldProc("hw")
	n.Add(hw)

	n.Run()
}

// ----------------------------------------------------------------------------
// Helpers
// ----------------------------------------------------------------------------

func NewHelloWorldProc(name string) *HelloWorldProc {
	return &HelloWorldProc{
		BaseProcess{
			name:    name,
			inPorts: make(map[string]*InPort),
			outPorts: map[string]*OutPort{
				"out": NewOutPort("hw"),
			},
		},
	}
}

type HelloWorldProc struct {
	BaseProcess
}

func (p *HelloWorldProc) Run() {
	defer p.outPorts["hw"].Close()
	for i := 1; i <= 10; i++ {
		p.outPorts["hw"].Send(fmt.Sprintf("Hello world for the %d:th time", i))
	}
}
