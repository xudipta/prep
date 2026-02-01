# Prep — Go Practice Workspace

This repository is a small workspace for practicing Go coding problems.

Structure

- `go.mod` — Go module file.
- `README.md` — this file.
- `problems/` — each problem should live in its own `problem-name.go` file, package `problems`.
- `utils/` — reusable helper functions for solving problems (`custom_lib.go`).
- `playground/` — quick testbed for snippets (`main.go`, `examples.go`, `run.sh`).

How to use

1. Write problem solutions in `problems/`.
2. Add helper functions to `utils/custom_lib.go`.
3. Use `playground/main.go` to quickly run and test code snippets.

Run the playground:

```bash
# make the script executable once
chmod +x ./playground/run.sh
# run
./playground/run.sh
```
