# Databases — Revision Notes

## SQL Fundamentals

```sql
SELECT column1, column2 FROM table WHERE condition
GROUP BY column1 HAVING aggregate_condition
ORDER BY column1
LIMIT n;
```

- **WHERE** filters rows before grouping; **HAVING** filters groups after
  aggregation — the single most common SQL interview gotcha (`WHERE` can't
  reference an aggregate like `COUNT(*)`; `HAVING` can).

## Joins

| Join | Result |
|---|---|
| INNER JOIN | Only rows with matches in both tables |
| LEFT (OUTER) JOIN | All rows from the left table; unmatched right columns are NULL |
| RIGHT (OUTER) JOIN | All rows from the right table; unmatched left columns are NULL |
| FULL OUTER JOIN | All rows from both tables; unmatched columns on either side are NULL |
| CROSS JOIN | Cartesian product — every row of the left table paired with every row of the right |

## GROUP BY / HAVING / Subqueries / CTEs / Window Functions

```sql
-- CTE (Common Table Expression) — improves readability of multi-step queries
WITH recent_orders AS (
    SELECT * FROM orders WHERE created_at > NOW() - INTERVAL '30 days'
)
SELECT customer_id, COUNT(*) FROM recent_orders GROUP BY customer_id;

-- Window function — aggregate without collapsing rows
SELECT employee_id, salary,
       RANK() OVER (PARTITION BY department ORDER BY salary DESC) AS dept_rank
FROM employees;
```

Window functions (`ROW_NUMBER`, `RANK`, `DENSE_RANK`, `LAG`/`LEAD`, running
`SUM`/`AVG` with `OVER`) compute per-row values relative to a window of
rows, without collapsing the result set the way `GROUP BY` does — the go-to
tool for "top N per group," running totals, and row-to-row comparisons.

## Indexes

An index is an auxiliary structure (typically a B-Tree, sometimes a hash
index) that speeds up lookups on a column at the cost of extra storage and
slower writes (every insert/update/delete must also update the index).
**Rule of thumb**: index columns used in `WHERE`, `JOIN`, and `ORDER BY`
clauses on large tables; over-indexing hurts write throughput.

## B-Tree Indexes

A balanced tree structure keeping data sorted, with O(log n) search,
insert, and delete, and efficient range queries (unlike a hash index, which
only supports exact-match lookups in O(1) but can't do `WHERE col > x`
efficiently). Most relational database indexes default to B-Tree (or a
B+Tree variant) specifically because of this range-query support.

## Transactions & ACID

- **Atomicity**: a transaction's operations either all happen or none do.
- **Consistency**: a transaction moves the database from one valid state to
  another (constraints/invariants hold before and after).
- **Isolation**: concurrent transactions appear to execute as if serialized
  (to a degree controlled by the isolation level — see below).
- **Durability**: once committed, a transaction's effects survive crashes
  (typically via a write-ahead log).

## Isolation Levels

| Level | Prevents | Allows |
|---|---|---|
| Read Uncommitted | Nothing | Dirty reads, non-repeatable reads, phantom reads |
| Read Committed | Dirty reads | Non-repeatable reads, phantom reads |
| Repeatable Read | Dirty + non-repeatable reads | Phantom reads (in some implementations) |
| Serializable | All of the above | Fully serialized execution (highest cost) |

**Dirty read**: reading uncommitted data from another transaction.
**Non-repeatable read**: re-reading the same row within a transaction gives
a different value because another transaction committed a change.
**Phantom read**: re-running the same query returns a different *set* of
rows because another transaction inserted/deleted matching rows.

## Normalization & Denormalization

**Normalization** organizes tables to eliminate redundancy (via normal
forms — 1NF: atomic columns; 2NF: no partial dependency on a composite key;
3NF: no transitive dependency on non-key columns), reducing update
anomalies at the cost of requiring more joins to reconstruct a full view.
**Denormalization** deliberately duplicates data to avoid expensive joins,
trading write complexity/redundancy for read performance — common in
read-heavy systems and data warehouses.

## Locks & Deadlocks

Databases use locks (row-level, table-level, shared/exclusive) to enforce
isolation. A **deadlock** occurs when two transactions each hold a lock the
other needs — most databases detect this automatically and abort one
transaction (see the OS deadlock notes for the general Coffman-conditions
framing, which applies here too).

## Query Optimization

- Add indexes on frequently filtered/joined/sorted columns.
- Use `EXPLAIN` / `EXPLAIN ANALYZE` to see the query planner's chosen
  strategy (index scan vs. sequential scan, join algorithm) and actual
  costs.
- Avoid `SELECT *` when only specific columns are needed (less data
  transferred, and can enable index-only scans).
- Be wary of functions applied to indexed columns in `WHERE` clauses
  (`WHERE LOWER(email) = ...`) — this typically prevents the index from
  being used unless a matching functional/expression index exists.

## Database Scaling

- **Vertical scaling**: bigger machine (more CPU/RAM/disk) — simple, but
  has a hard ceiling and a single point of failure.
- **Replication**: copy data to multiple nodes.
  - **Leader-follower (primary-replica)**: writes go to the leader,
    reads can be served by replicas — improves read throughput and
    provides failover, but replicas may lag (eventual consistency for
    reads from replicas).
  - **Multi-leader / leaderless**: more write availability, at the cost of
    needing conflict resolution.
- **Sharding (horizontal partitioning)**: split data across multiple nodes
  by a shard key (e.g., user ID range or hash), so each node holds a
  subset — necessary once a single machine can't hold or serve the full
  dataset. Introduces cross-shard query/transaction complexity and
  requires a good shard-key choice to avoid hot spots.

## Representative SQL Interview Problems

- Find the second-highest salary per department (window function
  `DENSE_RANK`, or a correlated subquery).
- Find duplicate rows in a table (`GROUP BY` all columns `HAVING COUNT(*) >
  1`).
- Running total of sales per day (`SUM(...) OVER (ORDER BY date)`).
- Find customers with no orders (`LEFT JOIN` + `WHERE orders.id IS NULL`,
  or `NOT EXISTS`).
- Top N products per category (`ROW_NUMBER() OVER (PARTITION BY category
  ORDER BY sales DESC)`, filtered to `<= N`).

## Quick Revision

- **WHERE vs HAVING**: pre-aggregation filter vs. post-aggregation filter.
- **Joins**: inner (matches only), left/right (preserve one side),
  full (preserve both).
- **ACID**: Atomicity, Consistency, Isolation, Durability.
- **Isolation levels**: Read Uncommitted → Read Committed → Repeatable
  Read → Serializable (increasing consistency, decreasing concurrency).
- **Normalization**: reduces redundancy, costs joins; denormalization is
  the opposite trade.
- **Indexes**: B-Tree for range queries; speeds reads, costs writes.
- **Scaling**: replication (read scaling + availability) vs. sharding
  (write/storage scaling, adds cross-shard complexity).
