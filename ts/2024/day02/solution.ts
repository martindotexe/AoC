export function partOne(input: string): number {
	return input
		.split("\n")
		.filter((line) => line.trim())
		.reduce((count, line) => {
			const levels = line.split(" ").map(Number)
			const diffs = levels.slice(1).map((num, i) => num - levels[i]!)

			const allIncreasing = diffs.every(d => d >= 1 && d <= 3)
			const allDecreasing = diffs.every(d => d >= -3 && d <= -1)

			return count + (allIncreasing || allDecreasing ? 1 : 0)
		}, 0)
}

export function partTwo(input: string): number {
	return 0
}
