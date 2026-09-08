# Case Study: URL Shortener

## Requirements

### Functional Requirements

- Given a long URL, generate a short alias (e.g., `short.ly/abc123`).
- Redirect requests for a short alias to the original long URL.
- Aliases should not be predictable/guessable sequentially (avoid easy
  enumeration), and optionally, users can supply a custom alias.
- Aliases expire after a configurable time (optional feature).

### Non-Functional Requirements

- **High availability** — a redirect failure is highly visible and
  damaging (broken links everywhere the short URL was shared).
- **Low latency** on redirect (this is the hot path — reads vastly
  outnumber writes).
- **Read-heavy**: expect roughly a 100:1 to 1000:1 read:write ratio (a link
  is created once, clicked many times).
- Uniqueness: no two long URLs should silently collide on the same short
  alias.

## Capacity Estimation

Assume 100 million new URLs/month, 10-year retention, average record size
~500 bytes, 100:1 read:write ratio.

- Write QPS: `100,000,000 / (30 × 86,400) ≈ 39` writes/sec average.
- Read QPS: `39 × 100 ≈ 3,900` reads/sec average; with a 3x peak factor,
  ~11,700 peak reads/sec.
- Total URLs over 10 years: `100,000,000 × 12 × 10 = 12` billion.
- Storage: `12,000,000,000 × 500 bytes ≈ 6 TB` (before replication, ~18 TB
  at 3x replication).

(See `system-design/estimation/estimation-cheatsheet.md` for the general
formulas used here.)

## API Design

```
POST /api/v1/shorten
  body: { "longUrl": "https://example.com/very/long/path", "customAlias": "optional", "ttl": "optional" }
  response: { "shortUrl": "https://short.ly/abc123" }

GET /{alias}
  response: 301/302 redirect to the original long URL
```

`301` (permanent) redirect caches at the browser/CDN, reducing load on the
service for repeat visits but making analytics on click-through harder
(the browser may not hit your server on subsequent clicks); `302`
(temporary) always hits the server, trading some load for accurate
click tracking. This is a real trade-off to raise explicitly.

## High-Level Architecture

```
Client -> Load Balancer -> App Servers -> Cache (hot aliases) -> Database
                                       -> ID Generation Service
```

- **App servers**: stateless, handle both the shorten and redirect paths.
- **Cache**: a read-through cache (e.g., Redis) in front of the database
  for hot aliases — the read-heavy workload makes this the single highest-
  leverage component.
- **ID generation service**: produces unique short codes (see below) —
  can be a separate service or logic embedded in app servers, depending on
  the chosen ID strategy.

## Data Model

```
urls
  alias        VARCHAR(10) PRIMARY KEY
  long_url     TEXT NOT NULL
  created_at   TIMESTAMP
  expires_at   TIMESTAMP NULL
  created_by   VARCHAR NULL  -- optional, if users are tracked
```

A key-value store (or a relational table used purely as one, keyed by
`alias`) is sufficient — there's no need for joins in the core hot path.

## Core Components

### Short Code Generation

Three common approaches, each with trade-offs:

1. **Base62 encode an auto-incrementing ID**: simple, guarantees
   uniqueness, but sequential IDs are predictable/enumerable (a privacy/
   abuse concern) and require a centralized ID generator (a potential
   bottleneck/single point of failure unless it's itself made
   distributed, e.g. via pre-allocated ID ranges per app server).
2. **Hash the long URL (e.g., MD5/SHA-256) and take the first N
   characters, base62-encoded**: no central coordination needed, but
   collisions are possible (two different URLs hashing to the same
   prefix) and must be detected and handled (e.g., append a salt/retry
   with a different prefix length on collision).
3. **Random generation with a uniqueness check**: generate a random
   Base62 string of fixed length (e.g., 7 characters ≈ 62⁷ ≈ 3.5 trillion
   possibilities), check the database for a collision, retry on the rare
   case one occurs. Simple and avoids predictability; collision
   probability is low enough at this ID space size to make retries rare.

This design recommends **option 3** (random + uniqueness check) as the
default: it avoids both the central-coordinator bottleneck of option 1 and
the collision-handling complexity of option 2, at the cost of an extra
existence check per creation (cheap, and creation is the low-volume path
anyway).

### Redirect Path (Hot Path)

1. Look up `alias` in the cache.
2. On a cache hit, redirect immediately.
3. On a cache miss, look up in the database, populate the cache, then
   redirect.

## Data Flow

**Write (shorten)**: client submits a long URL → app server generates a
candidate alias → checks the database for a collision (retry if needed) →
writes the mapping → returns the short URL.

**Read (redirect)**: client requests `/{alias}` → app server checks cache
→ on hit, redirect; on miss, query database, populate cache, redirect.

## Scaling Strategy

- **Cache**: the single most impactful lever given the read-heavy profile
  — a well-sized cache (sized to the hot working set, not the full 12
  billion URLs — see the estimation cheatsheet) can absorb the vast
  majority of read traffic.
- **Database**: shard by `alias` (e.g., hash the alias to pick a shard) —
  since almost all queries are point lookups by `alias`, this avoids
  cross-shard queries entirely on the hot path.
- **Read replicas**: additional read scaling beyond the cache layer, for
  cache misses.
- **CDN**: for `301` redirects, a CDN can cache the redirect response
  itself, further offloading the origin.

## Failure Scenarios

- **Cache node failure**: read traffic falls through to the database;
  ensure the database (with replicas) can absorb the temporary increase in
  direct traffic without cascading failure — a good discussion point is
  cache warming/pre-population strategies after a cache restart.
- **Database shard failure**: only aliases on that shard are affected if
  properly isolated; replicate each shard for availability.
- **ID generation collision spike**: monitor the collision-retry rate as
  the ID space fills up over years — a rising retry rate signals it's time
  to increase the alias length.

## Bottlenecks

- Database write throughput at very high creation rates (mitigated by
  sharding).
- Cache eviction thrashing if the working set is much larger than assumed
  (mitigated by monitoring cache hit rate and resizing/tuning eviction
  policy).

## Trade-offs

- **301 vs. 302 redirects**: caching/offload vs. accurate click analytics.
- **Random ID generation vs. sequential**: avoids predictability at the
  cost of an extra existence check per write (acceptable given the
  read-heavy workload).
- **Custom aliases**: adds a uniqueness-conflict UX case (must tell the
  user their requested alias is taken) but is a common, expected feature.

## Possible Improvements

- Analytics pipeline (click counts, geographic breakdown) — likely an
  asynchronous pipeline (e.g., publish a click event to a queue, process
  separately) so analytics recording never adds latency to the hot
  redirect path.
- Rate limiting on the shorten endpoint to prevent abuse (see
  `rate-limiter.md`).
- Expiration/garbage collection for TTL'd links, run as a background job
  rather than synchronously on read.

## Interview Discussion Points

- Why is this workload read-heavy, and how does that shape almost every
  design decision here (caching, replica count, redirect status code)?
- How would you detect and handle a short-code collision under each of the
  three ID-generation strategies?
- How would you shard the database, and why does sharding by `alias` avoid
  cross-shard queries here specifically?
- How would you add click-analytics without slowing down the redirect hot
  path?
