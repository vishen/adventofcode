Basic :: #import "Basic";
File :: #import "File";

main :: () {
	part1();
	part2();
}

part1 :: () {
	file_reader := load_file("input");

	total := 0;
	while true {
		line, success := next_line(*file_reader);
		if !success {
			break;
		}
		first, last := -1, -1;
		for line {
			if it >= #char "0" && it <= #char "9" {
				number := cast(int) (it - #char "0");
				if first == -1 {
					first = number;
					last = number;
				} else {
					last = number;
				}
			}
		}
		// Basic.print("first=%, last=%\n", first, last);
		total += (first * 10) + last;
	}
	Basic.print("total=%\n", total);
}

part2 :: () {
	file_reader := load_file("input");

	numbers_in_words := string.[
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	];

	total := 0;
	while true {
		line, success := next_line(*file_reader);
		if !success {
			break;
		}

		first, last := -1, -1;
		i := 0;

		while i < line.count {
			ch := line[i];
			number := -1;
			if ch >= #char "0" && ch <= #char "9" {
				number = cast(int) (ch - #char "0");
			} else {
				for numbers_in_words {
					s, success_peek := peek_as_string(line, i, it.count);
					if !success_peek  continue;

					if s == it {
						number = it_index + 1;
						i += it.count - 1; // STUPID, since we have i += 1 below
						break;
					}
				}
			}
			if number != -1 {
				if first == -1 {
					first = number;
					last = number;
				} else {
					last = number;
				}
			}

			i += 1;
		}
		// Basic.print("first=%, last=%\n", first, last);
		total += (first * 10) + last;
	}
	Basic.print("total=%\n", total);
}

peek_as_string :: (data: []u8, cursor: int, peek: int) -> string, success: bool {
	if cursor + peek > data.count return "", false;

	result: string;
	result.data = data.data + cursor;
	result.count = peek;

	return result, true;
}

File_Reader :: struct {
	data: *u8;
	cur: int;
	length: int;
}

advance :: (f: *File_Reader, amount := 1) {
	f.data += amount;
	f.cur += amount;
}

next_line :: (f: *File_Reader) -> []u8, success: bool {
	if f.cur >= f.length  return .[], false;

	start_pos := f.cur;
	start_data := f.data;
	while (<< f.data) != #char "\n" {
		advance(f);

		if f.cur >= f.length  break;
	}

	result: []u8;
	result.data = start_data;
	result.count = f.cur-start_pos;

	advance(f);

	return result, true;
}

load_file :: (filename: string) -> File_Reader {
	file, success := File.file_open(filename);
	Basic.assert(success);

	length := File.file_length(file);
	buffer := cast(*u8) Basic.alloc(length);

	success = File.file_read(file, buffer, length);
	Basic.assert(success);
	return File_Reader.{data=buffer, length=length};
}
