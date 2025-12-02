#!/usr/bin/env bash
set -euo pipefail

# Advent of Code README Star Counter Updater
# Updates README.md with current star progress from adventofcode.com/events

# Get AOC_SESSION from argument or environment
AOC_SESSION="${1:-${AOC_SESSION:-}}"

if [ -z "$AOC_SESSION" ]; then
  echo "Error: AOC_SESSION is required"
  echo "Usage: $0 <session_cookie>"
  echo "   or: AOC_SESSION=<cookie> $0"
  exit 1
fi

README="README.md"

# Check if README exists
if [ ! -f "$README" ]; then
  echo "Error: $README not found"
  exit 1
fi

echo "Fetching star progress from adventofcode.com/events..."

# Fetch events page
HTML=$(curl -s -H "Cookie: session=${AOC_SESSION}" https://adventofcode.com/events)

# Parse HTML to extract years and stars
# The HTML format is:
# <div class="eventlist-event"><a href="/2024">[2024]</a> <span class="star-count">11*</span> <span class="quiet">/ 50*</span></div>
# or for years with no stars:
# <div class="eventlist-event"><a href="/2020">[2020]</a>          </div>

echo "Parsing star counts..."

YEARS_DATA=$(echo "$HTML" | grep 'class="eventlist-event"' | while read -r line; do
  # Extract year
  year=$(echo "$line" | sed -n 's/.*\[\([0-9]\{4\}\)\].*/\1/p')

  # Extract star count (if present)
  if echo "$line" | grep -q 'star-count'; then
    # Remove leading spaces and extract number
    stars=$(echo "$line" | sed -n 's/.*<span class="star-count">[[:space:]]*\([0-9]\+\)\*.*/\1/p')
    # Extract total stars from the "/ XX*" pattern
    total=$(echo "$line" | sed -n 's/.*<span class="quiet">[[:space:]]*\/[[:space:]]*\([0-9]\+\)\*.*/\1/p')
  else
    stars=0
    total=50
  fi

  # Only include years with stars > 0
  if [ "$stars" -gt 0 ] && [ -n "$year" ]; then
    echo "- $year: $stars/$total ⭐"
  fi
done | sort -r)

if [ -z "$YEARS_DATA" ]; then
  echo "Warning: No star data found. Check your session cookie."
  exit 1
fi

echo "Found stars for the following years:"
echo "$YEARS_DATA"

# Update README.md
# Strategy: Replace everything between "## My Progress" and the next "##" section
echo "Updating $README..."

# Create temporary file with updated content
awk -v years="$YEARS_DATA" '
  # Print everything before "## My Progress"
  /^## My Progress/ {
    in_progress = 1
    print
    print ""
    print "Overall Advent of Code progress (solutions may be in this repository or elsewhere):"
    print ""
    print years
    next
  }

  # When we hit the next section, stop skipping
  /^## / && in_progress {
    in_progress = 0
  }

  # Skip lines while in the progress section
  in_progress {
    next
  }

  # Print everything else
  {
    print
  }
' "$README" > "${README}.tmp"

# Replace original file
mv "${README}.tmp" "$README"

echo "Successfully updated $README with current star progress!"
