#!/usr/bin/env python3
"""
Update README.md with benchmark results from JSON artifacts.

Usage:
    python update_benchmarks.py /path/to/benchmarks/directory/
"""

import argparse
import json
import re
from pathlib import Path
from collections import defaultdict
from typing import Dict, Tuple, List


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


def format_language_list(languages: List[str], fastest: str) -> str:
    """
    Format language list with fastest in bold.

    Args:
        languages: List of language names
        fastest: Name of fastest language

    Returns:
        Formatted string like "**Go**, PyPy, Bun"
    """
    sorted_langs = sorted(languages)
    formatted = []

    for lang in sorted_langs:
        if lang == fastest:
            formatted.append(f"**{lang}**")
        else:
            formatted.append(lang)

    return ", ".join(formatted)


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
        lang_list = format_language_list(languages, fastest)

        # Format with proper spacing to align with column headers
        time_str = f"{min_time:.3f}".ljust(13)
        rows.append(f"| {year} | {time_str} | xxx   | {lang_list.ljust(25)} |")

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

        # Format language list with fastest in bold
        lang_list = format_language_list(list(languages.keys()), fastest_lang)

        day_str = f"{day:02d}"
        # Format with proper spacing to align with column headers
        time_str = f"{min_time:.3f}".ljust(13)
        rows.append(f"| {day_str}  | {time_str} | xxx   | {lang_list.ljust(25)} |")

    return "\n".join(rows)


def update_readme(
    benchmarks: Dict[int, Dict[int, Dict[str, float]]], readme_path: Path
):
    """
    Update README.md with benchmark results.

    Replaces content between '# Benchmark' heading and '---' line with:
    - Overall stats table for all years
    - Individual tables for each year

    Args:
        benchmarks: Nested dict from load_benchmarks()
        readme_path: Path to README.md file
    """
    # Read current README
    with open(readme_path, "r") as f:
        content = f.read()

    # Calculate year totals
    year_totals = calculate_year_totals(benchmarks)

    # Build the new benchmark section content
    sections = []

    # Overall table
    sections.append("")
    sections.append("## Overall")
    sections.append("")
    sections.append("| Year | Min (seconds) | Stars | Language                  |")
    sections.append("|------|---------------|-------|---------------------------|")
    sections.append(generate_overall_table(benchmarks, year_totals))
    sections.append("")

    # Individual year tables
    for year in sorted(benchmarks.keys(), reverse=True):
        sections.append(f"## {year}")
        sections.append("")
        sections.append("| Day | Min (seconds) | Stars | Language                  |")
        sections.append("|-----|---------------|-------|---------------------------|")
        sections.append(generate_year_table(year, benchmarks[year]))
        sections.append("")

    new_benchmark_content = "\n".join(sections)

    # Find and replace content between # Benchmarks and ---
    # Pattern matches from "# Benchmarks" to the line before "---"
    pattern = r"(# Benchmarks\n)(.*?)(\n---)"

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

    args = parser.parse_args()

    try:
        benchmarks = load_benchmarks(args.directory)
        update_readme(benchmarks, args.readme)
    except Exception as e:
        print(f"Error: {e}")
        return 1

    return 0


if __name__ == "__main__":
    exit(main())
