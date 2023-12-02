Basic :: #import "Basic";
Utils :: #import,file "../utils.jai";

main :: () {
	part1();
	part2();
}

part1 :: () {
	// fr := Utils.load_file("test");
	fr := Utils.load_file("input");

	// only 12 red cubes, 13 green cubes, and 14 blue cubes

	max_red, max_green, max_blue := 12, 13, 14;

	ids := 0;
	while true {
		line, success := Utils.next_line(*fr);
		if !success {
			break;
		}
		//Basic.print("line=%\n", line);
		id, cursor := next_int(line, 5);
		// Basic.print("id=%, cursor=%\n", id, cursor);

		is_valid := true;
		number, word := 0, "";
		while true {
			cursor = eat_until(line, cursor, is_digit);
			number, cursor = next_int(line, cursor);
			cursor = eat_until(line, cursor, is_alpha);
			word, cursor = next_word(line, cursor);
			// Basic.print("word=%, number=%\n", word, number);
			if word == {
			case "red";
				if number > max_red {
					is_valid = false;
					break;
				}
			case "green";
				if number > max_green {
					is_valid = false;
					break;
				}
			case "blue";
				if number > max_blue {
					is_valid = false;
					break;
				}
			}

			if cursor >= line.count  break;
		}

		if is_valid  ids += id;
	}

	Basic.print("%\n", ids);
}

part2 :: () {
	// fr := Utils.load_file("test");
	fr := Utils.load_file("input");

	total := 0;
	while true {
		line, success := Utils.next_line(*fr);
		if !success {
			break;
		}
		id, cursor := next_int(line, 5);

		max_red, max_blue, max_green := 0;

		number, word := 0, "";
		while true {
			cursor = eat_until(line, cursor, is_digit);
			number, cursor = next_int(line, cursor);
			cursor = eat_until(line, cursor, is_alpha);
			word, cursor = next_word(line, cursor);
			if word == {
			case "red";  max_red = Basic.max(number, max_red);
			case "green";  max_green = Basic.max(number, max_green);
			case "blue";  max_blue = Basic.max(number, max_blue);
			}

			if cursor >= line.count  break;
		}

		total += max_red * max_green * max_blue;
	}

	Basic.print("%\n", total);
}

next_int :: (data: []u8, cursor: int) -> (value: int, cursor: int) {
	Basic.assert(is_digit(data[cursor]));

	value := 0;
	while true {
		if !is_digit(data[cursor])  break;

		value *= 10;
		value += cast(int) (data[cursor] - #char "0");
		cursor += 1;
	}

	return value, cursor;
}

next_word :: (data: []u8, cursor: int) -> (value: string, cursor: int) {
	Basic.assert(is_alpha(data[cursor]));

	start := cursor;
	while true {
		if cursor >= data.count  break;
		if !is_alpha(data[cursor])  break;

		cursor += 1;
	}

	s: string = ---;
	s.data = data.data + start;
	s.count = cursor - start;

	return s, cursor;
}

is_digit :: (c: u8) -> bool {
	return c >= #char "0" && c <= #char "9";
}

is_alpha :: (c: u8) -> bool {
	return (c >= #char "a" && c <= #char "z") || (c >= #char "A" && c <= #char "Z");
}

eat_until :: (data: []u8, cursor: int, until: (u8) -> bool) -> cursor: int {
	while true {
		if until(data[cursor]) break;

		cursor += 1;
	}
	return cursor;
}
