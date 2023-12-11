Basic :: #import "Basic";
File :: #import "File";
String :: #import "String";
Hash_Table :: #import "Hash_Table";
Sort :: #import "Sort";

main :: () {
	part1();
	// part2();
}

part1 :: () {
	data, success := File.read_entire_file("input");
	Basic.assert(success);

	char_at :: (data: string, offset: int) -> u8 {
		return << (data.data + offset);
	}

	start := 0;
	line_length := 0;
	for 0..data.count-1 {
		ch := char_at(data, it);

		if ch == #char "S" start = it;
		if line_length == 0 && ch == #char "\n" line_length = it+1;
	}

	Next :: struct {
		pos: int;
		count: int;
		last: int;
	}

	next: [..] Next;
	Basic.array_add(*next, .{start, 0, 0});

	can_continue :: (data: string, last: int, offset: int, looking_for: []u8) -> bool {
		if offset < 0 || offset  > data.count  return false;
		if last == offset  return false;

		ch := char_at(data, offset);
		for looking_for  if ch == it  return true;
		return false;
	}


	UP := -line_length;
	DOWN := line_length;
	LEFT := -1;
	RIGHT := 1;

	Dir :: struct {
		dir: int;
		looking_for: []u8;
	}

	steps := 0;
	while main_loop := steps < 100_000 {
		defer steps += 1;
		// Basic.print("%\n", next);
		new_next: [..] Next;
		for next: next {
			if steps > 0 && char_at(data, next.pos) == #char "S" {
				total := 0;
				if next.count & 0 == 0 {
					total = next.count / 2;
				} else {
					total = (next.count + 1) / 2;
				}
				Basic.print("%\n", total);
				break main_loop;
			}
			dirs: []Dir;
			ch := char_at(data, next.pos);
			if ch == {
			case #char "S";
				dirs = Dir.[
					.{DOWN, u8.[#char "L", #char "J", #char "|", #char "S"]},
					.{UP, u8.[#char "F", #char "7", #char "|", #char "S"]},
					.{LEFT, u8.[#char "L", #char "F", #char "-", #char "S"]},
					.{RIGHT, u8.[#char "7", #char "J", #char "-", #char "S"]},
				];
			case #char "|";
				dirs = Dir.[
					.{DOWN, u8.[#char "L", #char "J", #char "|", #char "S"]},
					.{UP, u8.[#char "F", #char "7", #char "|", #char "S"]},
				];
			case #char "-";
				dirs = Dir.[
					.{LEFT, u8.[#char "L", #char "F", #char "-", #char "S"]},
					.{RIGHT, u8.[#char "7", #char "J", #char "-", #char "S"]},
				];
			case #char "L";
				dirs = Dir.[
					.{UP, u8.[#char "F", #char "7", #char "|", #char "S"]},
					.{RIGHT, u8.[#char "7", #char "J", #char "-", #char "S"]},
				];
			case #char "J";
				dirs = Dir.[
					.{UP, u8.[#char "F", #char "7", #char "|", #char "S"]},
					.{LEFT, u8.[#char "L", #char "F", #char "-", #char "S"]},
				];
			case #char "7";
				dirs = Dir.[
					.{DOWN, u8.[#char "L", #char "J", #char "|", #char "S"]},
					.{LEFT, u8.[#char "L", #char "F", #char "-", #char "S"]},
				];
			case #char "F";
				dirs = Dir.[
					.{DOWN, u8.[#char "L", #char "J", #char "|", #char "S"]},
					.{RIGHT, u8.[#char "7", #char "J", #char "-", #char "S"]},
				];
			}

			for dirs {
				if can_continue(data, next.last, next.pos + it.dir, it.looking_for) {
					Basic.array_add(*new_next, .{next.pos + it.dir, next.count+1, next.pos});
				}
			}
		}
		Basic.free(next.data);
		next = new_next;
	}
};

/*
part2 :: () {
	data, success := File.read_entire_file("test5");
	Basic.assert(success);

	char_at :: (data: string, offset: int) -> u8 {
		return << (data.data + offset);
	}

	start := 0;
	line_length := 0;
	for 0..data.count-1 {
		ch := char_at(data, it);

		if ch == #char "S" start = it;
		if line_length == 0 && ch == #char "\n" line_length = it+1;
	}

	Next :: struct {
		pos: int;
		count: int;
		last: int;
		previous: [..]int;
	}

	next: [..] Next;
	previous: [..]int;
	Basic.array_add(*next, .{start, 0, 0, previous});

	can_continue :: (data: string, last: int, offset: int, looking_for: []u8) -> bool {
		if offset < 0 || offset  > data.count  return false;
		if last == offset  return false;

		ch := char_at(data, offset);
		for looking_for  if ch == it  return true;
		return false;
	}


	UP := -line_length;
	DOWN := line_length;
	LEFT := -1;
	RIGHT := 1;

	Dir :: struct {
		dir: int;
		looking_for: []u8;
	}

	route: [..]int;

	steps := 0;
	while main_loop := steps < 100_000 {
		defer steps += 1;
		// Basic.print("%\n", next);
		new_next: [..] Next;
		for next: next {
			if steps > 0 && char_at(data, next.pos) == #char "S" {
				// Basic.print("start=%, %\n", start, next.previous);
				route = next.previous;
				break main_loop;
			}
			dirs: []Dir;
			ch := char_at(data, next.pos);
			if ch == {
			case #char "S";
				dirs = Dir.[
					.{DOWN, u8.[#char "L", #char "J", #char "|", #char "S"]},
					.{UP, u8.[#char "F", #char "7", #char "|", #char "S"]},
					.{LEFT, u8.[#char "L", #char "F", #char "-", #char "S"]},
					.{RIGHT, u8.[#char "7", #char "J", #char "-", #char "S"]},
				];
			case #char "|";
				dirs = Dir.[
					.{DOWN, u8.[#char "L", #char "J", #char "|", #char "S"]},
					.{UP, u8.[#char "F", #char "7", #char "|", #char "S"]},
				];
			case #char "-";
				dirs = Dir.[
					.{LEFT, u8.[#char "L", #char "F", #char "-", #char "S"]},
					.{RIGHT, u8.[#char "7", #char "J", #char "-", #char "S"]},
				];
			case #char "L";
				dirs = Dir.[
					.{UP, u8.[#char "F", #char "7", #char "|", #char "S"]},
					.{RIGHT, u8.[#char "7", #char "J", #char "-", #char "S"]},
				];
			case #char "J";
				dirs = Dir.[
					.{UP, u8.[#char "F", #char "7", #char "|", #char "S"]},
					.{LEFT, u8.[#char "L", #char "F", #char "-", #char "S"]},
				];
			case #char "7";
				dirs = Dir.[
					.{DOWN, u8.[#char "L", #char "J", #char "|", #char "S"]},
					.{LEFT, u8.[#char "L", #char "F", #char "-", #char "S"]},
				];
			case #char "F";
				dirs = Dir.[
					.{DOWN, u8.[#char "L", #char "J", #char "|", #char "S"]},
					.{RIGHT, u8.[#char "7", #char "J", #char "-", #char "S"]},
				];
			}

			for dirs {
				if can_continue(data, next.last, next.pos + it.dir, it.looking_for) {
					previous: [..]int;
					Basic.array_reserve(*previous, next.previous.count+1);
					for next.previous  Basic.array_add(*previous, it);
					Basic.array_add(*previous, next.pos);
					Basic.array_add(*new_next, .{next.pos + it.dir, next.count+1, next.pos, previous});
				}
			}
		}
		Basic.free(next.data);
		next = new_next;
	}

	Basic.print("route=%\n", route);
	board: Hash_Table.Table(int, [..]int);
	for route {
		row := it / line_length;
		// col := it % line_length;
		arr, found := Hash_Table.table_find(*board, row);	
		if !found {
			arr = << Basic.New([..]int);
		} 
		Basic.array_add(*arr, it);
		Hash_Table.table_set(*board, row, arr);
	}

	// For for_expansion?
	#import "Hash_Table";

	total := 0;
	for values, key: board {
		Sort.quick_sort(values, (v1: int, v2: int) -> int { return v1 - v2; });

		next := 0;
		while true {
			start := values[next];


			// for start..end  if char_at(data, it + (key * line_length)) == #char "."  then total += 1;
			for start..end {
				if char_at(data, it) == #char "." {
					Basic.print("%, % -> start=%, end=%, it=%\n", key, values, start, end, it);
					Basic.print("%, %\n", key, it % line_length);
					total += 1;
				}
			}
		}
	}
	Basic.print("%\n", total);
};
*/
