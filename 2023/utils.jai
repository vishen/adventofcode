Basic :: #import "Basic";
File :: #import "File";

Line :: struct {
	data: []u8;
	end: *u8;
}

advance :: (line: *Line, amount := 1) {
	line.data.data += amount;
}

is_at_end :: (line: Line) -> bool {
	return line.data.data >= line.end;
}

next_int :: (line: *Line) -> (value: int) {
	Basic.assert(is_digit(line.data.data[0]));

	value := 0;
	while true {
		if !is_digit(line.data.data[0])  break;

		value *= 10;
		value += cast(int) (line.data.data[0] - #char "0");
		advance(line);
	}

	return value;
}

next_word :: (line: *Line) -> (value: string) {
	Basic.assert(is_alpha(line.data.data[0]));

	start := line.data.data;
	count := 0;
	while true {
		if is_at_end(line) break;
		if !is_alpha(line.data.data[0])  break;

		advance(line);
		count += 1;
	}

	s: string = ---;
	s.data = start;
	s.count = count;

	return s;
}


eat_until :: (line: *Line, until: (u8) -> bool) {
	while true {
		if until(line.data.data[0]) break;

		advance(line);
	}
}

is_digit :: (c: u8) -> bool {
	return c >= #char "0" && c <= #char "9";
}

is_alpha :: (c: u8) -> bool {
	return (c >= #char "a" && c <= #char "z") || (c >= #char "A" && c <= #char "Z");
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

next_line :: (f: *File_Reader) -> Line, success: bool {
	if f.cur >= f.length  return .{}, false;

	start_pos := f.cur;
	start_data := f.data;
	while (<< f.data) != #char "\n" {
		advance(f);

		if f.cur >= f.length  break;
	}

	d: []u8;
	d.data = start_data;
	d.count = f.cur-start_pos;

	line := Line.{data=d, end=d.data + d.count};

	advance(f);

	return line, true;
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
