Basic :: #import "Basic";
File :: #import "File";
String :: #import "String";
Sort :: #import "Sort";

main :: () {
	part1();
}

part1 :: () {
	data, success := File.read_entire_file("test");
	Basic.assert(success);

	grid: [..][..]u8;
	for line: String.split(data, "\n") {
		if line.count == 0  continue;
		row: [..]u8;
		for 0..line.count-1  Basic.array_add(*row, << (line.data + it));
		Basic.array_add(*grid, row);
	}

	Crucible :: struct {
		dir: Dir;
		pos: Pos;

		heat_loss := 0;
		direction_count := 0;
	}

	add :: (grid: [..][..]u8, cs: *[..]Crucible, c: Crucible, next_dir: Dir, seen: *[..]Crucible) {

		if c.dir == NORTH && next_dir == SOUTH  return;
		if c.dir == SOUTH && next_dir == NORTH  return;
		if c.dir == EAST && next_dir == WEST  return;
		if c.dir == WEST && next_dir == EAST  return;

		nc := c;
		nc.pos.Y += next_dir.Y;
		nc.pos.X += next_dir.X;

		if nc.pos.Y < 0 || nc.pos.Y >= grid.count  return;
		if nc.pos.X < 0 || nc.pos.X >= grid[0].count  return;

		if nc.dir == next_dir {
			nc.direction_count += 1;
			if nc.direction_count >= 4  return;
		}  else {
			nc.dir = next_dir;
			nc.direction_count = 1;
		}

		for << seen {
			if it.pos.Y == nc.pos.Y && it.pos.X == nc.pos.X && it.dir.Y == nc.dir.Y && it.dir.X == nc.dir.X && it.direction_count == nc.direction_count  return;
		}

		nc.heat_loss += (grid[nc.pos.Y][nc.pos.X] - #char "0");
		Basic.array_add(cs, nc);
		Basic.array_add(seen, nc);
	}

	seen: [..]Crucible;

	crucibles: [..]Crucible;
	{
		c: Crucible;
		c.pos = .{0, 0};
		add(grid, *crucibles, c, EAST, *seen);
		add(grid, *crucibles, c, SOUTH, *seen);
		Basic.array_add(*seen, c);
	}

	heat_loss := -1;

	i := 0;
	MAX :: 100000;
	while main_loop := i < MAX {
		defer i += 1;

		if crucibles.count == 0  break;

		Sort.quick_sort(crucibles, (c1: Crucible, c2: Crucible) -> int {
			return c1.heat_loss - c2.heat_loss;
		});

		// Basic.print("%: crucibles.count = %\n", i, crucibles.count);

		crucible := crucibles[0];
		Basic.array_unordered_remove_by_index(*crucibles, 0);
		if crucible.pos.Y >= grid.count-1 && crucible.pos.X >= grid[0].count - 1 {
			heat_loss = ifx heat_loss > -1 Basic.min(heat_loss, crucible.heat_loss) else crucible.heat_loss;
			Basic.print("heat_loss=%\n", crucible.heat_loss);
			// Basic.print("%\n", crucible.previous);
			// break main_loop;
			continue;
		}
		add(grid, *crucibles, crucible, NORTH, *seen);
		add(grid, *crucibles, crucible, SOUTH, *seen);
		add(grid, *crucibles, crucible, EAST, *seen);
		add(grid, *crucibles, crucible, WEST, *seen);
	}

	Basic.print("%\n", heat_loss);
}

Dir :: struct {
	Y, X := 0;
}

operator == :: (d1: Dir, d2: Dir) -> bool {
	return d1.Y == d2.Y && d1.X == d2.X;
}

NORTH :: Dir.{-1, 0};
EAST  :: Dir.{0, 1};
SOUTH :: Dir.{1, 0};
WEST  :: Dir.{0, -1};

Pos :: struct {
	Y, X := 0;
}
