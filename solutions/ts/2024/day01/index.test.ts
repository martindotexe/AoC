import { describe, expect, it } from "bun:test";
import { partOne, partTwo } from "./index"

export const challenge_input = `3   4
4   3
2   5
1   3
3   9
3   3`;


describe("Part one", () => {
	it("should return 11", () => {
		const result = partOne(challenge_input);
		expect(result).toBe(11);
	});
});


describe("Part two", () => {
	it("should return 31", () => {
		const result = partTwo(challenge_input);
		expect(result).toBe(31);
	});
});
