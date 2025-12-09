import { readFileSync } from "fs";

export function partOne(input: string): number {
	const { left, right } = input
		.split("\n")
		.filter((line) => line.trim())
		.reduce((acc, line) => {
			const [l, r] = line.split("   ").map(Number)
			acc.left.push(l!)
			acc.right.push(r!)
			return acc
		}, { left: [] as number[], right: [] as number[] })

	left.sort((a, b) => a - b)
	right.sort((a, b) => a - b)

	return left.reduce((sum, num, i) => sum + Math.abs(num - right[i]!), 0)
}

export function partTwo(input: string): number {
	const { left, right } = input
		.split("\n")
		.filter((line) => line.trim())
		.reduce((acc, line) => {
			const [l, r] = line.split("   ").map(Number)
			acc.left.push(l!)
			acc.right.set(r!, (acc.right.get(r!) ?? 0) + 1)
			return acc
		}, { left: [] as number[], right: new Map<number, number>() })

	return left.reduce((sum, num) => {
		const count = right.get(num) ?? 0
		return sum + num * count
	}, 0)
}

if (import.meta.main) {
	if (process.argv.length < 3) {
		console.error("Usage: bun index.ts <filepath>");
		process.exit(1);
	}

	const filepath = process.argv[2];
	const input = readFileSync(filepath, "utf-8");

	const result1 = partOne(input);
	const result2 = partTwo(input);

	console.log(`Part 1: ${result1}`);
	console.log(`Part 2: ${result2}`);
}
