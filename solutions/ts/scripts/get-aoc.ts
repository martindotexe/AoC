export async function fetchInput(year: number, day: number): Promise<string> {
	// Load environment variables
	const envPath = new URL("../.env", import.meta.url).pathname;
	const envFile = await Bun.file(envPath).text().catch(() => "");

	const sessionCookie = envFile
		.split("\n")
		.find((line) => line.startsWith("AOC_SESSION="))
		?.split("=")[1]
		?.trim();

	if (!sessionCookie) {
		throw new Error(
			"AOC_SESSION not found in .env file. Please add your session cookie."
		);
	}

	const url = `https://adventofcode.com/${year}/day/${day}/input`;

	const response = await fetch(url, {
		headers: {
			Cookie: `session=${sessionCookie}`,
			"User-Agent": "github.com/yourname/advent-of-code via bun",
		},
	});

	if (!response.ok) {
		if (response.status === 404) {
			throw new Error(`Puzzle not found: ${year} day ${day}`);
		} else if (response.status === 400) {
			throw new Error("Invalid session cookie. Please check your .env file.");
		}
		throw new Error(`Failed to fetch input: ${response.status} ${response.statusText}`);
	}

	return await response.text();
}

// CLI usage
if (import.meta.main) {
	const args = process.argv.slice(2);

	if (args.length < 2) {
		console.error("Usage: bun run scripts/fetch-input.ts <year> <day>");
		process.exit(1);
	}

	const year = parseInt(args[0]!);
	const day = parseInt(args[1]!);

	if (isNaN(year) || isNaN(day)) {
		console.error("Year and day must be numbers");
		process.exit(1);
	}

	try {
		const input = await fetchInput(year, day);
		console.log(input);
	} catch (error) {
		console.error("Error:", error instanceof Error ? error.message : error);
		process.exit(1);
	}
}
