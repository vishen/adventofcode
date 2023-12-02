Basic :: #import "Basic";
File :: #import "File";

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
