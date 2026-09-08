# OOP — Quick Revision

- **Encapsulation**: bundle data + behavior, restrict direct access (Go:
  exported vs. unexported identifiers, per-package).
- **Abstraction**: hide complexity behind a simpler interface (what, not
  how).
- **Inheritance**: Go has none — use composition (struct embedding) and
  interfaces instead.
- **Polymorphism**: Go interfaces are structurally typed — any type with
  the right methods satisfies an interface implicitly.
- **Composition > Inheritance**: loose coupling, no fragile base-class
  hierarchies.
- **Association / Aggregation / Composition (UML)**: uses-a / has-a
  (parts outlive whole) / has-a (whole owns parts' lifecycle).
- **Interfaces vs. abstract classes**: Go interfaces carry zero
  implementation; shared logic comes from composition or plain functions.
- **Overloading/Overriding**: neither exists in Go in the classical sense;
  use variadic params or interface params; embedding shadows, doesn't
  override.
- **Static vs. dynamic binding**: concrete-type method calls resolve at
  compile time; interface method calls resolve at runtime.
- **SOLID**: Single Responsibility, Open/Closed, Liskov Substitution,
  Interface Segregation, Dependency Inversion.
- **DRY / KISS / YAGNI**: don't duplicate logic / keep it simple / don't
  build for hypothetical futures.
- **Coupling vs. Cohesion**: want low coupling (narrow interfaces between
  modules), high cohesion (each module does one focused thing).
- **Dependency Injection**: pass interfaces into constructors instead of
  constructing dependencies internally — enables substitution/testing.
