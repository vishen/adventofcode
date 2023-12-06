Basic :: #import "Basic";
File :: #import "File";
String :: #import "String";
Hash_Table :: #import "Hash_Table";

main :: () {
	part1();
	part2();
}

Map :: struct {
	from, to: string;
	name: string;
	ranges: [..] Range;
}

Range :: struct {
	destination_start: int;
	source_start: int;
	length: int;
}

part1 :: () {
	data, success := File.read_entire_file("input");
	Basic.assert(success);

	initial_seeds: [..] int;
	maps: [..] Map;

	cur_map: Map;

	for line: String.split(data, "\n") {
		if line.count == 0  {
			if cur_map.ranges.count > 0 {
				Basic.array_add(*maps, cur_map);

				cur_map = Map.{};
			}
			continue;
		}

		if initial_seeds.count == 0 {
			seeds := String.split(line, ": ");
			for seed: String.split(seeds[1], " ") {
				if seed.count == 0  continue;
				Basic.array_add(*initial_seeds, to_int(seed));
			}
			continue;
		}
		if !cur_map.name {
			name := String.split(line, " map")[0];
			ids := String.split(name, "-");
			cur_map.name = name;
			cur_map.from = ids[0];
			cur_map.to = ids[2];
		} else {
			range := String.split(line, " ");
			Basic.assert(range.count == 3);
			Basic.array_add(*cur_map.ranges, Range.{
				destination_start=to_int(range[0]),
				source_start=to_int(range[1]),
				length=to_int(range[2]),
			});
		}
	}

	lowest_location := 0;
	for seed: initial_seeds {
		location := find_src("seed", seed, maps);
		if it_index == 0 {
			lowest_location = location;
			continue;
		}

		lowest_location = Basic.min(lowest_location, location);
	}
	Basic.print("%\n", lowest_location);
}

part2 :: () {
	data, success := File.read_entire_file("input");
	Basic.assert(success);

	initial_seeds: [..] int;
	maps: [..] Map;

	cur_map: Map;

	for line: String.split(data, "\n") {
		if line.count == 0  {
			if cur_map.ranges.count > 0 {
				Basic.array_add(*maps, cur_map);

				cur_map = Map.{};
			}
			continue;
		}

		if initial_seeds.count == 0 {
			seeds := String.split(line, ": ");
			for seed: String.split(seeds[1], " ") {
				if seed.count == 0  continue;
				Basic.array_add(*initial_seeds, to_int(seed));
			}
			continue;
		}
		if !cur_map.name {
			name := String.split(line, " map")[0];
			ids := String.split(name, "-");
			cur_map.name = name;
			cur_map.from = ids[0];
			cur_map.to = ids[2];
		} else {
			range := String.split(line, " ");
			Basic.assert(range.count == 3);
			Basic.array_add(*cur_map.ranges, Range.{
				destination_start=to_int(range[0]),
				source_start=to_int(range[1]),
				length=to_int(range[2]),
			});
		}
	}

	lowest_location := -1;
	next_seed := 0;
	for 0..(initial_seeds.count/2)-1 {
		start := initial_seeds[next_seed];
		length := initial_seeds[next_seed+1];
		for seed: start..start+length {
			location := find_src("seed", seed, maps);
			if lowest_location == -1 {
				lowest_location = location;
				continue;
			}
			lowest_location = Basic.min(lowest_location, location);
		}
		next_seed += 2;
	}
	Basic.print("%\n", lowest_location);
}


/*
Key :: struct {
	from: string;
	dest: int;
}

compare :: (x: Key, y: Key) -> bool {
	return x.from == y.from && x.dest == y.dest;
}

cache: Hash_Table.Table(Key_Type=Key, Value_Type=int, given_compare_function=compare);
*/

cache: Hash_Table.Table(Key_Type=string, Value_Type=int);

find_src :: (from: string, dest: int, maps: [] Map) -> int {
	key := Basic.sprint("%.%", from, dest);
	src, found := Hash_Table.table_find(*cache, key=key);
	if found {
		return src;
	}
	m := find(from, maps);
	src = dest;
	for range: m.ranges {
		if in_range(dest, range) {
			src = (range.destination_start - range.source_start) + dest;
			break;
		}
	}
	if m.to == "location" {
		return src;
	}
	Hash_Table.table_add(*cache, key=key, value=src);
	return find_src(m.to, src, maps);
}

in_range :: (src: int, range: Range) -> bool {
	return (src >= range.source_start) && (src <= range.source_start + range.length);
}

find :: (from: string, maps: [] Map) -> Map {
	for maps  if it.from == from  return it;
	Basic.assert(false);
}

to_int :: (s: string) -> int {
	i, success := Basic.to_integer(s);
	Basic.assert(success);
	return i;
}
