Basic :: #import "Basic";
// Utils :: #import,file "../utils.jai";
File :: #import "File";

main :: () {
	part1();
	part2();
}

part1 :: () {
	data_, success := File.read_entire_file("input");
	Basic.assert(success);
	
	data: []u8;
	data.data = data_.data;
	data.count = data_.count;

	line_length := 0;
	for data {
		if it == #char "\n" {
			line_length = it_index + 1;
			break;
		}
	}

	cursor := 0;
	eating_number := false;
	number := 0;
	adjacent_symbol := false;
	total := 0;
	while cursor < data.count {
		ch := data[cursor];
		defer cursor += 1;
		
		if ch >= #char "0" && ch <= #char "9" {
			if !eating_number {
				eating_number = true;
			}
			number *= 10;
			number += cast(int) (ch - #char "0");

			for directions {
				pos := cursor + it.x + (it.y*line_length);
				if is_valid_and_symbol(data, pos) {
					adjacent_symbol = true;
					break;
				}
			}
		} else {
			if eating_number {
				if adjacent_symbol  total += number;

				eating_number = false;
				number = 0;
				adjacent_symbol = false;
			}
		}
	}
	Basic.print("%\n", total);
}

part2 :: () {
	data_, success := File.read_entire_file("input");
	Basic.assert(success);
	
	data: []u8;
	data.data = data_.data;
	data.count = data_.count;

	line_length := 0;
	for data {
		if it == #char "\n" {
			line_length = it_index + 1;
			break;
		}
	}


	Number_And_Gear_Location :: struct {
		number: int;
		gear_location: int;
	}

	locs: [..] Number_And_Gear_Location;

	cursor := 0;
	eating_number := false;
	number := 0;
	gear_location := -1;

	while cursor < data.count {
		ch := data[cursor];
		defer cursor += 1;
		
		if ch >= #char "0" && ch <= #char "9" {
			if !eating_number {
				eating_number = true;
			}
			number *= 10;
			number += cast(int) (ch - #char "0");

			for directions {
				pos := cursor + it.x + (it.y*line_length);
				if is_gear(data, pos) {
					gear_location = pos;
					break;
				}
			}
		} else {
			if eating_number {
				if gear_location != -1  Basic.array_add(*locs, Number_And_Gear_Location.{number=number, gear_location=gear_location});

				eating_number = false;
				number = 0;
				gear_location = -1;
			}
		}
	}

	gear_ratio := 0;
	for loc, loc_index: locs {
		_gear_ratio := loc.number;
		count := 1;
		for loc_index+1..locs.count-1 {
			if loc.gear_location == locs[it].gear_location {
				_gear_ratio *= locs[it].number;
				count += 1;
			}
		}
		if count == 2  gear_ratio += _gear_ratio;
		
	}

	Basic.print("%\n", gear_ratio);
}

Direction :: struct {
	x, y: int;
}

directions :: Direction.[
	.{0, 1},
	.{0, -1},
	.{1, 0},
	.{-1, 0},
	.{1, 1},
	.{-1, 1},
	.{1, -1},
	.{-1, -1},
];

is_valid_and_symbol :: (data: []u8, cursor: int) -> bool {
	if cursor < 0 || cursor >= data.count   return false;

	ch := data[cursor];
	return !(ch == #char "." || ch == #char "\n" || is_digit(ch));
}

is_gear :: (data: []u8, cursor: int) -> bool {
	if cursor < 0 || cursor >= data.count   return false;

	return data[cursor] == #char "*";
}

is_digit :: (c: u8) -> bool {
	return c >= #char "0" && c <= #char "9";
}
