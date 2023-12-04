Basic :: #import "Basic";
File :: #import "File";
String :: #import "String";

main :: () {
	part1();
	part2();
}

part1 :: () {
	data, success := File.read_entire_file("input");
	Basic.assert(success);

	total := 0;
	lines := String.split(data, "\n");
	for lines {
		if it.count == 0  continue;

		numbers := String.split(it, ":")[1];
		numbers_split := String.split(numbers, "|");
		winning_numbers, our_numbers := numbers_split[0], numbers_split[1];
		
		points := 0;
		for w: String.split(winning_numbers, " ") {
			if w == ""  continue;

			for o: String.split(our_numbers, " ") {
				if o == ""  continue;

				if o == w {
					if points == 0 {
						points = 1;
					} else {
						points *= 2;
					}
				}
			}
		}
		total += points;
	}
	Basic.print("%\n", total);
}

part2 :: () {
	data, success := File.read_entire_file("input");
	Basic.assert(success);

	cards: [500] int;

	total := 0;
	lines := String.split(data, "\n");
	for line, line_index: lines {
		if line.count == 0  continue;

		copies := cards[line_index];
		total += 1 + copies;

		numbers := String.split(line, ":")[1];
		numbers_split := String.split(numbers, "|");
		winning_numbers, our_numbers := numbers_split[0], numbers_split[1];
		
		matches := 0;
		for w: String.split(winning_numbers, " ") {
			if w == ""  continue;

			for o: String.split(our_numbers, " ") {
				if o == ""  continue;
				if o == w  matches += 1;
			}
		}
		for 0..matches-1 {
			index := line_index + 1 + it;
			cards[index] += 1+copies;
		}
	}
	Basic.print("%\n", total);
}
