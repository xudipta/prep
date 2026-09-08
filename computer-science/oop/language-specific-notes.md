# OOP — Language-Specific Notes (C++, Java, Go)

The main [`README.md`](README.md) covers OOP *principles* language-agnostically
(and uses Go for examples, since Go is this repo's primary language). Several
classical OOP mechanisms — virtual functions, abstract classes, multiple
inheritance, constructors/destructors, operator overloading — are
implemented very differently across languages, and interviewers often probe
exactly these differences. This page is a language-by-language comparison
for C++, Java, and Go.

## Quick Comparison Table

| Concept | C++ | Java | Go |
|---|---|---|---|
| Classes | `class` keyword | `class` keyword | No classes — `struct` + methods |
| Inheritance | Multiple (classes can inherit from several base classes) | Single (one base class only) | None — composition via struct embedding |
| Interfaces | No separate keyword — a class with only pure virtual functions acts as one | `interface` keyword, explicit `implements` | Implicit/structural — no `implements` needed |
| Virtual dispatch | Explicit: `virtual` keyword, vtable-based | Implicit: all non-static, non-final, non-private methods are virtual by default | No virtual methods — dispatch happens through interface values instead |
| Abstract classes | Class with ≥1 pure virtual function (`= 0`) | `abstract class` keyword | Not supported (use an interface) |
| Access control | `public` / `protected` / `private`, plus `friend` | `public` / `protected` / `private` / package-private (default) | Exported (`Capitalized`) / unexported, per-**package**, not per-type |
| Constructors | Yes, plus copy constructors and move constructors | Yes | No constructors — convention: a `NewX()` function |
| Destructors | Yes (deterministic, RAII-based) | No (finalizers existed, deprecated; GC-based cleanup) | No destructors (GC-based); use `defer` for cleanup |
| Operator overloading | Yes | No (except built-in `+` for `String`) | No |
| Method overloading | Yes (resolved by signature at compile time) | Yes (resolved by signature at compile time) | No — use different names or variadic/interface parameters |
| Memory management | Manual, or RAII / smart pointers (`unique_ptr`, `shared_ptr`) | Garbage collected | Garbage collected |

## Virtual Functions and the vtable

**Question:** What actually happens at runtime when a virtual function is
called?

**Short Answer:** In C++ and Java, each object (well, each class) has a
hidden **virtual method table (vtable)** — an array of function pointers.
A virtual call looks up the correct function pointer through the object's
vtable pointer at runtime, instead of calling a fixed address at compile
time. Go achieves the same *effect* through interface values, but the
mechanism looks different.

**Detailed Explanation:**

- **C++**: virtual functions must be explicitly marked `virtual` in the
  base class. Every object of a class with at least one virtual function
  carries a hidden pointer (traditionally called `vptr`) to its class's
  vtable. Calling a virtual function through a base-class pointer or
  reference indirects through this vtable, which is what allows a
  `Shape*` pointing to a `Circle` to correctly call `Circle::Area()`. Any
  method *not* marked `virtual` is resolved at compile time
  (non-polymorphic), which is a common source of bugs when a base class
  "forgets" to mark a method virtual and a derived class's override is
  silently never called through a base pointer.

  ```cpp
  class Shape {
  public:
      virtual double Area() const = 0;  // pure virtual: makes Shape abstract
      virtual ~Shape() = default;        // virtual destructor: essential when
                                          // deleting derived objects through a
                                          // base pointer
  };

  class Circle : public Shape {
      double radius;
  public:
      explicit Circle(double r) : radius(r) {}
      double Area() const override { return 3.14159 * radius * radius; }
  };
  ```

- **Java**: every non-static, non-private, non-final method is virtual by
  default — there's no explicit `virtual` keyword because it's the
  default behavior, the opposite of C++. `final` disables overriding (and
  therefore the need for virtual dispatch) for a specific method or class.

  ```java
  abstract class Shape {
      abstract double area(); // implicitly "virtual" — must be overridden
  }

  class Circle extends Shape {
      private final double radius;
      Circle(double radius) { this.radius = radius; }
      @Override
      double area() { return Math.PI * radius * radius; }
  }
  ```

- **Go**: there is no virtual function mechanism at all, and no vtable
  attached to a concrete type. Instead, an **interface value** is a pair
  of (concrete type, pointer/value), and calling a method on an interface
  value looks up that concrete type's method in a table associated with
  the *interface value itself* (conceptually similar to a vtable, but
  attached to the interface value rather than the object, and built from
  structural method matching rather than explicit inheritance).

  ```go
  type Shape interface{ Area() float64 }

  type Circle struct{ Radius float64 }
  func (c Circle) Area() float64 { return 3.14159 * c.Radius * c.Radius }

  // Calling shape.Area() where shape is a Shape interface value dispatches
  // to Circle.Area at runtime — without Circle ever declaring "virtual" or
  // "implements Shape" anywhere.
  ```

**Common Misconception:** That "Go has no polymorphism because it has no
virtual functions." Go has polymorphism — it just achieves dynamic
dispatch through interfaces rather than a class-based vtable mechanism.

**Interview angle:** be ready to explain *why* C++ requires a virtual
destructor on a base class that will be deleted through a base pointer —
without it, only the base class's destructor runs (not the derived
class's), leaking any resources the derived class owns. Java and Go don't
have this footgun because neither has manual destructors tied to pointer
type.

## Abstract Classes vs. Interfaces

**Question:** How do C++, Java, and Go each express "a contract without an
implementation"?

**Detailed Explanation:**

- **C++** has no separate `interface` keyword. A class consisting entirely
  of pure virtual functions (`virtual ... = 0;`) is *used* as an interface
  by convention, but the language treats it identically to any other
  abstract class. A class can also mix pure virtual functions with real
  (concrete) methods — a genuine abstract class, not just an
  interface-in-disguise.
- **Java** has both: `interface` (originally pure method signatures only;
  since Java 8, interfaces can also have `default` methods with bodies)
  and `abstract class` (can hold state/fields and a mix of implemented and
  abstract methods). A class can implement multiple interfaces but extend
  only one abstract/concrete class — this is Java's answer to avoiding
  C++'s multiple-inheritance diamond problem while still allowing a type
  to satisfy several contracts.
- **Go** has no abstract classes at all. Interfaces are pure method-set
  contracts with zero implementation, satisfied structurally (no
  `implements` keyword needed) by any type with matching methods.

**Interview angle:** "Why did Java add default methods to interfaces in
Java 8?" — Primarily to let library authors add new methods to existing
interfaces (like `Collection`) without breaking every class that already
implements them, by supplying a default implementation instead of forcing
every implementer to add the new method immediately.

## Multiple Inheritance and the Diamond Problem

**Question:** What is the diamond problem, and how does each language
avoid or handle it?

**Short Answer:** The diamond problem arises when a class inherits from
two classes that both inherit from a common base — which version of an
inherited member does the bottom class get?

**Detailed Explanation:**

```
      Base
     /    \
  Left    Right
     \    /
    Bottom
```

- **C++** allows this directly and provides **virtual inheritance**
  (`class Left : virtual public Base`) to solve it: marking the inheritance
  from `Base` as `virtual` in both `Left` and `Right` ensures `Bottom`
  gets exactly one shared `Base` subobject instead of two separate copies.
  Without `virtual`, `Bottom` would have two independent `Base` subobjects
  and ambiguous access to `Base`'s members (`Bottom::SomeBaseMember` would
  need explicit qualification, e.g. `Left::SomeBaseMember`).
- **Java** sidesteps the problem entirely by disallowing multiple *class*
  inheritance — a class extends at most one other class. Multiple
  *interface* implementation is allowed, and Java resolves conflicting
  `default` methods from two interfaces by **requiring** the implementing
  class to override the method explicitly and choose (or combine) a
  behavior — the ambiguity becomes a compile error until resolved by hand.
- **Go** has no inheritance at all, so there's no diamond problem in the
  classical sense. Struct embedding *can* create a similar-looking
  ambiguity (embedding two structs that both have a method of the same
  name), but Go resolves it much more bluntly: calling the ambiguous
  method name directly on the outer struct is a **compile error**, and you
  must qualify which embedded field's method you mean
  (`outer.Left.SomeMethod()`).

**Common Misconception:** That Go "supports multiple inheritance" because
you can embed multiple structs. Embedding is composition with method
*promotion*, not inheritance — there's no shared base subobject, no
polymorphic dispatch through the embedding, and ambiguous promoted methods
simply fail to compile rather than being silently resolved by some
precedence rule.

## Constructors, Destructors, and RAII

**Question:** How is resource cleanup handled differently across these
three languages?

**Detailed Explanation:**

- **C++** uses **RAII** (Resource Acquisition Is Initialization):
  acquire a resource in a constructor, release it in the destructor, and
  let the language's deterministic, scope-based destruction guarantee
  cleanup — including on early returns or exceptions (stack unwinding
  calls destructors automatically). This is why idiomatic modern C++
  avoids raw `new`/`delete` in favor of `std::unique_ptr`/`std::shared_ptr`,
  which wrap RAII around dynamic allocation.

  ```cpp
  class FileHandle {
      FILE* f;
  public:
      explicit FileHandle(const char* path) : f(fopen(path, "r")) {}
      ~FileHandle() { if (f) fclose(f); }  // guaranteed to run on scope exit
  };
  ```

- **Java** has constructors but no deterministic destructors — cleanup
  relies on the garbage collector, whose timing is not guaranteed.
  `finalize()` existed for pre-cleanup hooks but is deprecated; the
  modern idiom is `try-with-resources` plus the `AutoCloseable` interface,
  which is deterministic (like RAII) but requires the caller to opt in via
  a `try (...)` block rather than getting it automatically from scope
  exit.

  ```java
  try (FileReader reader = new FileReader(path)) {
      // reader.close() is guaranteed to run at the end of this block
  }
  ```

- **Go** has no constructors or destructors as language features — object
  creation is just a struct literal or a conventional `NewX()` function,
  and cleanup uses `defer` (guaranteed to run when the enclosing function
  returns, similar in spirit to C++ RAII's scope-based guarantee, but
  explicit rather than automatic) combined with garbage collection for
  memory itself.

  ```go
  func readFile(path string) error {
      f, err := os.Open(path)
      if err != nil {
          return err
      }
      defer f.Close() // guaranteed to run when readFile returns
      // ...
      return nil
  }
  ```

**Interview angle:** "Why doesn't Java have deterministic destructors like
C++?" — Java's garbage collector doesn't track *when* the last reference
to an object disappears in a way that's cheap to act on immediately (unlike
C++'s stack-based, deterministic object lifetimes) — tying cleanup to GC
timing would make resource release unpredictable, which is exactly why
`try-with-resources` exists as an explicit, deterministic alternative for
resources (files, sockets) that can't wait for GC.

## Operator Overloading

**Question:** Why does C++ allow operator overloading but Java and Go
don't (with minor exceptions)?

**Detailed Explanation:** C++ allows redefining what operators like `+`,
`==`, `[]`, and even function-call syntax `()` mean for user-defined
types, enabling natural syntax for mathematical/container types (e.g.,
`matrix1 + matrix2`). Java deliberately omits this (aside from the
built-in overload of `+` for `String` concatenation) — the language
designers considered operator overloading a common source of confusing,
hard-to-read code in C++ (an operator silently doing something
unexpected, like `+` triggering a database call). Go's designers made the
same choice for the same reason: predictability and readability over
expressive flexibility. Go requires named methods (`a.Add(b)`) instead.

**Common Misconception:** That Go's lack of generics-era operator support
means numeric-like custom types are awkward — Go does support standard
arithmetic operators (`+`, `-`, etc.) on *named types with an underlying
numeric type* (e.g., `type Meters float64` supports `+` natively), but you
cannot define custom operator behavior for struct types the way C++ does.

## Method Overloading, Revisited Per Language

The main OOP doc covers this at a high level; here's the per-language
mechanism:

- **C++**: overload resolution happens at **compile time**, matching the
  call site's argument types/count against all visible overloads —
  ambiguous matches are a compile error.
- **Java**: same idea — compile-time overload resolution by signature.
  Combined with virtual dispatch for *overriding*, this means Java has two
  independent mechanisms that are easy to conflate: overloading (same
  name, different signature, resolved at compile time) and overriding
  (same name, same signature, resolved at runtime via the vtable).
- **Go**: neither exists. A type can only have one method with a given
  name. The idiomatic replacements are variadic parameters
  (`func Sum(nums ...int)`), accepting an interface parameter, or simply
  using distinct names (`ParseInt`, `ParseFloat` instead of overloaded
  `Parse`) — which is also why Go's standard library has many
  differently-named functions where a C++/Java library would have one
  overloaded name.

## Quick Revision

- **Virtual dispatch**: C++ explicit (`virtual` + vtable), Java implicit
  (default for non-final/non-private/non-static methods), Go has none —
  polymorphism comes from interface values instead.
- **Abstract classes**: C++ (pure virtual functions), Java (`abstract`
  keyword, can hold state), Go (not supported — use an interface).
- **Interfaces**: Java has an explicit keyword + `implements`; Go is
  fully structural (no keyword needed); C++ has no dedicated construct at
  all.
- **Multiple inheritance**: C++ allows it (virtual inheritance solves the
  diamond problem); Java disallows multiple class inheritance but allows
  multiple interfaces; Go has no inheritance, only composition.
- **Constructors/destructors**: C++ has deterministic RAII destructors;
  Java relies on GC + `try-with-resources`; Go has no destructors, uses
  `defer` + GC.
- **Operator overloading**: C++ yes, Java/Go no (by deliberate design
  choice favoring readability).
- **Method overloading**: C++/Java yes (compile-time signature
  resolution); Go no (use variadic params, interfaces, or distinct
  names).
