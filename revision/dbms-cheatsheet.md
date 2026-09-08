# DBMS Cheatsheet

- **WHERE vs. HAVING**: pre-aggregation filter vs. post-aggregation filter
  (HAVING can reference aggregates like `COUNT(*)`).
- **Joins**: INNER (matches only), LEFT/RIGHT (preserve one side, NULL for
  the other), FULL OUTER (preserve both), CROSS (Cartesian product).
- **ACID**: Atomicity, Consistency, Isolation, Durability.
- **Isolation levels** (increasing strictness): Read Uncommitted → Read
  Committed → Repeatable Read → Serializable. Prevents, respectively:
  nothing → dirty reads → + non-repeatable reads → + phantom reads.
- **Normalization**: reduces redundancy (1NF/2NF/3NF), costs joins.
  **Denormalization**: duplicates data, saves joins, costs write
  consistency.
- **Indexes**: B-Tree is the default (supports range queries); hash indexes
  are O(1) but exact-match only. Indexing speeds reads, costs writes.
- **Locks & deadlocks**: same Coffman-conditions framing as OS deadlocks;
  most DBs auto-detect and abort one transaction.
- **Scaling**:
  - **Replication** (leader-follower or multi-leader) → read scaling +
    availability, possible replica lag.
  - **Sharding** (horizontal partitioning by key) → write/storage scaling,
    adds cross-shard query complexity; shard-key choice determines hot
    spots.
- **Window functions**: `ROW_NUMBER`, `RANK`, `DENSE_RANK`, `LAG`/`LEAD`,
  running aggregates via `OVER (...)` — compute per-row without collapsing
  rows the way `GROUP BY` does.

Full notes + SQL examples: `computer-science/databases/README.md`.
