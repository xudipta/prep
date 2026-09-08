# Operating Systems — Revision Notes

## Process vs. Thread

A **process** is an independent unit of execution with its own memory
address space, file descriptors, and OS resources. A **thread** is a unit
of execution *within* a process, sharing that process's memory and
resources with other threads in the same process. Threads are cheaper to
create/switch between than processes because there's no address-space
switch, but shared memory means threads must coordinate to avoid races.

**Interview angle:** why are threads "lighter weight"? Because creating a
thread doesn't require setting up a new page table / address space — only a
new stack and register state.

## Process States

Typical states: **New** → **Ready** (waiting for CPU) → **Running** (on
CPU) → **Waiting/Blocked** (waiting on I/O or an event) → **Terminated**. A
process cycles between Ready and Running as the scheduler time-slices the
CPU, and moves to Blocked when it issues a blocking I/O call, returning to
Ready once that I/O completes.

## Context Switching

Saving the current process/thread's CPU register state and loading
another's, so the CPU can switch which execution flow it's running.
Context switches aren't free — saving/restoring registers, flushing/
reloading CPU caches and the TLB (for process switches, since the address
space changes) all cost time. Excessive context switching ("thrashing" at
the scheduler level) can dominate CPU time over actual useful work.

## CPU Scheduling

| Algorithm | Idea | Trade-off |
|---|---|---|
| FCFS (First-Come First-Served) | Run in arrival order | Simple, but a long job blocks short ones ("convoy effect") |
| SJF (Shortest Job First) | Run the shortest job next | Optimal average wait time, but requires knowing job length in advance |
| Round Robin | Fixed time slice per process, cycle through | Fair, good for interactivity; too-short slices waste time on context switching |
| Priority Scheduling | Run highest-priority job next | Risk of starvation for low-priority jobs (mitigated by aging) |
| Multilevel Feedback Queue | Multiple queues with different priorities/slices, jobs move between them based on behavior | Approximates SJF without knowing job length upfront; used by real OS schedulers |

## Synchronization

**Mutex** (mutual exclusion lock): ensures only one thread executes a
critical section at a time. **Semaphore**: a counter-based primitive that
allows up to N threads into a section concurrently (a mutex is a binary
semaphore, semantically, though implementations differ — a mutex typically
also tracks ownership).

## Deadlock

Four necessary conditions (Coffman conditions), all must hold
simultaneously for deadlock to be possible:

1. **Mutual exclusion** — resources can't be shared.
2. **Hold and wait** — a process holds one resource while waiting for
   another.
3. **No preemption** — resources can't be forcibly taken away.
4. **Circular wait** — a cycle of processes each waiting on the next.

**Prevention**: break any one condition (e.g., impose a global lock
ordering to prevent circular wait). **Detection**: build a wait-for graph
and check for cycles. **Avoidance**: the Banker's algorithm (check if
granting a resource request could ever lead to an unsafe state).

## Race Conditions

Occur when multiple threads access shared data concurrently and the
outcome depends on timing/interleaving. Fixed by synchronization (mutexes,
atomic operations) around the shared state, or by avoiding sharing
altogether (message passing, immutable data, thread-local state).

## Virtual Memory

Gives each process the illusion of a large, contiguous, private address
space, backed by physical memory (and disk, for swapped-out pages) via a
translation layer (page tables). Benefits: process isolation, the ability
to run programs larger than physical RAM, and simplified memory allocation
(each process's addresses start from a similar layout).

## Paging

Divides virtual and physical memory into fixed-size **pages**/**frames**.
The **page table** maps virtual pages to physical frames. A **page fault**
occurs when a referenced page isn't currently in physical memory, requiring
the OS to load it (possibly evicting another page first). The
**Translation Lookaside Buffer (TLB)** caches recent virtual-to-physical
translations to avoid a full page-table walk on every memory access.

## Segmentation

Divides memory into variable-sized logical segments (code, stack, heap),
each with its own base/limit — closer to how programmers think about
memory than fixed-size pages, but more prone to external fragmentation.
Modern systems often combine segmentation and paging.

## Page Replacement

When physical memory is full and a new page must be loaded, an existing
page must be evicted:

| Algorithm | Idea |
|---|---|
| FIFO | Evict the oldest-loaded page |
| LRU (Least Recently Used) | Evict the page unused for the longest time — approximates the optimal policy well in practice |
| Optimal (Belady's) | Evict the page that won't be used for the longest time in the future — theoretical baseline, requires future knowledge |
| Clock (Second-Chance) | Approximates LRU cheaply using a reference bit and a circular scan |

## Memory Allocation

- **Contiguous allocation** strategies: first-fit, best-fit, worst-fit —
  trade off allocation speed against fragmentation.
- **Fragmentation**: **external** (free memory scattered in small chunks
  that individually can't satisfy a request) vs. **internal** (allocated
  memory larger than requested, wasting the difference — common with
  fixed-size pages/blocks).

## System Calls & User vs. Kernel Mode

A **system call** is the controlled interface a user-space program uses to
request a service from the kernel (file I/O, process creation, network
access). The CPU has (at least) two privilege levels: **user mode**
(restricted — can't directly access hardware or other processes' memory)
and **kernel mode** (full privilege). A system call triggers a controlled
mode switch so the kernel can perform the privileged operation on the
process's behalf, then returns control to user mode.

## File Systems

Organize data on persistent storage into files and directories, tracking
metadata (permissions, timestamps, block locations) typically via
**inodes** (Unix-style) that separate a file's metadata/location from its
name (allowing hard links: multiple names for one inode).

## IPC (Inter-Process Communication)

| Mechanism | Notes |
|---|---|
| Pipes | Unidirectional byte stream between related processes |
| Message Queues | Structured messages, decoupled sender/receiver |
| Shared Memory | Fastest (no copying), but requires explicit synchronization |
| Sockets | IPC across machines (or locally via Unix domain sockets) |
| Signals | Asynchronous notifications (e.g., `SIGKILL`, `SIGTERM`) |

## Concurrency

See [`computer-science/concurrency/README.md`](../concurrency/README.md)
for goroutines, channels, and Go-specific concurrency primitives built on
top of these OS-level concepts.

## Interview Questions

- Explain the difference between a process and a thread, and why context
  switching between threads is cheaper.
- Walk through the four Coffman conditions and how to break each one to
  prevent deadlock.
- What happens on a page fault, end to end?
- Why is LRU a good page-replacement heuristic, and how would you
  implement it efficiently (hash map + doubly linked list — see
  `data-structures/hash-map.md` and `data-structures/linked-list.md`)?
- What's the difference between a mutex and a semaphore?

## Quick Revision

- **Process vs. thread**: separate address space vs. shared address space
  within a process.
- **Scheduling**: FCFS, SJF, Round Robin, Priority, Multilevel Feedback
  Queue.
- **Deadlock (Coffman conditions)**: mutual exclusion, hold-and-wait, no
  preemption, circular wait — break any one to prevent deadlock.
- **Virtual memory**: page tables map virtual → physical; TLB caches
  translations; page faults trigger loading from disk.
- **Page replacement**: LRU is the practical standard; Belady's Optimal is
  the theoretical baseline.
- **User vs. kernel mode**: system calls are the controlled boundary
  between them.
