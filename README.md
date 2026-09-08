# Interview Prep — DSA, Systems, and CS Fundamentals

A structured, pattern-first repository for preparing for technical interviews and
competitive programming. Solutions are in Go.

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
│           ├── solution.go        Go solution, package per problem directory
│           └── solution_test.go   table-driven tests
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
├── playground/, problems/, utils/, scripts/, Makefile
│                                 the original Go scratch workspace (unrelated
│                                 to the curriculum above — see below)
└── .github/workflows/            CI + GitHub Pages deployment
```

> Note on `problems/` vs `dsa/*/problems/`: the top-level `problems/` and
> `playground/` directories predate this curriculum and remain as a free-form
> scratch workspace for quick experiments (see the bottom of this file). All
> curated, documented interview-prep content lives under `dsa/`,
> `data-structures/`, `computer-science/`, and `system-design/`.

## Learning path

Follow this order if you're starting from scratch; jump around freely if
you're revising:

```
Programming Fundamentals
        ↓
Complexity Analysis            → revision/algorithms-cheatsheet.md
        ↓
Basic Data Structures          → data-structures/
        ↓
Arrays / Strings / Hashing     → dsa/hashing
        ↓
Two Pointers / Sliding Window  → dsa/two-pointers, dsa/sliding-window
        ↓
Stack / Queue / Heap           → data-structures/
        ↓
Binary Search                  → dsa/binary-search
        ↓
Recursion / Backtracking       → dsa/backtracking
        ↓
Trees                          → dsa/trees
        ↓
Graphs                         → dsa/graphs
        ↓
Greedy / Divide & Conquer
        ↓
Dynamic Programming            → dsa/dynamic-programming
        ↓
Advanced Data Structures       → Trie, DSU, Segment Tree, Fenwick Tree
        ↓
CS Fundamentals                → computer-science/
        ↓
System Design                  → system-design/
```

**Why this order:** each stage builds a prerequisite for the next. Hashing and
two pointers give you the base toolkit for array/string problems. Binary
search and recursion are the mental building blocks for backtracking, trees,
and eventually DP (which is recursion + memoization). Graphs generalize
trees. DP is placed after recursion/trees/graphs because most DP transitions
are best understood as "recursion on a smaller subproblem, cached." CS
fundamentals and system design come last because interview loops typically
test them in later rounds, and system design draws on data structure and
scaling intuition built earlier.

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

## Browsing as a website

This repository is also published as a browsable site via GitHub Pages
(docsify, no build step — see `.github/workflows/pages.yml`). If Pages isn't
yet enabled for this repo: **Settings → Pages → Build and deployment → Source:
GitHub Actions**. Once enabled, the workflow publishes on every push to
`main`.

## Contributing a new problem or note

1. Pick the technique directory it belongs to (or propose a new one in
   `dsa/` if it's a genuinely distinct pattern).
2. Ask: *"What new idea does this problem teach that isn't already covered?"*
   If the answer is "nothing new," it's a duplicate — don't add it.
3. Copy `templates/problem-template.md` into
   `dsa/<technique>/problems/<slug>/README.md` and fill it in.
4. Add `solution.go` (package name = the slug in `snake_case` or a short
   camelCase package name) and a table-driven `solution_test.go`.
5. Add a row to `PROBLEMS.md` and, if relevant, an entry in `PATTERN-MAP.md`.
6. Update `PROGRESS.md`.
7. Run `go build ./... && go vet ./... && go test ./...` before opening a PR.

## Source quality

Conceptual material (OS/DBMS/Networking/System Design/OOP) is written in
original language, cross-checked against multiple reputable sources, and not
copied verbatim from any single source. See `REFERENCES.md` for what was
consulted.
