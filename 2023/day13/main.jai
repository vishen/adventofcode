Basic :: #import "Basic";
File :: #import "File";
String :: #import "String";

main :: () {
	// part1();
	part2();
}

part1 :: () {
	data, success := File.read_entire_file("input");
	Basic.assert(success);

	total := 0;
	for lines: String.split(data, "\n\n") {
		if lines.count == 0  continue;

		// Horizontal
		horizontal := String.split(String.trim(lines, "\n"), "\n");

		split, success := find(horizontal);
		if success  {
			total += split * 100;
			continue;
		}

		// Vertical
		width := horizontal[0].count;
		height := horizontal.count;

		vertical: [..]string;
		for w: 0..width-1 {
			s: [..]u8;
			for h: 0..height-1 Basic.array_add(*s, << (horizontal[h].data + w));

			s_: string;
			s_.data = s.data;
			s_.count = s.count;
			Basic.array_add(*vertical, s_);
		}

		split, success = find(vertical);
		Basic.assert(success);
		total += split;
	}
	Basic.print("%\n", total);
}

part2 :: () {
	data, success := File.read_entire_file("input");
	Basic.assert(success);

	total := 0;
	for lines_: String.split(data, "\n\n") {
		lines := String.trim(lines_, "\n");
		if lines.count == 0  continue;

		original, success := attempt(lines);
		Basic.assert(success);

		found := false;
		for 0..lines.count-1 {
			ch := << (lines.data + it);
			if ch == #char "\n"  continue;
			nline := string_copy(lines);
			set_to := ifx ch == #char "#" then #char "." else #char "#";
			<< (nline.data+it) = xx set_to;
			split, success := attempt(nline, original);
			if success && split != original {
				total += split;
				found = true;
				break;
			}
		}
		if !found  total += original;

	}
	Basic.print("%\n", total);
}

string_copy :: (s: string) -> string {
	n: [..]u8;
	for 0..s.count-1  Basic.array_add(*n, << (s.data + it));

	ns: string;
	ns.data = n.data;
	ns.count = n.count;
	return ns;
}

attempt :: (lines: string, original := 0) -> (int, bool) {
	horizontal := String.split(String.trim(lines, "\n"), "\n");
	for 0..horizontal.count-1 {
		split, success := find(horizontal, it);
		if success && (split * 100) != original  {
			return split * 100, true;
		}
	}

	// Vertical
	width := horizontal[0].count;
	height := horizontal.count;

	vertical: [..]string;
	for w: 0..width-1 {
		s: [..]u8;
		for h: 0..height-1 Basic.array_add(*s, << (horizontal[h].data + w));

		s_: string;
		s_.data = s.data;
		s_.count = s.count;
		Basic.array_add(*vertical, s_);
	}

	for 0..vertical.count-1 {
		split, success := find(vertical, it);
		if success && split != original {
			return split, true;
		}
	}
	return 0, false;
}

find :: (lines: []string, start := 0) -> int, bool {
	cur := start + 1;
	for cur..lines.count-1 {
		defer cur += 1;

		valid := true;
		count := 1;
		while true {
			defer count += 1;

			if cur+(count-1) >= lines.count || cur-count < 0  break;

			if lines[cur+(count-1)] != lines[cur-count] {
				valid = false;
				break;
			}
		}
		if valid {
			return cur, true;
		}
	}
	return 0, false;
}
