VERSION 0.8
FROM earthly/dind:alpine

# Variables (global so they propagate to all targets)
ARG --global iterations=3
ARG --global warmups=2
ARG --global timeas="second"

# Benchmarks and outputs results to ./results
BENCH:
  FUNCTION
  ARG --required cmd
  ARG --required lang
  ARG --required year
  ARG --required day
  ARG day_padded =$(printf "%02d" $day)
  ARG index=0
  RUN --no-cache hyperfine "$cmd" --warmup $warmups --runs $iterations --time-unit $timeas --export-json "./hyperfine.json"
  SAVE ARTIFACT ./hyperfine.json AS LOCAL ./results/$lang-$year-$day_padded.json

# Prepare debian based containers
PREPARE_DEBIAN:
  FUNCTION
  RUN apt-get update && apt-get install -y wget curl
  RUN ARCH=$(dpkg --print-architecture) && \
      wget -q https://github.com/sharkdp/hyperfine/releases/download/v1.18.0/hyperfine_1.18.0_${ARCH}.deb && \
      dpkg -i hyperfine_1.18.0_${ARCH}.deb

# Prepare alpine based containers
PREPARE_ALPINE:
  FUNCTION
  RUN apk add --no-cache hyperfine curl

# Get AoC input based on year and day
# First tries to use pre-downloaded input from CI (aoc-inputs/)
# Falls back to downloading from AOC website (local development)
GET_INPUT:
  FUNCTION
  ARG --required year
  ARG --required day
  # Check if input was pre-downloaded (CI environment)
  COPY --if-exists aoc-inputs/$year-$day.txt ./input.txt
  # Fallback to download if not found (local development)
  RUN --mount=type=cache,target=/cache/aoc \
      --secret AOC_SESSION \
      if [ ! -f ./input.txt ]; then \
        mkdir -p /cache/aoc && \
        if [ ! -f /cache/aoc/$year-$day.txt ]; then \
          curl -b "session=$AOC_SESSION" "https://adventofcode.com/$year/day/$day/input" -o /cache/aoc/$year-$day.txt || exit 1; \
        fi && \
        cp /cache/aoc/$year-$day.txt ./input.txt; \
      fi

# Languages

pypy:
  FROM pypy:3.10
  ARG --required year
  ARG --required day
  ARG day_padded =$(printf "%02d" $day)
  DO +PREPARE_DEBIAN
  DO +GET_INPUT --year=$year --day=$day
  COPY ./solutions/py/$year/day$day_padded/main.py ./
  DO +BENCH --cmd="pypy main.py input.txt" --lang="py" --year=$year --day=$day

go:
  FROM golang:1.23-alpine
  ARG --required year
  ARG --required day
  ARG day_padded =$(printf "%02d" $day)
  DO +PREPARE_ALPINE
  DO +GET_INPUT --year=$year --day=$day
  COPY ./solutions/go/$year/day$day_padded/main.go ./
  RUN --no-cache go build main.go
  DO +BENCH --cmd="./main input.txt" --lang="go" --year=$year --day=$day

bun:
  FROM oven/bun:1.2-alpine
  ARG --required year
  ARG --required day
  ARG day_padded =$(printf "%02d" $day)
  DO +PREPARE_ALPINE
  DO +GET_INPUT --year=$year --day=$day
  COPY ./solutions/ts/$year/day$day_padded/index.ts ./
  DO +BENCH --cmd="bun run index.ts input.txt" --lang="ts" --year=$year --day=$day
