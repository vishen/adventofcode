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
		Utils.advance(*line, 5);
		id := Utils.next_int(*line);

		is_valid := true;
		while true {
			Utils.eat_until(*line, Utils.is_digit);
			number := Utils.next_int(*line);
			Utils.eat_until(*line, Utils.is_alpha);
			word := Utils.next_word(*line);

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

			if Utils.is_at_end(line)  break;
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

		Utils.advance(*line, 5);
		id := Utils.next_int(*line);

		max_red, max_blue, max_green := 0;

		while true {
			Utils.eat_until(*line, Utils.is_digit);
			number := Utils.next_int(*line);
			Utils.eat_until(*line, Utils.is_alpha);
			word := Utils.next_word(*line);
			if word == {
			case "red";  max_red = Basic.max(number, max_red);
			case "green";  max_green = Basic.max(number, max_green);
			case "blue";  max_blue = Basic.max(number, max_blue);
			}

			if Utils.is_at_end(line)  break;
		}

		total += max_red * max_green * max_blue;
	}

	Basic.print("%\n", total);
}

