package main

import (
	"fmt"
)

type workloadType string

const (
	workloadSleep       workloadType = "sleep"
	workloadCPUBurn     workloadType = "cpu_burn"
	workloadMemorySpike workloadType = "memory_spike"
	workloadFakeIO      workloadType = "fake_io"
)

type dag struct {
	vertexes []vertex
	edges    []edge
}

func (d *dag) run() {
	// Detect cycles -> Dag or not?
	// Find executable vertexes (i.e., vertexes with no incoming edges)
	// Execute vertexes with injected workloads
	fmt.Println("Running DAG...")
}

type vertex struct {
	id       uint8
	workload workloadType
}

type edge struct {
	from uint8
	to   uint8
}

func main() {
	v0 := vertex{id: 0, workload: workloadCPUBurn}
	v1 := vertex{id: 1, workload: workloadMemorySpike}
	v2 := vertex{id: 2, workload: workloadFakeIO}
	v3 := vertex{id: 3, workload: workloadSleep}
	e01 := edge{from: 0, to: 1}
	e12 := edge{from: 1, to: 2}
	e13 := edge{from: 1, to: 3}
	d := dag{vertexes: []vertex{v0, v1, v2, v3}, edges: []edge{e01, e12, e13}}
	fmt.Printf("DAG: %+v\n", d)
}
