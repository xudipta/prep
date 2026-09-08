# Networking Cheatsheet

- **OSI layers**: Physical → Data Link → Network → Transport → Session →
  Presentation → Application.
- **TCP**: reliable, ordered, connection-oriented (3-way handshake:
  SYN → SYN-ACK → ACK). **UDP**: best-effort, connectionless, low overhead
  — video/audio streaming, DNS, gaming.
- **DNS**: hierarchical name → IP resolution (root → TLD → authoritative),
  cached via TTL.
- **DHCP**: dynamic IP assignment (Discover/Offer/Request/Acknowledge).
- **TLS/HTTPS**: encryption + authentication (certificate chain of trust)
  + integrity, layered on top of a transport connection.
- **Cookies vs. sessions**: cookie is client-stored data sent with each
  request; session is server-side state usually keyed by a cookie-held ID.
- **REST**: stateless, resource-oriented, standard HTTP verbs, status codes
  communicate outcome.
- **WebSockets**: persistent full-duplex connection, used for server-push
  use cases (chat, live updates) that plain request/response can't serve
  efficiently.
- **Load balancer**: L4 (fast, TCP/UDP-level) vs. L7 (content-aware,
  HTTP-level, can terminate TLS/route by path).
- **Forward vs. reverse proxy**: sits in front of clients vs. in front of
  servers.
- **CDN**: geographically distributed caching close to users; cache
  invalidation is the hard problem.

Full notes: `computer-science/networking/README.md`.
