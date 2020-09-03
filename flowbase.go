package flowbase

import (
	"fmt"
)

var (
	BUFSIZE = 128
)

// ----------------------------------------------------------------------------
// IP
// ----------------------------------------------------------------------------

type IP interface{}

type BaseIP struct {
	data []byte
}

// ----------------------------------------------------------------------------
// Port
// ----------------------------------------------------------------------------

type Port interface {
	Name() string
}

func NewBasePort(name string) *BasePort {
	return &BasePort{
		name:        name,
		remotePorts: make(map[string]Port),
		channel:     make(chan IP, BUFSIZE),
	}
}

type BasePort struct {
	name        string
	remotePorts map[string]Port
	channel     chan IP
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
	return &OutPort{
		BasePort: *NewBasePort(name),
	}
}

type OutPort struct {
	BasePort
}

func (op *OutPort) To(ipt *OutPort) {
	op.Connect(ipt)
}

func (op *OutPort) Send(ip IP) {}

func (op *OutPort) Close() {
	close(op.channel)
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
		name:     name,
		inPorts:  make(map[string]*InPort),
		outPorts: make(map[string]*OutPort),
	}
}

type BaseProcess struct {
	name     string
	inPorts  map[string]*InPort
	outPorts map[string]*OutPort
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
		processes:   []Process{},
		sink:        NewBaseProcess("sink"),
	}
}

type BaseNetwork struct {
	BaseProcess
	processes []Process
	sink      Process
}

func (n *BaseNetwork) Add(p Process) {
	n.processes = append(n.processes, p)
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
