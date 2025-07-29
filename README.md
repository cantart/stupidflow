# Stupid Flow

Stupid Flow is a lightweight task flow simulator written in Go.

## Features
- YAML-based DAG config
- Injected workloads: sleep, CPU burn, memory spike, fake I/O
- CPU and memory profiler (pprof)
- Lightweight dockerized execution via `--docker-image`
- Profiles are written to `cpu.pprof` and `mem.pprof` by default (use `--profile` and `--mem-profile` to change)

## Usage
Create a YAML file describing a single task:

```yaml
tasks:
  - id: only
    type: cpu_burn
    duration: 1s
```

The configuration must contain exactly one task.

Run the simulator:

```bash
go run . --config path/to/dag.yaml
```

CPU and memory profiles will be saved to `cpu.pprof` and `mem.pprof`.

Inspect a profile with `go tool pprof`:

```bash
go tool pprof cpu.pprof
go tool pprof mem.pprof
```

Pass `-http=:8080` to open a browser-based viewer.

Run the entire DAG inside a Docker container (requires Docker):

```bash
go run . --config path/to/dag.yaml --docker-image alpine
```

Profiles are stored on the host as `cpu.pprof` and `mem.pprof`.

Start the pprof server with `--pprof` and access profiles at `http://localhost:6060/debug/pprof/`.
