# Project TODOs

## Testing & Automation First
- [ ] Smoke-test harness for `core` (`go test ./core -run TestSmoke` → pass before engine code starts).
- [ ] Table-driven `BaseNode.Validate` tests (cases slice with missing `id` / bad config → `len(errors)` matches expectation).
- [ ] Integration test for `core.RunDAG` with the three base node types (fixture DAG → summary `{Total:3, Failed:0}`).
- [ ] CI workflow running `go test ./...` on push (GitHub Actions YAML → green check on `main`).

## Core Engine Package
- [ ] Scaffold `core` package with README (create `core/` dir → compilable `package core` + documented boundaries).
- [ ] Add `BaseNode` struct (Go file with `ID string` and `Config map[string]any` → exported type with JSON tags).
- [ ] Implement `(*BaseNode).Validate()` (inputs: sample nodes → output: slice describing missing IDs or config errors).
- [ ] Define `Node` interface (`type Node interface { Base() *BaseNode; Execute(context.Context) Result }` → compiler accepts).
- [ ] Implement `Start`, `Task`, and `Control` nodes (structs embedding `BaseNode` → satisfy `Node`; e.g., `Start` takes DAG metadata → emits initial payload, `Task` consumes `map[string]any` → returns success/failure, `Control` evaluates condition/loop config → routes to next node ID).
- [ ] Provide node registry (`core.RegisterNode(kind, factory)` call → custom node instantiation works in tests).
- [ ] Introduce IO mapping table (`map[string]any` JSON payload → nodes read/write via simple hash lookup).
- [ ] Expose `RunDAG` orchestrator (DAG spec with base + custom nodes → returns `Summary`).

## DAG & Scheduling
- [ ] Describe `DAGSpec` (YAML example with nodes/edges → Go struct with YAML tags).
- [ ] Load DAG from file (`core.LoadDAG(path)` with valid / invalid YAML → struct or `ErrInvalidDAG`).
- [ ] Build dependency scheduler (DAG with fan-out + `MaxParallel:2` → execution order honors dependencies).
- [ ] Propagate cancellation (cancel context mid `TaskNode` → summary flags canceled tasks).

## Scaling Experiments
- [ ] Author benchmark DAG (`benchmarks/base_vs_app.yaml` → pipeline mixing base + app nodes stored).
- [ ] Add benchmark `BenchmarkRunDAG_BaseVsApp` (compare 1 vs 8 workers → ns/op delta logged).
- [ ] Summarize results in `docs/scaling.md` (benchmark CSV/table → documented throughput comparison).

## Observability & Diagnostics
- [ ] Emit metrics (`core_node_duration_seconds` per node kind → Prometheus sample visible in tests).
- [ ] Add structured logs (`{"node":"Task","event":"completed"}` → JSON log line from helper).
- [ ] Wire tracing hooks (OpenTelemetry API calls → spans created per node execution).

## Execution Environment
- [ ] Goroutine executor honoring `MaxParallel` (DAG with `MaxParallel:1` → sequential log).
- [ ] Container executor stub (`RunContainer` without config → returns `ErrContainerDisabled`).
- [ ] Resource hints in scheduler (node with `Resources.CPU:2` → scheduler blocks when capacity hit).

## Interfaces & Applications
- [ ] CLI at `cmd/stupidflow/main.go` (`go run ./cmd/stupidflow --dag sample.yaml` → prints `Summary: Total=4, Failed=0`).
- [ ] Go API example (`examples/app/main.go` registering custom node → log shows custom node executed).
- [ ] README quickstart snippet (`import "github.com/.../stupidflow/core"` → instructions for embedding engine).

## Stretch Goals
- [ ] Node extension docs in `docs/nodes.md` (markdown draft → examples showing registry usage).
- [ ] Web dashboard prototype (HTTP handler → HTML table listing node states).
- [ ] Distributed execution proposal (`docs/distributed.md` → RPC design write-up).

---

**Current Snapshot:** We have only 2 packages today (`main` and `core`). See `DESIGN_DECISIONS.md` for the latest architectural choices.
