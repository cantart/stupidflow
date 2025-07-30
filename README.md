# Stupid Flow

Stupid Flow is a lightweight task flow simulator written in Go.

## Features
- YAML-based DAG config
- Injected workloads: sleep, CPU burn, memory spike, fake I/O
- CPU and memory profiler (pprof)
- Optionally run tasks inside a Docker container with `--docker-image`
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

Run the DAG locally:

```bash
go run . --config path/to/dag.yaml
```

Run the DAG inside a Docker container (requires Docker):

```bash
go run . --config path/to/dag.yaml --docker-image alpine
```

Profiles are stored on the host as `cpu.pprof` and `mem.pprof`.

Start the pprof server with `--pprof` and access profiles at `http://localhost:6060/debug/pprof/`.
