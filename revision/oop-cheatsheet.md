# OOP Cheatsheet

- **Encapsulation**: bundle data + behavior; restrict access (Go: exported
  vs. unexported, per-package).
- **Abstraction**: hide complexity behind a simpler interface.
- **Inheritance**: not in Go — use composition (embedding) + interfaces.
- **Polymorphism**: Go interfaces are structural — any type with the right
  methods satisfies them, no `implements` needed.
- **Composition > Inheritance**: loose coupling; Go has no inheritance at
  all by design.
- **Association / Aggregation / Composition (UML)**: uses-a / has-a
  (independent lifetime) / has-a (owned lifetime).
- **Overloading/Overriding**: neither exists in classical form in Go.
- **Static vs. dynamic binding**: concrete-type calls resolve at compile
  time; interface method calls resolve at runtime.
- **SOLID**:
  - **S**ingle Responsibility — one reason to change.
  - **O**pen/Closed — extend via new implementations, don't modify existing
    code.
  - **L**iskov Substitution — implementations must honor the interface's
    behavioral contract, not just its signatures.
  - **I**nterface Segregation — many small interfaces over one large one.
  - **D**ependency Inversion — depend on interfaces, not concrete types.
- **DRY / KISS / YAGNI**: don't duplicate / keep it simple / don't build
  for hypothetical futures.
- **Coupling vs. Cohesion**: want low coupling, high cohesion.
- **Dependency Injection**: pass interfaces into constructors instead of
  constructing dependencies internally.

Full notes + Q&A: `computer-science/oop/README.md`. Design patterns:
`system-design/design-patterns/README.md`.
