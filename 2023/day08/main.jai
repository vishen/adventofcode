Basic :: #import "Basic";
File :: #import "File";
String :: #import "String";
Hash_Table :: #import "Hash_Table";

main :: () {
	part1();
	part2();
}

part1 :: () {
	data, success := File.read_entire_file("input");
	Basic.assert(success);

	dirs: []u8;
	map: Hash_Table.Table(string, [2]string);

	for line: String.split(data, "\n") {
		if line.count == 0  continue;

		if !dirs {
			dirs.data = line.data;
			dirs.count = line.count;
			continue;
		}

		key := String.slice(line, 0, 3);
		left := String.slice(line, 7, 3);
		right := String.slice(line, 12, 3);
		Hash_Table.table_add(*map, key, string.[left, right]);
	}

	next_dir :: (dirs: []u8, cursor: int) -> u8 {
		return dirs[cursor % dirs.count];
	}

	next := "AAA";
	steps := 0;
	while true {
		if next == "ZZZ"  break;

		defer steps += 1;

		rooms, found := Hash_Table.table_find(*map, next);
		Basic.assert(found);

		dir := next_dir(dirs, steps);
		if dir == {
			case #char "L";  next = rooms[0];
			case #char "R";  next = rooms[1];
		}
	}

	Basic.print("%\n", steps);
}

part2 :: () {
	data, success := File.read_entire_file("input");
	Basic.assert(success);

	dirs: []u8;
	map: Hash_Table.Table(string, [2]string);
	nodes: [..] string;

	for line: String.split(data, "\n") {
		if line.count == 0  continue;

		if !dirs {
			dirs.data = line.data;
			dirs.count = line.count;
			continue;
		}

		key := String.slice(line, 0, 3);
		left := String.slice(line, 7, 3);
		right := String.slice(line, 12, 3);
		Hash_Table.table_add(*map, key, string.[left, right]);

		if << (key.data + 2) == #char "A"  Basic.array_add(*nodes, key);
	}

	next_dir :: (dirs: []u8, cursor: int) -> u8 {
		return dirs[cursor % dirs.count];
	}

	steps_for_nodes: [..] int;
	for nodes {
		next := it;
		steps := 0;
		while true {
			if << (next.data + 2) == #char "Z"  break;

			defer steps += 1;

			rooms, found := Hash_Table.table_find(*map, next);
			Basic.assert(found);

			dir := next_dir(dirs, steps);
			if dir == {
				case #char "L";  next = rooms[0];
				case #char "R";  next = rooms[1];
			}
		}
		Basic.array_add(*steps_for_nodes, steps);
	}

	Basic.print("%\n", least_common_multiple(steps_for_nodes));
}

// Below taken from https://www.geeksforgeeks.org/program-to-find-lcm-of-two-numbers/

greatest_common_divisor :: (a: int, b: int) -> int {
	if a == 0  return b;
	return greatest_common_divisor(b % a, a);
}

least_common_multiple :: (arr: []int) -> int {
	// (a / gcd(a, b)) * b;
	value := 1;
	for arr  {
		value = (value / greatest_common_divisor(value, it)) * it;
	}
	return value;
}
