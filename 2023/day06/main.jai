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

	times: [..] int;
	distances: [..] int;

	lines := String.split(data, "\n");

	for String.split(String.split(lines[0], ":")[1], " ") {
		if it  Basic.array_add(*times, to_int(it));
	}
	for String.split(String.split(lines[1], ":")[1], " ") {
		if it  Basic.array_add(*distances, to_int(it));
	}

	total := 1;
	for time, time_index: times {
		results := simulate_results(time);
		ways_to_beat_record := 0;
		for results {
			if it > distances[time_index]  ways_to_beat_record += 1;
		}
		total *= ways_to_beat_record;
	}
	Basic.print("%\n", total);
}

part2 :: () {
	data, success := File.read_entire_file("input");
	Basic.assert(success);

	lines := String.split(data, "\n");

	time := parse_int(String.split(lines[0], ":")[1]);
	distance := parse_int(String.split(lines[1], ":")[1]);

	results := simulate_results(time);
	ways_to_beat_record := 0;
	for results {
		if it > distance  ways_to_beat_record += 1;
	}
	Basic.print("%\n", ways_to_beat_record);
}

simulate_results :: (time: int) -> [..] int {
	results: [..] int;
	// TODO(vishen): This can be reduced in half, since the numbers are
	// exactly the same after the half-way point. But seems to run fast enough
	// as is.
	for ms: 0..time {
		Basic.array_add(*results, (time-ms) * ms);
	}
	return results;
}

parse_int :: (s: string) -> int {
	value := 0;

	bytes: []u8;
	bytes.data = s.data;
	bytes.count = s.count;
	for bytes {
		if it >= #char "0" && it <= #char "9" {
			if value == 0 {
				value = it - #char "0";
			} else {
				value *= 10;
				value += it - #char "0";
			}
		}
	}
	return value;
}


to_int :: (s: string) -> int {
	i, success := Basic.to_integer(s);
	Basic.assert(success);
	return i;
}
