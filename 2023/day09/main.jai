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
	for line: String.split(data, "\n") {
		if line.count == 0  continue;
		numbers: [..]int;
		for String.split(line, " ")  Basic.array_add(*numbers, to_int(it));

		_, next := solve(numbers);
		total += next;
	}
	Basic.print("%\n", total);
}

part2 :: () {
	data, success := File.read_entire_file("input");
	Basic.assert(success);

	total := 0;
	for line: String.split(data, "\n") {
		if line.count == 0  continue;
		numbers: [..]int;
		for String.split(line, " ")  Basic.array_add(*numbers, to_int(it));

		previous, _ := solve(numbers);
		total += previous;
	}
	Basic.print("%\n", total);
}

solve :: (numbers: []int) -> int, int {
	end := numbers.count-1;
	previous: [..] int;

	next :=  numbers[end];
	while true {
		Basic.array_add(*previous, numbers[0]);
		done := true;
		for 0..end-1 {
			n1, n2 := numbers[it], numbers[it+1];
			diff := n2 - n1;
			if diff != 0  done = false;
			numbers[it] = diff;
		}
		end -= 1;

		next += numbers[end];

		if done break;
	}
	p := 0;
	for < previous p = it-p;
	return p, next;
}

to_int :: (s: string) -> int {
	i, success := Basic.to_integer(s);
	Basic.assert(success);
	return i;
}
