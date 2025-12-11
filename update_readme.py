#!/usr/bin/env python3
"""
Update README.md with benchmark results from JSON artifacts.

Usage:
    python update_benchmarks.py /path/to/benchmarks/directory/
"""

import argparse
import json
import os
import re
from pathlib import Path
from collections import defaultdict
from typing import Dict, Tuple, List, Optional
from urllib.request import Request, urlopen
from urllib.error import URLError, HTTPError


def parse_filename(filename: str) -> Tuple[str, int, int]:
    """
    Extract language, year, and day from filename.

    Example: "py-2025-01.json" -> ("py", 2025, 1)

    Args:
        filename: Benchmark JSON filename

    Returns:
        Tuple of (language, year, day)
    """
    match = re.match(r"(\w+)-(\d{4})-(\d{2})\.json", filename)
    if not match:
        raise ValueError(f"Invalid filename format: {filename}")

    language, year_str, day_str = match.groups()
    year = int(year_str)
    day = int(day_str)

    return language, year, day


def load_benchmarks(directory_path: Path) -> Dict[int, Dict[int, Dict[str, float]]]:
    """
    Load all benchmark JSON files from directory.

    Args:
        directory_path: Path to directory containing JSON files

    Returns:
        Nested dict: {year: {day: {language: min_time}}}
    """
    benchmarks = defaultdict(lambda: defaultdict(dict))

    if not directory_path.exists():
        raise FileNotFoundError(f"Directory not found: {directory_path}")

    json_files = list(directory_path.glob("*.json"))
    if not json_files:
        raise ValueError(f"No JSON files found in {directory_path}")

    for json_file in json_files:
        try:
            language, year, day = parse_filename(json_file.name)

            with open(json_file, "r") as f:
                data = json.load(f)

            min_time = data["results"][0]["min"]
            benchmarks[year][day][language] = min_time

        except (ValueError, KeyError, IndexError) as e:
            print(f"Warning: Skipping {json_file.name}: {e}")
            continue

    return benchmarks


def calculate_year_totals(benchmarks: Dict) -> Dict[int, Dict[str, float]]:
    """
    Calculate total time per language for each year.

    Args:
        benchmarks: Nested dict from load_benchmarks()

    Returns:
        Dict: {year: {language: total_time}}
    """
    year_totals = {}

    for year, days in benchmarks.items():
        totals = defaultdict(float)

        for day, languages in days.items():
            for language, min_time in languages.items():
                totals[language] += min_time

        year_totals[year] = dict(totals)

    return year_totals


def find_fastest_language(year_totals: Dict[str, float]) -> str:
    """
    Find the language with the lowest total time.

    Args:
        year_totals: Dict of {language: total_time}

    Returns:
        Name of fastest language
    """
    return min(year_totals.items(), key=lambda x: x[1])[0]


def format_language_list(
    languages: List[str], fastest: str, year: int, day: Optional[int] = None
) -> str:
    """
    Format language list with fastest in bold and links to solution files.

    Args:
        languages: List of language names
        fastest: Name of fastest language
        year: Year number
        day: Day number (optional, None for overall table)

    Returns:
        Formatted string like "[**Go**](path), [PyPy](path), [Bun](path)"
    """
    # Map language names to their file names
    lang_file_map = {"py": "main.py", "go": "main.go", "ts": "index.ts"}

    sorted_langs = sorted(languages)
    formatted = []

    for lang in sorted_langs:
        # Construct relative path to solution
        if day is not None:
            # For specific day: solutions/py/2025/day01/main.py
            day_str = f"{day:02d}"
            file_name = lang_file_map.get(lang, "main.py")
            path = f"solutions/{lang}/{year}/day{day_str}/{file_name}"
        else:
            # For overall table, link to year directory
            path = f"solutions/{lang}/{year}"

        # Format with bold if fastest
        if lang == fastest:
            formatted.append(f"[**{lang}**]({path})")
        else:
            formatted.append(f"[{lang}]({path})")

    return ", ".join(formatted)


def fetch_star_counts(session_token: str) -> Dict[int, int]:
    """
    Fetch star counts from adventofcode.com/events.

    Args:
        session_token: AOC session cookie value

    Returns:
        Dict mapping year to star count: {2025: 12, 2024: 11, ...}
        Returns empty dict on failure
    """
    if not session_token:
        print("Warning: No AOC session token provided, skipping star count update")
        return {}

    try:
        url = "https://adventofcode.com/events"
        request = Request(url)
        request.add_header("Cookie", f"session={session_token}")
        request.add_header(
            "User-Agent", "github.com/yourusername/aoc via python/update_readme.py"
        )

        with urlopen(request, timeout=10) as response:
            html = response.read().decode("utf-8")

        # Parse star counts from events page
        # Pattern looks for: [YYYY] followed by star count
        # The events page shows something like: "[2025] 12*" or similar
        star_counts = {}

        # Try to find patterns like: /YYYY">...N*
        # This regex captures year and stars from links
        pattern = r'/(\d{4})">.*?(\d+)\s*\*'
        matches = re.findall(pattern, html, re.DOTALL)

        for year_str, stars_str in matches:
            year = int(year_str)
            stars = int(stars_str)
            # Only include years from 2015 onwards (when AoC started)
            if 2015 <= year <= 2030:
                star_counts[year] = stars

        if star_counts:
            print(f"Fetched star counts for {len(star_counts)} years")
        else:
            print(
                "Warning: No star counts found in response, pattern may need updating"
            )

        return star_counts

    except HTTPError as e:
        print(f"HTTP Error fetching star counts: {e.code} - {e.reason}")
        if e.code == 401:
            print("Invalid session token - please check your AOC_SESSION value")
        return {}
    except URLError as e:
        print(f"Network error fetching star counts: {e.reason}")
        return {}
    except Exception as e:
        print(f"Unexpected error fetching star counts: {e}")
        return {}


def update_progress_section(content: str, star_counts: Dict[int, int]) -> str:
    """
    Update the "My Progress" section in README with star counts.

    Args:
        content: Current README content
        star_counts: Dict mapping year to star count

    Returns:
        Updated README content
    """
    if not star_counts:
        return content

    # Pattern to match the My Progress section
    pattern = r"(## My Progress\n\nOverall Advent of Code progress.*?\n\n)((?:- \d{4}: \d+/\d+ ⭐\n)+)"
    match = re.search(pattern, content, re.DOTALL)

    if not match:
        print("Warning: Could not find 'My Progress' section in README")
        return content

    header = match.group(1)
    progress_lines = match.group(2)

    # Parse existing year lines
    year_pattern = r"- (\d{4}): \d+/(\d+) ⭐"
    year_lines = []

    for line in progress_lines.strip().split("\n"):
        year_match = re.match(year_pattern, line)
        if year_match:
            year = int(year_match.group(1))
            max_stars = year_match.group(2)
            # Use fetched count if available, otherwise keep existing
            stars = star_counts.get(year)
            if stars is not None:
                year_lines.append(f"- {year}: {stars}/{max_stars} ⭐")
            else:
                year_lines.append(line)

    # Reconstruct the section
    new_progress = header + "\n".join(year_lines) + "\n"

    # Replace in content
    new_content = re.sub(pattern, new_progress, content, flags=re.DOTALL)

    return new_content


def generate_overall_table(benchmarks: Dict, year_totals: Dict) -> str:
    """
    Generate markdown table rows for the Overall section.

    Args:
        benchmarks: Benchmark data
        year_totals: Year totals per language

    Returns:
        Markdown table rows (without header)
    """
    rows = []

    for year in sorted(year_totals.keys(), reverse=True):
        totals = year_totals[year]
        fastest = find_fastest_language(totals)
        languages = list(totals.keys())

        min_time = totals[fastest]
        lang_list = format_language_list(languages, fastest, year)

        # Format with proper spacing to align with column headers
        time_str = f"{min_time:.3f}"
        year_link = f"[{year}](https://adventofcode.com/{year})"
        rows.append(f"| {year_link} | {time_str} | {lang_list} |")

    return "\n".join(rows)


def generate_year_table(year: int, days: Dict[int, Dict[str, float]]) -> str:
    """
    Generate markdown table rows for a specific year.

    Args:
        year: Year number
        days: Dict of {day: {language: min_time}}

    Returns:
        Markdown table rows (without header)
    """
    rows = []

    for day in sorted(days.keys()):
        languages = days[day]

        # Find fastest language for this day
        fastest_lang = min(languages.items(), key=lambda x: x[1])[0]
        min_time = languages[fastest_lang]

        # Format language list with fastest in bold and links
        lang_list = format_language_list(
            list(languages.keys()), fastest_lang, year, day
        )

        day_str = f"{day:02d}"
        # Format with proper spacing to align with column headers
        time_str = f"{min_time:.3f}"
        day_link = f"[{day_str}](https://adventofcode.com/{year}/day/{day})"
        rows.append(f"| {day_link} | {time_str} | {lang_list} |")

    return "\n".join(rows)


def update_readme(
    benchmarks: Dict[int, Dict[int, Dict[str, float]]],
    readme_path: Path,
    session_token: Optional[str] = None,
):
    """
    Update README.md with benchmark results and star counts.

    Replaces content between '# Benchmark' heading and '---' line with:
    - Overall stats table for all years
    - Individual tables for each year

    Optionally updates "My Progress" section with star counts from AoC.

    Args:
        benchmarks: Nested dict from load_benchmarks()
        readme_path: Path to README.md file
        session_token: Optional AOC session token for fetching star counts
    """
    # Read current README
    with open(readme_path, "r") as f:
        content = f.read()

    # Fetch and update star counts if session token provided
    if session_token:
        star_counts = fetch_star_counts(session_token)
        content = update_progress_section(content, star_counts)

    # Calculate year totals
    year_totals = calculate_year_totals(benchmarks)

    # Build the new benchmark section content
    sections = []

    # Overall table
    sections.append("")
    sections.append("### Overall")
    sections.append("")
    sections.append("| Year | Min (seconds) | Language |")
    sections.append("|------|---------------|----------|")
    sections.append(generate_overall_table(benchmarks, year_totals))
    sections.append("")

    # Individual year tables
    for year in sorted(benchmarks.keys(), reverse=True):
        sections.append(f"### {year}")
        sections.append("")
        sections.append("| Day | Min (seconds) | Language |")
        sections.append("|-----|---------------|----------|")
        sections.append(generate_year_table(year, benchmarks[year]))
        sections.append("")

    new_benchmark_content = "\n".join(sections)

    # Find and replace content between # Benchmarks and ---
    # Pattern matches from "# Benchmarks" to the line before "---"
    pattern = r"(## Benchmarks\n)(.*?)(\n---)"

    replacement = r"\1" + new_benchmark_content + r"\3"

    updated_content = re.sub(pattern, replacement, content, flags=re.DOTALL)

    # Write updated README
    with open(readme_path, "w") as f:
        f.write(updated_content)

    print(f"Successfully updated {readme_path}")


def main():
    """Main entry point."""
    parser = argparse.ArgumentParser(
        description="Update README.md with benchmark results from JSON artifacts"
    )
    parser.add_argument(
        "directory", type=Path, help="Directory containing benchmark JSON files"
    )
    parser.add_argument(
        "--readme",
        type=Path,
        default=Path("README.md"),
        help="Path to README.md (default: ./README.md)",
    )
    parser.add_argument(
        "--session",
        type=str,
        help="AOC session token for fetching star counts (defaults to AOC_SESSION env var)",
    )

    args = parser.parse_args()

    # Get session token from args or environment variable
    session_token = args.session or os.environ.get("AOC_SESSION")

    try:
        benchmarks = load_benchmarks(args.directory)
        update_readme(benchmarks, args.readme, session_token)
    except Exception as e:
        print(f"Error: {e}")
        return 1

    return 0


if __name__ == "__main__":
    exit(main())
