# Concurrency Cheatsheet (Go)

- **Concurrency ≠ parallelism**: structuring independent tasks vs.
  literally running them at the same instant (needs multiple cores).
- **Goroutines**: cheap, runtime-scheduled; the main goroutine exiting does
  NOT wait for others — use `sync.WaitGroup` or channels.
- **Channels**: unbuffered = synchronization point (send blocks until a
  receiver is ready); buffered = async up to capacity.
- **Mutex/RWMutex**: protect shared mutable state directly accessed by
  multiple goroutines; `RWMutex` when reads ≫ writes.
- **atomic**: lock-free single-variable ops (`atomic.Int64`, etc.) — faster
  than a mutex for simple counters.
- **Race conditions**: unsynchronized concurrent access with ≥1 writer;
  detect with `go test -race` (doesn't catch every possible race, only
  ones that occur during the run).
- **Deadlock**: goroutines waiting on each other in a cycle; Go detects
  *global* deadlock (all goroutines asleep) and panics, not partial
  deadlocks.
- **Worker pool**: fixed number of goroutines reading from a shared job
  channel — bounds concurrency.
- **Fan-out/fan-in**: distribute work across goroutines / merge multiple
  channels into one.
- **context.Context**: standard cancellation/deadline propagation across
  goroutines and API boundaries.
- **When to use channels vs. mutex**: channels for coordinating/signaling
  between goroutines (pipelines); mutexes for protecting shared state
  accessed directly by multiple goroutines — both are idiomatic Go.

Full notes + code: `computer-science/concurrency/README.md`.
