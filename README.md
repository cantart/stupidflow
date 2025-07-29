# Stupid Flow

Stupid Flow is a lightweight task flow simulator written in Go.

## Features
- YAML-based DAG config
- Injected workloads: sleep, CPU burn, memory spike, fake I/O
- CPU and Memory profiler (pprof)
- Lightweight dockerized execution via `--docker-image`
- CPU profile reports via `--report`

## Usage
Create a YAML file describing the tasks:

```yaml
tasks:
  - id: first
    type: sleep
    duration: 1s
  - id: second
    type: cpu_burn
    duration: 1s
    depends_on:
      - first
```

Run the simulator:

```bash
go run . --config path/to/dag.yaml
```

Run the entire DAG inside a Docker container (requires Docker):

```bash
go run . --config path/to/dag.yaml --docker-image alpine
```

Generate and show a CPU profile report:

```bash
go run . --config path/to/dag.yaml --docker-image alpine --report
```

Start the pprof server with `--pprof` and access profiles at `http://localhost:6060/debug/pprof/`.
