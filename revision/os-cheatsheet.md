# Operating Systems Cheatsheet

- **Process vs. thread**: separate address space vs. shared address space
  within a process; threads are cheaper to context-switch.
- **Process states**: New → Ready → Running → Waiting/Blocked →
  Terminated.
- **Scheduling**: FCFS (simple, convoy effect), SJF (optimal wait, needs
  foreknowledge), Round Robin (fair, interactive), Priority (starvation
  risk), Multilevel Feedback Queue (approximates SJF adaptively).
- **Deadlock (Coffman conditions)**: mutual exclusion, hold-and-wait, no
  preemption, circular wait — all four required; break any one to prevent.
- **Race condition**: unsynchronized concurrent access to shared state with
  at least one writer — fix with mutexes/atomics or avoid sharing.
- **Virtual memory**: page tables map virtual → physical addresses; TLB
  caches translations; page fault triggers loading from disk.
- **Paging vs. segmentation**: fixed-size pages (internal fragmentation) vs.
  variable-size logical segments (external fragmentation).
- **Page replacement**: LRU (practical standard) vs. Belady's Optimal
  (theoretical baseline, needs future knowledge).
- **User vs. kernel mode**: system calls are the controlled boundary
  between restricted and privileged execution.
- **IPC**: pipes, message queues, shared memory (fastest, needs sync),
  sockets, signals.

Full notes: `computer-science/operating-systems/README.md`.
