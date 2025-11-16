#!/usr/bin/env bash
set -euo pipefail

# Advent of Code Puzzle Input Fetcher
# Fetches puzzle inputs for all days with solution files

# Immediate failure if session cookie is missing
if [ -z "$AOC_SESSION" ]; then
  echo "::error::AOC_SESSION secret is missing. Add it to GitHub Secrets."
  exit 1
fi

# Initialize error tracking
HAS_ERRORS=0

# Enhanced fetch function with retries
fetch_input() {
  local year=$1
  local day=$2
  local day_padded=$(printf "%02d" $day)
  local solution_dir="go/${year}/day${day_padded}"
  local data_dir="data/${year}"
  local data_file="${data_dir}/day${day_padded}.txt"
  local max_retries=3
  local attempt=1
  local retry_delay=1

  # Skip if directory structure is invalid or solution doesn't exist
  if [ ! -d "$solution_dir" ] || [ ! -f "${solution_dir}/solution.go" ]; then
    return 0
  fi

  # Create data directory if it doesn't exist
  mkdir -p "$data_dir"

  echo "Fetching input for ${year} day ${day} (${data_file})"

  while [ $attempt -le $max_retries ]; do
    # Use temporary file to capture response
    local tmpfile=$(mktemp)
    local status_code=$(curl -s -w "%{http_code}" \
      -H "Cookie: session=${AOC_SESSION}" \
      "https://adventofcode.com/${year}/day/${day}/input" \
      -o "$tmpfile" \
      --silent --show-error)

    if [ $status_code -eq 200 ]; then
      mv "$tmpfile" "$data_file"
      echo "Successfully fetched input for ${year} day ${day}"
      sleep 1  # Maintain server-friendly delay
      return 0
    else
      echo "Attempt ${attempt} failed: HTTP ${status_code}"
      rm -f "$tmpfile"
      sleep $retry_delay
      retry_delay=$((retry_delay * 2))
      attempt=$((attempt + 1))
    fi
  done

  echo "::error::Failed to fetch input for ${year} day ${day} after ${max_retries} attempts"
  HAS_ERRORS=1
  return 1
}

# Process year directories in go/
for dir in go/*; do
  if [ -d "$dir" ]; then
    year=$(basename "$dir")
    if [[ "$year" =~ ^20[0-9]{2}$ ]]; then
      echo "Processing year directory: $year"
      for day in {1..25}; do
        fetch_input "$year" "$day"
      done
    fi
  fi
done

# Final failure if any input failed to fetch
if [ $HAS_ERRORS -ne 0 ]; then
  echo "::error::Some inputs failed to fetch after retries"
  exit 1
fi
