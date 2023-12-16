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

	Basic.print("%\n", run(grid, EAST, .{0, 0}));
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

	highest := 0;
	for 0..grid.count-1 {
		highest = Basic.max(highest, run(grid, EAST, .{it, 0}));
		Basic.print("%: highest=%\n", it, highest);
		highest = Basic.max(highest, run(grid, WEST, .{it, grid[0].count-1}));
		Basic.print("%: highest=%\n", it, highest);
	}
	for 0..grid[0].count-1 {
		highest = Basic.max(highest, run(grid, SOUTH, .{0, it}));
		Basic.print("%: highest=%\n", it, highest);
		highest = Basic.max(highest, run(grid, NORTH, .{grid.count-1, it}));
		Basic.print("%: highest=%\n", it, highest);
	}
	Basic.print("%\n", highest);
}


run :: (grid: [..][..]u8, start_dir: Dir, start_pos: Pos) -> int {
	beams: [..]Beam;
	Basic.array_add(*beams, .{start_dir, start_pos});

	seen: [..]Beam;
	add :: (beams: *[..]Beam, beam: Beam, seen: *[..]Beam) {
		nb := beam;
		nb.pos.Y += beam.dir.Y;
		nb.pos.X += beam.dir.X;

		for << seen  {
			if it.pos == nb.pos && it.dir == nb.dir  return;
		}
		Basic.array_add(beams, nb);
		Basic.array_add(seen, nb);
	}

	visited: [..]Pos;

	MAX :: 1_000_000;
	i := 0;
	while i <= MAX {
		defer i+=1;
		if beams.count == 0  break;

		new_beams: [..]Beam;
		for beam_: beams {
			beam := beam_;

			if beam.pos.Y < 0 || beam.pos.Y >= grid.count  continue;
			if beam.pos.X < 0 || beam.pos.X >= grid[0].count  continue;

			add_if_unique(*visited, beam.pos);

			ch := grid[beam.pos.Y][beam.pos.X];
			if ch == {
			case;
				Basic.assert(false);
			case #char ".";
				add(*new_beams, beam, *seen);
			case #char "/";
				// rightward-moving beam that encounters a / mirror would continue upward in the mirror's column
				if beam.dir == NORTH  beam.dir = EAST;
				else if beam.dir == EAST  beam.dir = NORTH;
				else if beam.dir == SOUTH  beam.dir = WEST;
				else if beam.dir == WEST  beam.dir = SOUTH;
				add(*new_beams, beam, *seen);
			case #char "\\";
				// while a rightward-moving beam that encounters a \ mirror would continue downward from the mirror's column
				if beam.dir == NORTH  beam.dir = WEST;
				else if beam.dir == WEST  beam.dir = NORTH;
				else if beam.dir == EAST  beam.dir = SOUTH;
				else if beam.dir == SOUTH  beam.dir = EAST;
				add(*new_beams, beam, *seen);
			case #char "|";
				/*
				If the beam encounters the flat side of a splitter (| or -), the beam is split into two beams going in each of the two directions the splitter's pointy ends are pointing. For instance, a rightward-moving beam that encounters a | splitter would split into two beams: one that continues upward from the splitter's column and one that continues downward from the splitter's column.
				*/
				if beam.dir == EAST || beam.dir == WEST {
					add(*new_beams, .{NORTH, beam.pos}, *seen);
					add(*new_beams, .{SOUTH, beam.pos}, *seen);
				} else {
					add(*new_beams, beam, *seen);
				}
			case #char "-";
				/*
				If the beam encounters the flat side of a splitter (| or -), the beam is split into two beams going in each of the two directions the splitter's pointy ends are pointing. For instance, a rightward-moving beam that encounters a | splitter would split into two beams: one that continues upward from the splitter's column and one that continues downward from the splitter's column.
				*/
				if beam.dir == NORTH || beam.dir == SOUTH {
					add(*new_beams, .{EAST, beam.pos}, *seen);
					add(*new_beams, .{WEST, beam.pos}, *seen);
				} else {
					add(*new_beams, beam, *seen);
				}
			}
		}
		Basic.free(beams.data);
		beams = new_beams;
	}
	return visited.count;
}

Dir :: struct {
	Y, X := 0;
}

NORTH :: Dir.{-1, 0};
EAST  :: Dir.{0, 1};
SOUTH :: Dir.{1, 0};
WEST  :: Dir.{0, -1};

Pos :: struct {
	Y, X := 0;
}

Beam :: struct {
	dir: Dir;
	pos: Pos;
}

operator == :: (d1: Dir, d2: Dir) -> bool {
	return d1.Y == d2.Y && d1.X == d2.X;
}

operator == :: (p1: Pos, p2: Pos) -> bool {
	return p1.Y == p2.Y && p1.X == p2.X;
}

add_if_unique :: (visited: *[..]Pos, p: Pos) {
	for << visited  if it == p  return;
	Basic.array_add(visited, p);
}


print :: (grid: [..][..]u8) {
	for row: grid {
		s: string;
		s.data = row.data;
		s.count = row.count;
		Basic.print("%\n", s);
	}
}
