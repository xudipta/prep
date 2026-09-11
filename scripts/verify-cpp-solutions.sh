#!/usr/bin/env bash
# Compiles and runs every C++ solution port under dsa/*/problems/*/cpp/solution.cpp,
# failing if any file doesn't compile warning-free or doesn't print exactly
# "All tests passed" (each file's own assert-based main() is its test suite).
set -euo pipefail

root_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$root_dir"

fail=0
count=0

for f in $(find dsa -path '*/cpp/solution.cpp' | sort); do
  count=$((count + 1))
  bin="$(mktemp)"
  if ! g++ -std=c++17 -Wall -Wextra -Werror -O1 -o "$bin" "$f" 2>&1; then
    echo "COMPILE FAILED: $f"
    fail=$((fail + 1))
    rm -f "$bin"
    continue
  fi

  out="$("$bin" 2>&1)" || {
    echo "RUNTIME FAILURE (assertion or crash): $f"
    echo "$out"
    fail=$((fail + 1))
    rm -f "$bin"
    continue
  }

  if [ "$out" != "All tests passed" ]; then
    echo "UNEXPECTED OUTPUT: $f -> $out"
    fail=$((fail + 1))
  fi
  rm -f "$bin"
done

echo "Checked $count C++ solution files, $fail failure(s)."
if [ "$fail" -ne 0 ]; then
  exit 1
fi
