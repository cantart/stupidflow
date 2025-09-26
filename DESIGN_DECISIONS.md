# Design Decisions

## Shared Base Node for Engine Primitives
- **Problem:** We need a consistent way to define Start, Task, Control, and future application nodes without duplicating validation or execution plumbing.
- **Possible Solutions:**
  - Hand-roll each node type independently with ad-hoc structs.
  - Create a lightweight base struct that centralizes shared lifecycle hooks and configuration parsing.
  - Rely solely on interfaces and let every node embed its own helpers.
- **Choice:** Introduce a reusable `BaseNode` abstraction that Start, Task, Control, and extension nodes can embed to share validation logic while still allowing specialized behavior.

## Package Layout for the Core Engine
- **Problem:** The project currently ships as a single `main` package, but we need to grow an engine that can be reused outside of the CLI entrypoint.
- **Possible Solutions:**
  - Keep everything in `main`, accepting tighter coupling between the engine and the executable.
  - Break the engine into a `core` package and leave the CLI in `main`.
  - Split into multiple packages (e.g., `core`, `scheduler`, `runtime`) from day one.
- **Choice:** House the engine inside a dedicated `core` package while keeping `main` for the CLI, giving us reuse without premature package sprawl.

## Data Exchange Between Nodes
- **Problem:** Nodes must pass inputs and outputs even before a richer type system exists.
- **Possible Solutions:**
  - Implement a full schema definition language now.
  - Use custom structs for every node pair.
  - Exchange data as JSON-parseable hash tables until stricter typing is required.
- **Choice:** Standardize on JSON-friendly hash tables for early experiments so nodes can communicate with minimal ceremony while leaving room for future schema evolution.
