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
	for part: String.split(data, ",") {
		if part.count == 0  continue;
		total += hash(part);
	}
	Basic.print("%\n", total);
}

hash :: (data: string) -> int {
	h := 0;
	for 0..data.count-1 {
		ch := << (data.data + it);
		if ch == #char "\n"  continue;

		h += ch;
		h *= 17;
		h %= 256;
	}
	return h;
}

part2 :: () {
	data, success := File.read_entire_file("input");
	Basic.assert(success);

	Lense :: struct {
		label: string;
		focal_length: u8;
		deleted := false;
	}

	boxes: [256][..]Lense;

	for part: String.split(data, ",") {
		label_end := 0;
		for 0..part.count-1 {
			ch := << (part.data + it);
			if ch < #char "a" || ch > #char "z" {
				label_end = it;
				break;
			}
		}

		label: string;
		label.data = part.data;
		label.count = label_end;

		label_hash := hash(label);

		op := << (part.data + label_end);
		if op == {
		case #char "=";
			focal_length  := << (part.data + label_end + 1);
			found := false;
			for * boxes[label_hash] {
				if it.label == label && !it.deleted {
					found = true;
					it.focal_length = focal_length;
					break;
				}
			}
			if !found Basic.array_add(*boxes[label_hash], .{label, focal_length, false});
		case #char "-";
			for * boxes[label_hash]  if it.label == label && !it.deleted  it.deleted = true;
		case;
			Basic.assert(false);
		}
	}

	total := 0;
	for lenses, box_index: boxes {
		cur := 1;
		for lenses {
			if it.deleted  continue;
			total += (box_index + 1) * cur * (xx it.focal_length - #char "0");
			cur += 1;
		}
	}
	Basic.print("%\n", total);
}
