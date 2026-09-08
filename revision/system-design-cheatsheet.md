# System Design Cheatsheet

- **CAP theorem**: during a network partition, choose Consistency or
  Availability — not both (outside a partition, you can have both).
- **Scalability**: vertical (bigger machine, hard ceiling) vs. horizontal
  (more machines, needs statelessness/partitioning).
- **Latency vs. throughput**: time-per-request vs. requests-per-time — not
  the same axis; clarify which a requirement actually means.
- **Caching**: cache-aside is the default pattern; LRU is the default
  eviction policy; decide TTL-based vs. explicit invalidation up front.
- **Replication**: read scaling + availability (leader-follower or
  multi-leader). **Sharding**: write/storage scaling by partitioning data;
  shard-key choice determines hot spots and cross-shard query cost.
- **Queue vs. Pub/Sub**: one consumer per message (decoupling, load
  smoothing) vs. fan-out to many subscribers (event broadcasting).
- **Rate limiting algorithms**: Fixed Window (simple, boundary bursts),
  Sliding Window Log (precise, memory-heavy), Sliding Window Counter (good
  middle ground), Token Bucket (allows controlled bursts), Leaky Bucket
  (smooths to a steady rate).
- **Idempotency**: essential for safe retries — use a client-supplied
  idempotency key the server deduplicates against.
- **Observability**: logs (specific events) → metrics (aggregated trends)
  → traces (cross-service request path).
- **Estimation**: QPS = daily requests / 86,400, × a peak factor (2-5x);
  storage = items × size × replication factor; cache for the hot working
  set, not the full dataset. Full formulas:
  `system-design/estimation/estimation-cheatsheet.md`.
- **Design-problem structure**: Requirements (functional/non-functional) →
  Capacity Estimation → API Design → High-Level Architecture → Data Model
  → Core Components → Data Flow → Scaling → Failure Scenarios →
  Bottlenecks → Trade-offs → Improvements.

Full notes: `system-design/fundamentals/README.md`. Worked examples:
`system-design/case-studies/`.
