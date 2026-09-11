# Interview Prep — DSA, Systems, and CS Fundamentals

**📖 Browse as a website: https://xudipta.github.io/prep/**

A structured, pattern-first repository for preparing for technical interviews and
competitive programming. Solutions are in Go, the primary and fully-tested
reference implementation; every problem also has a C++ port at
`cpp/solution.cpp` for readers who want to see (or practice in) both
languages.

This is not a dump of random problems. It is organized around **problem-solving
techniques** so that studying it builds transferable intuition — the ability to
look at a new problem and recognize *which* pattern applies — rather than
memorized answers to specific questions.

## Who this is for

- Engineers preparing for coding interviews who want pattern recognition, not
  rote memorization.
- Anyone who wants a single place to revise DSA, core CS (OS/DBMS/Networking/
  Concurrency/OOP), and system design before an interview loop.
- Competitive programmers who want concise technique notes and templates.

## Repository philosophy

1. **Learn concepts from fundamentals** — every technique has a theory note
   before any problem.
2. **Recognize patterns** — every technique document lists recognition
   signals: what in a problem statement should make you think "this is a
   two-pointers problem" or "this is interval DP."
3. **Practice representative problems** — a handful of problems that each
   teach a *distinct* idea, not dozens of near-duplicates.
4. **Understand multiple approaches** where the comparison is instructive
   (brute force → optimized), skipped where it adds no value.
5. **Revise quickly** — every technique has a one-screen "Quick Revision"
   section, and `revision/` holds cheatsheets across every topic.

## Repository structure

```
.
├── README.md                  this file
├── PROBLEMS.md                 index of every problem, with pattern + difficulty
├── PATTERN-MAP.md               "problem characteristic -> pattern" decision map
├── PROGRESS.md                  checkbox tracker for the whole curriculum
├── REFERENCES.md                sources used for conceptual material
├── templates/
│   └── problem-template.md      the canonical problem write-up template
│
├── dsa/                         one directory per technique/pattern
│   └── <technique>/
│       ├── README.md            theory: definition, intuition, template, when (not) to use
│       ├── notes.md              one-screen quick revision for the technique
│       └── problems/<slug>/
│           ├── README.md         problem write-up (uses templates/problem-template.md)
│           ├── solution.go        Go solution, package per problem directory (primary, tested)
│           ├── solution_test.go   table-driven tests for solution.go
│           └── cpp/
│               └── solution.cpp   C++ port with its own assert-based main() as its test
│
├── data-structures/             revision notes + Go snippets per data structure
│
├── computer-science/
│   ├── oop/
│   ├── operating-systems/
│   ├── databases/
│   ├── networking/
│   └── concurrency/
│
├── system-design/
│   ├── fundamentals/
│   ├── estimation/
│   ├── design-patterns/
│   └── case-studies/
│
├── tips-tricks/                  checklists and problem-solving frameworks
├── revision/                     one cheatsheet per subject, single-screen
│
├── playground/, problems/, utils/, Makefile
│                                 the original Go scratch workspace (unrelated
│                                 to the curriculum above — see below)
├── scripts/                     repo tooling: new_problem.sh (scratch-workspace
│                                 helper) and verify-cpp-solutions.sh (compiles +
│                                 runs every dsa/*/problems/*/cpp/solution.cpp;
│                                 used by CI)
└── .github/workflows/            CI + GitHub Pages deployment
```

> Note on `problems/` vs `dsa/*/problems/`: the top-level `problems/` and
> `playground/` directories predate this curriculum and remain as a free-form
> scratch workspace for quick experiments (see the bottom of this file). All
> curated, documented interview-prep content lives under `dsa/`,
> `data-structures/`, `computer-science/`, and `system-design/`.

## Learning path

Follow this order if you're starting from scratch; jump around freely if
you're revising. This is also the order every technique appears in the
sidebar, `PROBLEMS.md`, and `PROGRESS.md`, so you can move between them
without losing your place. The sidebar groups these 14 techniques into 5
collapsible clusters so the nav doesn't show all of them at once — each
cluster below is a contiguous stretch of this same order, not a
reordering:

- **Arrays & Strings**: Hashing, Two Pointers, Sliding Window, Prefix Sum
- **Linear Structures**: Stack, Linked List
- **Search & Recursion**: Binary Search, Backtracking
- **Trees & Graphs**: Trees, Graphs, Union-Find
- **Optimization**: Bit Manipulation, Greedy, Dynamic Programming

```
Programming Fundamentals
        ↓
Complexity Analysis            → revision/algorithms-cheatsheet.md
        ↓
Basic Data Structures          → data-structures/
        ↓
Hashing                        → dsa/hashing            (arrays/strings base toolkit)
        ↓
Two Pointers                   → dsa/two-pointers
        ↓
Sliding Window                 → dsa/sliding-window
        ↓
Prefix Sum                     → dsa/prefix-sum
        ↓
Stack / Monotonic Stack        → dsa/stack
        ↓
Linked List                    → dsa/linked-list
        ↓
Binary Search                  → dsa/binary-search
        ↓
Backtracking                   → dsa/backtracking       (recursion)
        ↓
Trees                          → dsa/trees
        ↓
Graphs                         → dsa/graphs
        ↓
Union-Find                     → dsa/union-find          (graph connectivity, DSU)
        ↓
Bit Manipulation               → dsa/bit-manipulation
        ↓
Greedy                         → dsa/greedy
        ↓
Dynamic Programming            → dsa/dynamic-programming (capstone: recursion + memoization)
        ↓
CS Fundamentals                → computer-science/
        ↓
System Design                  → system-design/
```

**Why this order:** each stage builds a prerequisite for the next. Hashing,
two pointers, sliding window, and prefix sum are the base array/string
toolkit. Stack and linked list are the next-simplest linear-structure
techniques. Binary search and backtracking (recursion) are the mental
building blocks for trees, then graphs, which generalize trees. Union-Find
follows graphs since its two problems here are graph-connectivity problems
(cycle detection, component counting) — the same shape as the graph
traversal problems just solved, but with a different data structure. Bit
manipulation is a mostly independent toolkit slotted in once the core
patterns are solid. Greedy comes right before DP because they're often two
answers to the same question ("is there a locally-optimal choice that's
always safe?") — and DP is placed last because most DP transitions are best
understood as "recursion on a smaller subproblem, cached," so it builds on
everything above it. CS fundamentals and system design come last because
interview loops typically test them in later rounds, and system design draws
on data structure and scaling intuition built earlier.

## Difficulty levels

Every problem is tagged Easy / Medium / Hard, and every technique aims for
roughly:

- 20–30% foundational (Beginner/Easy) problems
- 50–60% Medium / standard-interview problems
- 15–25% Hard / advanced problems

We do not pad techniques with near-duplicate problems just to hit a count.
Every problem in this repo should teach something the others don't — see
`PROBLEMS.md` for the full index and what each problem teaches.

## How to use the revision notes

- **Level 1 — Quick Revision**: `revision/*.md` and each technique's
  `notes.md`. Read these the night before / morning of an interview.
- **Level 2 — Concept Revision**: each technique's `README.md` and each
  `computer-science/<topic>/README.md`. Definitions, patterns, examples,
  common interview questions.
- **Level 3 — Deep Dive**: the individual problem write-ups under
  `dsa/*/problems/*/README.md`, and `system-design/case-studies/*.md`.

## Tracking progress

`PROGRESS.md` has checkboxes for every technique, CS topic, and system-design
area. Check items off as you complete them. It doubles as a table of contents
for what exists today versus what's still planned.

## Running the Go code

```bash
go build ./...
go vet ./...
go test ./...
```

Each problem lives in its own package (its directory), so tests can be run
per-problem too:

```bash
go test ./dsa/two-pointers/problems/three-sum/...
```

## Running the C++ ports

Each `cpp/solution.cpp` is self-contained, with its own `assert`-based
`main()` acting as its test suite — compile and run it directly:

```bash
g++ -std=c++17 -Wall -Wextra -O1 dsa/two-pointers/problems/three-sum/cpp/solution.cpp -o /tmp/three-sum
/tmp/three-sum
```

To check every C++ port at once (the same check CI runs):

```bash
bash scripts/verify-cpp-solutions.sh
```

## Browsing as a website

This repository is also published as a browsable site via GitHub Pages:
**https://xudipta.github.io/prep/** (docsify, no build step — see
`.github/workflows/pages.yml`). It redeploys automatically on every push to
`main`.

If the link 404s, Pages hasn't been enabled on the repo yet — this is a
one-time manual step (`Settings → Pages → Build and deployment → Source:
GitHub Actions`) that a GitHub Actions workflow can't do on its own, since
creating a Pages site requires repository-admin permissions that the
workflow's token doesn't have. Once it's enabled once, every future push
deploys automatically with no further action needed.

## Contributing a new problem or note

1. Pick the technique directory it belongs to (or propose a new one in
   `dsa/` if it's a genuinely distinct pattern).
2. Ask: *"What new idea does this problem teach that isn't already covered?"*
   If the answer is "nothing new," it's a duplicate — don't add it.
3. Copy `templates/problem-template.md` into
   `dsa/<technique>/problems/<slug>/README.md` and fill it in.
4. Add `solution.go` (package name = the slug in `snake_case` or a short
   camelCase package name) and a table-driven `solution_test.go`.
5. Optionally add a C++ port at `cpp/solution.cpp`, faithful to the same
   algorithm and complexity, with its own `assert`-based `main()` in place
   of a separate test file.
6. Add a row to `PROBLEMS.md` and, if relevant, an entry in `PATTERN-MAP.md`.
7. Update `PROGRESS.md`.
8. Run `go build ./... && go vet ./... && go test ./...`, and
   `bash scripts/verify-cpp-solutions.sh` if you added a C++ port, before
   opening a PR.

## Source quality

Conceptual material (OS/DBMS/Networking/System Design/OOP) is written in
original language, cross-checked against multiple reputable sources, and not
copied verbatim from any single source. See `REFERENCES.md` for what was
consulted.
