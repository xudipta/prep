# Case Study: Rate Limiter

## Requirements

### Functional Requirements

- Limit the number of requests a client (identified by API key, user ID,
  or IP) can make within a time window.
- Reject (or queue/delay) requests exceeding the limit, typically with an
  HTTP `429 Too Many Requests` response.
- Support different limits for different clients/tiers (e.g., free vs.
  paid).

### Non-Functional Requirements

- **Low latency**: the rate-limit check sits on every request's hot path —
  it must add negligible overhead.
- **Accuracy vs. cost trade-off**: perfectly precise limiting across a
  distributed fleet of servers requires coordination (a shared store),
  which adds latency — most real systems accept a small amount of
  imprecision for speed.
- **Consistency across a distributed set of application servers**: a
  client's requests may hit different servers; the limit must be
  (approximately) global, not per-server, or clients could bypass it by
  spreading requests across servers.

## Capacity Estimation

Assume 10,000 requests/sec across the fleet needing a rate-limit check,
each check requiring one round trip to a shared store (e.g., Redis):

- Redis needs to sustain ~10,000 ops/sec for checks alone — well within a
  single well-provisioned Redis instance's typical capacity (Redis
  commonly handles 100,000+ simple ops/sec), but plan for redundancy
  (replica) for availability, not raw throughput.
- Per-client state size: a counter + timestamp per client per active
  window — for 1 million active clients, a few tens of bytes each,
  totaling tens of MB — comfortably fits in memory.

## API / Interface Design

This is typically a library/middleware rather than a public API:

```go
type RateLimiter interface {
    // Allow reports whether the request identified by key should be
    // permitted under the current rate limit.
    Allow(key string) bool
}
```

Applied as middleware: check `Allow(clientID)` before processing a
request; return `429` if it returns `false`.

## High-Level Architecture

```
Client -> Load Balancer -> App Servers (rate-limit middleware) -> Shared Store (Redis) -> Backend
```

A **shared store** (Redis, or a purpose-built rate-limiting service) is
what makes the limit global across all app servers instead of
per-instance — critical, since a per-server-only limiter would let a
client get `N × (number of servers)` requests through by hitting different
servers.

## Core Algorithms

| Algorithm | Idea | Trade-off |
|---|---|---|
| **Fixed Window Counter** | Count requests in the current fixed time window (e.g., per-minute bucket), reset each window | Simple, O(1) memory per client, but allows up to 2x the limit in a burst straddling a window boundary (e.g., a burst at the end of one window plus a burst at the start of the next) |
| **Sliding Window Log** | Store a timestamp per request; count timestamps within the trailing window | Precise, but memory grows with request volume per client |
| **Sliding Window Counter** | Weighted average of the current and previous fixed windows, approximating a sliding window | Good accuracy/memory trade-off; the common production choice |
| **Token Bucket** | A bucket holds tokens, refilled at a fixed rate; each request consumes a token; requests are rejected if the bucket is empty | Allows controlled bursts up to the bucket size while enforcing a long-term average rate; widely used (e.g., AWS API limits) |
| **Leaky Bucket** | Requests enter a queue processed at a fixed rate; overflow is dropped/rejected | Smooths bursts into a steady output rate, at the cost of added latency for queued requests |

This design recommends **Token Bucket** as the default: it naturally
allows short bursts (good UX for legitimate clients with bursty traffic)
while still enforcing a long-term rate, and is simple to reason about.

### Token Bucket — Sketch

```go
type TokenBucket struct {
    capacity     int64
    tokens       int64
    refillRate   int64 // tokens per second
    lastRefill   time.Time
    mu           sync.Mutex
}

func (b *TokenBucket) Allow() bool {
    b.mu.Lock()
    defer b.mu.Unlock()

    now := time.Now()
    elapsed := now.Sub(b.lastRefill).Seconds()
    b.tokens = min64(b.capacity, b.tokens+int64(elapsed*float64(b.refillRate)))
    b.lastRefill = now

    if b.tokens > 0 {
        b.tokens--
        return true
    }
    return false
}
```

In a distributed deployment, this state (tokens, lastRefill) lives in
Redis (via an atomic Lua script or `INCR`/`EXPIRE` combination) instead of
a local struct, so all app servers see a consistent view.

## Data Model (Redis)

```
key:   ratelimit:{clientID}
value: { tokens: int, lastRefill: timestamp }  -- or a simple counter for fixed/sliding window counter approaches
TTL:   set to the window size, so idle clients' state expires automatically
```

## Data Flow

1. Request arrives at an app server with client identifier `key`.
2. Middleware calls the rate limiter's `Allow(key)`.
3. The limiter atomically reads/updates the client's token count in Redis
   (via a Lua script to avoid a read-then-write race across concurrent
   requests for the same client).
4. If allowed, the request proceeds to the backend; if not, return `429`
   (optionally with a `Retry-After` header).

## Scaling Strategy

- Redis is fast enough for very high request rates on a single instance,
  but for extreme scale, shard by `clientID` hash across multiple Redis
  instances/clusters — each client's state only ever needs one shard, so
  this requires no cross-shard coordination.
- Consider a **local + global** hybrid: an approximate local counter per
  app server for a fast-path check (reject obviously-over-limit requests
  without a network call), backed by periodic reconciliation against the
  global store — trades some precision for reduced Redis load, appropriate
  when perfect precision isn't required.

## Failure Scenarios

- **Redis unavailable**: decide a fail-open (allow all requests — protects
  availability, risks abuse during the outage) vs. fail-closed (reject all
  requests — protects backend from overload, risks unnecessarily blocking
  legitimate traffic during a transient store outage) policy. Fail-open is
  more common for rate limiting specifically, since backend overload
  protection usually has other layers (autoscaling, circuit breakers) as
  backstops.
- **Clock skew across app servers**: matters for time-window-based
  algorithms; token bucket refill computed relative to a stored
  `lastRefill` timestamp is more robust to this than wall-clock-window
  approaches, but Redis's own server time (via `TIME` command) can be used
  as a single source of truth if precision matters.

## Bottlenecks

- A single very high-volume client (e.g., a misbehaving script) causing
  disproportionate load on its shard — mitigate with a maximum burst cap
  independent of the bucket size, or additional coarser-grained limits
  (e.g., per-IP in addition to per-API-key).

## Trade-offs

- **Precision vs. latency/cost**: sliding window log is most precise but
  most expensive; fixed window is cheapest but allows boundary bursts;
  token bucket and sliding window counter are the practical middle ground.
- **Fail-open vs. fail-closed** on rate-limiter-store failure.
- **Centralized (Redis) vs. hybrid local+global**: precision vs. Redis
  load reduction.

## Possible Improvements

- Per-endpoint limits in addition to per-client limits (protect
  specific expensive endpoints more aggressively).
- Dynamic limits based on client tier, adjustable without a deploy (store
  tier limits in a fast-lookup config store, not hardcoded).
- Return standard rate-limit headers (`X-RateLimit-Limit`,
  `X-RateLimit-Remaining`, `Retry-After`) so well-behaved clients can
  self-throttle.

## Interview Discussion Points

- Compare all five algorithms above and justify picking Token Bucket (or
  argue for a different one given different stated priorities).
- Why must the rate-limit state be shared (not per-app-server) in a
  load-balanced deployment?
- What happens to a client's UX under fail-open vs. fail-closed during a
  Redis outage, and which would you choose for a payments API vs. a public
  read-only API?
