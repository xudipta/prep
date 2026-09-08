# Computer Networks — Revision Notes

## OSI Model

| Layer | Responsibility | Examples |
|---|---|---|
| 7. Application | End-user protocols | HTTP, DNS, FTP |
| 6. Presentation | Data formatting/encryption | TLS (often placed here or at the session layer) |
| 5. Session | Manages sessions/connections | — |
| 4. Transport | End-to-end delivery, reliability | TCP, UDP |
| 3. Network | Routing between networks | IP |
| 2. Data Link | Node-to-node delivery on the same network | Ethernet, MAC addresses |
| 1. Physical | Raw bit transmission | Cables, radio |

In practice, most engineers reason in terms of the simpler **TCP/IP model**
(Application, Transport, Internet, Link), which maps roughly onto the OSI
layers above.

## TCP vs. UDP

| | TCP | UDP |
|---|---|---|
| Connection | Connection-oriented (handshake) | Connectionless |
| Reliability | Guaranteed delivery, ordering, retransmission | Best-effort, no guarantees |
| Overhead | Higher (headers, ACKs, congestion control) | Lower |
| Use cases | Web (HTTP), file transfer, anything needing correctness | Video/audio streaming, DNS, gaming — latency-sensitive, tolerates loss |

## TCP Three-Way Handshake

1. Client sends `SYN` (synchronize, with an initial sequence number).
2. Server responds `SYN-ACK` (acknowledges client's SYN, sends its own).
3. Client sends `ACK` (acknowledges server's SYN).

Connection teardown is a separate four-step process (`FIN`/`ACK` from each
side), since either side can close its half of the connection independently.

## IP

**IPv4**: 32-bit addresses (~4.3 billion), commonly written as
dotted-decimal (`192.168.1.1`). **IPv6**: 128-bit addresses, designed to
solve IPv4 exhaustion, written in hextets (`2001:0db8::1`). Routing moves
packets between networks based on destination IP, hop by hop.

## DNS

Translates human-readable domain names to IP addresses via a hierarchical,
distributed lookup: root servers → TLD servers (`.com`, `.org`) →
authoritative name servers for the specific domain. Heavily cached (by
resolvers and clients) using each record's TTL to reduce lookup latency and
load on authoritative servers.

## DHCP

Dynamically assigns IP addresses (and other network config: subnet mask,
gateway, DNS servers) to devices joining a network, via a
Discover/Offer/Request/Acknowledge exchange, avoiding manual IP
configuration.

## TLS (and HTTPS)

TLS provides encryption, integrity, and authentication (via certificates)
on top of a transport connection (typically TCP). The TLS handshake
negotiates a protocol version and cipher suite, verifies the server's
certificate (chain of trust to a trusted root CA), and establishes shared
session keys (via asymmetric crypto for the handshake, then fast symmetric
crypto for the actual data). HTTPS is simply HTTP running inside a TLS
connection.

## Cookies vs. Sessions

- **Cookie**: a small piece of data the server asks the browser to store
  and resend on subsequent requests to the same domain — used for session
  identifiers, preferences, tracking.
- **Session**: server-side state associated with a client, usually keyed by
  a session ID stored in a cookie. The cookie itself doesn't have to
  contain sensitive data — often just an opaque ID that maps to
  server-side (or signed/stateless, e.g. JWT-based) session data.

## REST

An architectural style for HTTP APIs: resources identified by URLs,
manipulated via standard HTTP methods (GET, POST, PUT/PATCH, DELETE),
ideally stateless (each request carries all context needed; no
server-side session tied to a specific request's processing). Status codes
communicate outcome (2xx success, 4xx client error, 5xx server error).

## WebSockets

A persistent, full-duplex connection between client and server, established
via an HTTP Upgrade handshake, then switching to a lightweight framed
protocol. Used when the server needs to push data to the client without
the client polling (chat apps, live updates, multiplayer games) — unlike
plain HTTP request/response, either side can send a message at any time.

## Load Balancers

Distribute incoming traffic across multiple backend servers to improve
throughput and availability. Common algorithms: round robin, least
connections, IP hash (for session affinity). Can operate at Layer 4
(TCP/UDP, faster, less flexible) or Layer 7 (HTTP-aware, can route based on
path/headers, terminate TLS).

## Proxies

**Forward proxy**: sits in front of clients, making requests on their
behalf (e.g., corporate egress filtering, anonymization). **Reverse
proxy**: sits in front of servers, receiving client requests and routing
them to backend servers (e.g., nginx, also often does TLS termination,
caching, load balancing).

## CDN (Content Delivery Network)

A geographically distributed network of caching servers that serve content
(static assets, and increasingly dynamic content via edge compute) from a
location close to the requesting client, reducing latency and offloading
the origin server. Cache invalidation (knowing when to fetch fresh content
from the origin) is the central hard problem.

## Common Networking Interview Questions

- What happens when you type a URL into a browser and press enter?
  (DNS resolution → TCP handshake → TLS handshake (if HTTPS) → HTTP
  request/response → rendering.)
- Why does TCP need a three-way handshake, not a two-way one? (Both sides
  need to confirm they can send *and* receive — a two-way handshake
  wouldn't let the client confirm the server received its ACK-equivalent.)
- Explain the difference between HTTP/1.1, HTTP/2, and HTTP/3 at a high
  level (head-of-line blocking, multiplexing, and the underlying transport
  — HTTP/3 runs over QUIC/UDP instead of TCP specifically to fix
  transport-level head-of-line blocking).
- When would you choose UDP over TCP?
- What's the difference between a forward and reverse proxy?

## Quick Revision

- **OSI**: Physical → Data Link → Network → Transport → Session →
  Presentation → Application.
- **TCP**: reliable, ordered, connection-oriented (3-way handshake). **UDP**:
  best-effort, connectionless, lower overhead.
- **DNS**: hierarchical name → IP resolution, heavily cached via TTL.
- **TLS**: encryption + authentication (certificates) + integrity, on top
  of a transport connection.
- **REST**: stateless, resource-oriented, standard HTTP methods.
- **WebSockets**: persistent full-duplex connection for server-push use
  cases.
- **Load balancer**: distributes traffic (L4 fast/simple vs. L7
  content-aware).
- **CDN**: edge caching close to users, cache invalidation is the hard
  part.
