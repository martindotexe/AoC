import { describe, expect, it } from "bun:test";
import { partOne, partTwo } from "./solution"

export const challenge_input = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`;


describe("Part one", () => {
	it("should return 11", () => {
		const result = partOne(challenge_input);
		expect(result).toBe(2);
	});
});


describe("Part two", () => {
	it("should return 31", () => {
		const result = partTwo(challenge_input);
		expect(result).toBe(4);
	});
});
