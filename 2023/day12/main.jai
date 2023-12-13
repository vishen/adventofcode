Basic :: #import "Basic";
File :: #import "File";
String :: #import "String";
Hash_Table :: #import "Hash_Table";

main :: () {
	//part1();
	part2();
}

part1 :: () {
	data, success := File.read_entire_file("input");
	Basic.assert(success);

	total := 0;
	for line: String.split(data, "\n") {
		if line.count == 0  continue;

		// Basic.print("%\n", line);
		input_and_groups := String.split(line, " ");

		input: []u8;
		input.data = input_and_groups[0].data;
		input.count = input_and_groups[0].count;

		groups: [..]int;
		for String.split(input_and_groups[1], ",")  Basic.array_add(*groups, to_int(it));

		count := permutations(input, 0, groups);
		/*
		count := 0;
		for perms  {
			if is_valid(it, groups)  count += 1;
		}
		*/
		//Basic.print("%: %, count=% perms.count=%\n", it_index, to_string(input), count, perms.count);
		Basic.print("%: %\n", it_index, count);
		total += count;
		// break;
	}

	Basic.print("%\n", total);
}

part2 :: () {
	data, success := File.read_entire_file("test");
	Basic.assert(success);

	total := 0;
	for line: String.split(data, "\n") {
		if line.count == 0  continue;

		// Basic.print("%\n", line);
		input_and_groups := String.split(line, " ");

		input: [..]u8;
		input.data = input_and_groups[0].data;
		input.count = input_and_groups[0].count;

		groups: [..]int;
		for String.split(input_and_groups[1], ",")  Basic.array_add(*groups, to_int(it));

		expanded_input: [..]u8;
		expanded_groups: [..]int;
		for 0..4 {
			Basic.array_add(*expanded_input, .. input);
			if it < 4  Basic.array_add(*expanded_input, #char "?");
			
			Basic.array_add(*expanded_groups, .. groups);
		}


		//Basic.print("expanded_input=%\n", to_string(expanded_input));
		//Basic.print("expanded_groups=%\n", expanded_groups);

		count := permutations(expanded_input, 0, expanded_groups);
		Basic.print("%: %\n", it_index, count);
		total += count;
	}

	Basic.print("%\n", total);
}

to_string :: (v: []u8) -> string {
	s: string;
	s.data = v.data;
	s.count = v.count;
	return s;
}

permutations :: (input: []u8, cur: int, groups: [..]int, g: int, depth := 0) -> int {
	total := 0;

	q := -1;
	for cur..input.count-1 {
		ch := input[cur];
		if ch == #char "?" {
			q = it;
			break;
		}
		cur += 1;
	}
	if q == -1  {
		return ifx is_valid(input, groups) then 1 else 0;
	}

	{
		i: []u8;
		i.data = input.data;
		i.count = q;
		hg := hash_groups(i);
		defer Basic.free(hg.data);

		if hg.count > groups.count  return 0;
		for hg  {
			if it > groups[it_index]  return 0;
			if it_index < hg.count-1 && it < groups[it_index]  return 0;
		}

		//Basic.print("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$\n");
		// Basic.print("%,%: % - % - % - %\n", depth, permutations_ran, to_string(input), to_string(i), hg, groups);
	}

	for u8.[#char "#", #char "."] {
		// Basic.print("%, %\n", input.count, q);
		i := Basic.array_copy(input);
		i[q] = it;
		total += permutations(i, q+1, groups, depth+1);
		//defer Basic.free(i.data);
	}

	return total;
}

/*
permutations :: (input: []u8, depth := 0) -> [..][]u8 {
	perms: [..][]u8;
	q := -1;
	for input {
		if it == #char "?" {
			q = it_index;
			break;
		}
	}
	if q == -1 {
		Basic.array_add(*perms, input);
		return perms;
	}

	for u8.[#char "#", #char "."] {
		i := Basic.array_copy(input);
		i[q] = it;
		Basic.array_add(*perms, .. permutations(i, depth + 1));
	}

	return perms;
}
*/

hash_groups :: (input: []u8) -> [..]int {
	groups: [..]int;

	start_hash := -1;
	for input {
		if it == #char "#" {
			if start_hash == -1  start_hash = it_index;
			continue;
		}
		if start_hash >= 0 {
			Basic.array_add(*groups, it_index - start_hash);
			start_hash = -1;
		}
	}
	if start_hash >= 0 {
		Basic.array_add(*groups, input.count - start_hash);
	}
	return groups;
}

is_valid :: (input: []u8, groups: []int) -> bool {
	hg := hash_groups(input);
	defer Basic.free(hg.data);

	if groups.count != hg.count  return false;
	for hg  if it != groups[it_index]  return false;

	return true;
}

to_int :: (s: string) -> int {
	i, success := Basic.to_integer(s);
	Basic.assert(success, Basic.tprint("unable to convert % to int", s));
	return i;
}
