package flowbase

import (
	"fmt"
)

// ----------------------------------------------------------------------------
// Port
// ----------------------------------------------------------------------------

type Port interface {
	Name() string
}

type BasePort struct {
	name        string
	remotePorts map[string]Port
}

func (bp *BasePort) Name() string {
	return bp.name
}

func (bp *BasePort) Connect(rpt Port) {
	bp.remotePorts[rpt.Name()] = rpt
}

// ----------------------------------------------------------------------------
// InPort
// ----------------------------------------------------------------------------

func NewInPort(name string) *InPort {
	return &InPort{BasePort{name: name}}
}

type InPort struct {
	BasePort
}

func (ip *InPort) From(op *OutPort) {
	ip.Connect(op)
}

// ----------------------------------------------------------------------------
// OutPort
// ----------------------------------------------------------------------------

func NewOutPort(name string) *OutPort {
	return &OutPort{BasePort{name: name}}
}

type OutPort struct {
	BasePort
}

func (op *OutPort) To(ip *OutPort) {
	op.Connect(ip)
}

// ----------------------------------------------------------------------------
// Process
// ----------------------------------------------------------------------------

type Process interface {
	Run()
	Name() string
}

func NewBaseProcess(name string) *BaseProcess {
	return &BaseProcess{
		name: name,
	}
}

type BaseProcess struct {
	name     string
	inPorts  map[string]InPort
	outPorts map[string]OutPort
}

func (bp *BaseProcess) Name() string {
	return bp.name
}

func (bp *BaseProcess) Run() {}

// ----------------------------------------------------------------------------
// Network
// ----------------------------------------------------------------------------

type Network interface {
	Process
}

func NewBaseNetwork(name string) *BaseNetwork {
	return &BaseNetwork{
		BaseProcess: *NewBaseProcess(name),
	}
}

type BaseNetwork struct {
	BaseProcess
	processes []Process
	sink      Process
}

func (n *BaseNetwork) Run() {
	if len(n.processes) == 0 {
		fmt.Printf("WARNING: No processes in network [%s]\n", n.Name())
		return
	}
	for _, p := range n.processes {
		go p.Run()
	}
	n.sink.Run()
}
