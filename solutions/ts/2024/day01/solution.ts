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
