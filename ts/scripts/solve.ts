/**
 * Solves an Advent of Code puzzle by fetching input and running the solution
 * Usage: bun run scripts/solve.ts <year> <day>
 * Example: bun run scripts/solve.ts 2024 1
 */

import { fetchInput } from "./get-aoc"

async function solve(year: number, day: number) {
	console.log(`🎄 Solving Advent of Code ${year} Day ${day}...\n`);

	// Fetch the input
	console.log("📥 Fetching puzzle input...");
	const input = await fetchInput(year, day);
	console.log(`✓ Input fetched (${input.length} characters)\n`);

	// Construct the path to the solution file
	const dayPadded = day.toString().padStart(2, "0");
	const solutionPath = `../${year}/day${dayPadded}/solution.ts`;
	const solutionUrl = new URL(solutionPath, import.meta.url).pathname;

	// Import the solution
	console.log(`📂 Loading solution from ${year}/day${dayPadded}/solution.ts...`);

	let solution;
	try {
		solution = await import(solutionUrl);
	} catch (error) {
		throw new Error(
			`Failed to load solution file. Make sure ${year}/day${dayPadded}/solution.ts exists.`
		);
	}

	// Run part 1
	if (solution.partOne) {
		console.log("\n🌟 Running Part 1...");
		const start1 = performance.now();
		const result1 = solution.partOne(input);
		const end1 = performance.now();
		console.log(`Answer: ${result1}`);
		console.log(`Time: ${(end1 - start1).toFixed(2)}ms`);
	} else {
		console.log("\n⚠️  Part 1 function not found");
	}

	// Run part 2
	if (solution.partTwo) {
		console.log("\n🌟🌟 Running Part 2...");
		const start2 = performance.now();
		const result2 = solution.partTwo(input);
		const end2 = performance.now();
		console.log(`Answer: ${result2}`);
		console.log(`Time: ${(end2 - start2).toFixed(2)}ms`);
	} else {
		console.log("\n⚠️  Part 2 function not found");
	}

	console.log("\n✨ Done!");
}

// CLI usage
if (import.meta.main) {
	const args = process.argv.slice(2);

	if (args.length < 2) {
		console.error("Usage: bun run scripts/solve.ts <year> <day>");
		console.error("Example: bun run scripts/solve.ts 2024 1");
		process.exit(1);
	}

	const year = parseInt(args[0]!);
	const day = parseInt(args[1]!);

	if (isNaN(year) || isNaN(day)) {
		console.error("Year and day must be numbers");
		process.exit(1);
	}

	try {
		await solve(year, day);
	} catch (error) {
		console.error("\n❌ Error:", error instanceof Error ? error.message : error);
		process.exit(1);
	}
}
