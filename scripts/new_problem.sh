#!/bin/sh
set -e

NAME="$1"
if [ -z "$NAME" ]; then
  echo "Usage: new_problem.sh name" >&2
  exit 1
fi

FILE="problems/$NAME.go"
if [ -e "$FILE" ]; then
  echo "$FILE already exists" >&2
  exit 1
fi

# Sanitize to a valid Go identifier: replace non-alnum with _
SAFE=$(echo "$NAME" | sed 's/[^A-Za-z0-9]/_/g')

cat > "$FILE" <<EOF
package problems

// $NAME is a stub for the problem named $NAME
// Add a description and implement the solution below.

// $SAFE TODO: update signature to match the problem requirements
func $SAFE(input ...interface{}) interface{} {
	// TODO: implement
	return nil
}
EOF

echo "Created $FILE (function: $SAFE)"
