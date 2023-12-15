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

	grid: [..][..]u8;

	for line: String.split(data, "\n") {
		if line.count == 0  continue;
		row: [..]u8;
		for 0..line.count-1  Basic.array_add(*row, << (line.data + it));
		Basic.array_add(*grid, row);
	}

	load := 0;

	row := 0;
	while row < grid.count {
		cols := grid[row];
		defer row += 1;

		for col, col_index: cols {
			if col == #char "#" || col == #char "."  continue;
			
			new_row := row;
			for 1..row {
				if grid[row-it][col_index] == {
				case #char "#"; #through;
				case #char "O"; 
					break;
				}
				new_row -= 1;
			}
			grid[row][col_index] = #char ".";
			grid[new_row][col_index] = #char "O";
			load += grid.count-new_row;
		}
		
	}

	Basic.print("%\n", load);
}

part2 :: () {
	data, success := File.read_entire_file("input");
	Basic.assert(success);

	grid: [..][..]u8;

	for line: String.split(data, "\n") {
		if line.count == 0  continue;
		row: [..]u8;
		for 0..line.count-1  Basic.array_add(*row, << (line.data + it));
		Basic.array_add(*grid, row);
	}

	Dir :: enum {
		NORTH :: 1;
		WEST;
		SOUTH;
		EAST;
	}

	dirs := Dir.[.NORTH, .WEST, .SOUTH, .EAST];

	previous: [..]string;
	loads: [..]int;

	first_iteration := 0;
	cycled_iteration := 0;

	MAX :: 1000000000;
	for iteration: 1..MAX {

		for dir: dirs {
			row_ := 0;
			while row_ < grid.count {
				row := ifx dir == .SOUTH  grid.count-1 - row_ else row_;
				col_ := 0;
				cols := grid[row];
				defer row_ += 1;

				while col_ < cols.count {
					col := ifx dir == .EAST  cols.count-1 - col_ else col_;
					defer col_ += 1;
					ch := cols[col];
					if ch == #char "#" || ch == #char "."  continue;
					
					col_d := 0;
					row_d := 0;
					length := 0;

					if dir == {
					case .NORTH;
						row_d = -1;
						length = row - 1;
					case .SOUTH;
						row_d = 1;
						length = grid.count - row - 1;
					case .EAST;
						col_d = 1;
						length = cols.count - col - 1;
					case .WEST;
						col_d = -1;
						length = col - 1;
					}

					new_row := row;
					new_col := col;
					moving := false;

					for 0..length {
						next_row := new_row + row_d;
						next_col := new_col + col_d;
						if next_row < 0 || next_row >= grid.count  break;
						if next_col < 0 || next_col >= cols.count  break;

						if grid[next_row][next_col] == {
						case #char "#"; #through;
						case #char "O"; 
							break;
						}
						new_row = next_row;
						new_col = next_col;
						moving = true;
					}
					if !moving  continue;

					grid[row][col] = #char ".";
					grid[new_row][new_col] = #char "O";
				}
				
			}
		}
		for previous {
			if grid_to_string(grid) == it {
				cycled_iteration = iteration;
				first_iteration = it_index;
				break iteration;
			}
		}
		Basic.array_add(*previous, grid_to_string(grid));
		load := 0;
		for row, row_index: grid {
			for col: row {
				if col == #char "O"  load += grid.count - row_index;
			}
		}
		Basic.array_add(*loads, load);
	}

	expected_iteration := (MAX - first_iteration) % (cycled_iteration - 1 - first_iteration);
	Basic.print("%\n", loads[first_iteration + (expected_iteration - 1)]);
}

grid_to_string :: (grid: [..][..]u8) -> string {
	flattened: [..]u8;
	for row: grid {
		for col: row  Basic.array_add(*flattened, col);
	}

	s: string;
	s.data = flattened.data;
	s.count = flattened.count;
	return s;
}
