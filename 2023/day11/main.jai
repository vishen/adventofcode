Basic :: #import "Basic";
File :: #import "File";
Sort :: #import "Sort";

main :: () {
	part1();
	part2();
}

part1 :: () {
	data, success := File.read_entire_file("input");
	Basic.assert(success);

	char_at :: (data: string, offset: int) -> u8 {
		return << (data.data + offset);
	}

	galaxies: [..]Pos;
	galaxies_by_row: [..]int;
	galaxies_by_col: [..]int;

	line_length := 0;
	rows := 0;
	for 0..data.count-1 {
		ch := char_at(data, it);
		if ch == #char "\n" {
			if line_length == 0  line_length = it + 1;
			if rows == 0 rows = it;
			continue;
		}
		row := ifx line_length > 0 then it/line_length else 0;
		col := ifx line_length > 0 then it % line_length else it;
		if ch == #char "#"{
			Basic.array_add(*galaxies, .{row=row, col=col});
			Basic.array_add_if_unique(*galaxies_by_row, row);
			Basic.array_add_if_unique(*galaxies_by_col, col);
		}
	}

	cols := line_length - 2;
	rows -= 1;

	Sort.quick_sort(galaxies_by_col, (i: int, j: int) -> int {
		return i - j;
	});

	{
		prev := 0;
		expanded := 0;
		for row: galaxies_by_row {
			if it_index > 0 && prev + 1 != row {
				expanded += row - (prev + 1);
			}
			for * galaxies {
				if it.row == row && !it.row_expanded  {
					it.row += expanded;
					it.row_expanded = true;
				}
			}
			prev = row;
		}
		rows += expanded;
	}
	{
		prev := 0;
		expanded := 0;
		for col: galaxies_by_col {
			if it_index > 0 && prev + 1 != col {
				expanded += col - (prev + 1);
			}
			for * galaxies {
				if it.col == col && !it.col_expanded  {
					it.col += expanded;
					it.col_expanded = true;
				}
			}
			prev = col;
		}
		cols += expanded;
	}

	total := 0;
	for outer: 0..galaxies.count-1 {
		for outer+1..galaxies.count-1 {
			total += distance(galaxies[outer], galaxies[it]);
		}
	}
	Basic.print("%\n", total);
}


part2 :: () {
	data, success := File.read_entire_file("input");
	Basic.assert(success);

	char_at :: (data: string, offset: int) -> u8 {
		return << (data.data + offset);
	}

	galaxies: [..]Pos;
	galaxies_by_row: [..]int;
	galaxies_by_col: [..]int;

	line_length := 0;
	rows := 0;
	for 0..data.count-1 {
		ch := char_at(data, it);
		if ch == #char "\n" {
			if line_length == 0  line_length = it + 1;
			if rows == 0 rows = it;
			continue;
		}
		row := ifx line_length > 0 then it/line_length else 0;
		col := ifx line_length > 0 then it % line_length else it;
		if ch == #char "#"{
			Basic.array_add(*galaxies, .{row=row, col=col});
			Basic.array_add_if_unique(*galaxies_by_row, row);
			Basic.array_add_if_unique(*galaxies_by_col, col);
		}
	}

	cols := line_length - 2;
	rows -= 1;

	Sort.quick_sort(galaxies_by_col, (i: int, j: int) -> int {
		return i - j;
	});

	{
		prev := 0;
		expanded := 0;
		for row: galaxies_by_row {
			expanded_row := row + expanded;
			if it_index > 0 && prev + 1 != expanded_row {
				expanded += (expanded_row - (prev + 1)) * (1000000 - 1);
				expanded_row = row + expanded;
			}
			for * galaxies {
				if it.row == row && !it.row_expanded  {
					it.row += expanded;
					it.row_expanded = true;
				}
			}
			prev = expanded_row;
		}
		rows += expanded;
	}
	{
		prev := 0;
		expanded := 0;
		for col: galaxies_by_col {
			expanded_col := col + expanded;
			if it_index > 0 && prev + 1 != expanded_col {
				expanded += (expanded_col - (prev + 1)) * (1000000 - 1);
				expanded_col = col + expanded;
			}
			for * galaxies {
				if it.col == col && !it.col_expanded  {
					it.col += expanded;
					it.col_expanded = true;
				}
			}
			prev = expanded_col;
		}
		cols += expanded;
	}

	total := 0;
	for outer: 0..galaxies.count-1 {
		for outer+1..galaxies.count-1 {
			total += distance(galaxies[outer], galaxies[it]);
		}
	}
	Basic.print("%\n", total);
}

Pos :: struct {
	row, col: int;

	// TODO: This is so stupid;
	row_expanded := false;
	col_expanded := false;
}


abs :: (i: int, j: int) -> int {
	diff := i - j;
	if diff >= 0 return diff;
	return -diff;
}

distance :: (pos: Pos, looking_for: Pos) -> int {
	diff_row := abs(looking_for.row, pos.row);
	diff_col := abs(looking_for.col, pos.col);
	return diff_row + diff_col;
}
