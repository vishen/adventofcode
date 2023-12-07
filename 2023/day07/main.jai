Basic :: #import "Basic";
File :: #import "File";
String :: #import "String";
Sort :: #import "Sort";

main :: () {
	part1();
	part2();
}

part1 :: () {
	data, success := File.read_entire_file("input");
	Basic.assert(success);

	hands: [..] Hand;

	for line: String.split(data, "\n") {
		if line.count == 0  continue;

		line_u8: []u8;
		line_u8.data = line.data;
		line_u8.count = line.count;

		totals: [Card.Ace+1] int;
		cards: [5] Card;
		for Basic.array_view(line_u8, 0, 5) {
			card := to_card(it);
			cards[it_index] = card;
			totals[card] += 1;
		}

		bet := 0;
		for Basic.array_view(line_u8, 6) {
			if it >= #char "0" && it <= #char "9" {
				bet *= 10;
				bet += it - #char "0";
			}
		}
		Basic.array_add(*hands, Hand.{
			score=score(totals),
			cards=cards,
			bet=bet,
		});
	}

	sorted_hands := Sort.quick_sort(hands, (h1: Hand, h2: Hand) -> int {
		if h1.score > h2.score  return 1;
		if h1.score < h2.score  return -1;

		for h1.cards {
			h2_card := h2.cards[it_index];
			if it > h2_card  return 1;
			if it < h2_card  return -1;
		}
		return 0;
	});

	total := 0;
	for sorted_hands  total += ((it_index + 1) * it.bet);
	Basic.print("%\n", total);
}

part2 :: () {
	data, success := File.read_entire_file("input");
	Basic.assert(success);

	hands: [..] Hand;

	for line: String.split(data, "\n") {
		if line.count == 0  continue;

		line_u8: []u8;
		line_u8.data = line.data;
		line_u8.count = line.count;

		totals: [Card.Ace+1] int;
		cards: [5] Card;
		jacks := 0;
		for Basic.array_view(line_u8, 0, 5) {
			card := to_card(it);
			cards[it_index] = card;
			if card == .Jack {
				jacks += 1;
			} else {
				totals[card] += 1;
			}
		}

		score_:= -1;
		if jacks == 5 {
			score_ = 6;
		} else if jacks >= 1 {
			// If we have any Jacks, we need to add them to the card with the highest count
			highest, pos := 0;
			for totals {
				if highest < it  {
					highest = it;
					pos = it_index;
				}
			}
			totals[pos] += jacks;
		}
		if score_ == -1  score_ = score(totals);

		bet := 0;
		for Basic.array_view(line_u8, 6) {
			if it >= #char "0" && it <= #char "9" {
				bet *= 10;
				bet += it - #char "0";
			}
		}
		Basic.array_add(*hands, Hand.{
			score=score_,
			cards=cards,
			bet=bet,
		});
	}

	sorted_hands := Sort.quick_sort(hands, (h1: Hand, h2: Hand) -> int {
		if h1.score > h2.score  return 1;
		if h1.score < h2.score  return -1;

		for h1.cards {
			h2_card := h2.cards[it_index];
			
			if it != .Jack && h2_card == .Jack  return 1;
			if it == .Jack && h2_card != .Jack  return -1;

			if it > h2_card  return 1;
			if it < h2_card  return -1;
		}
		return 0;
	});

	total := 0;
	for sorted_hands  total += ((it_index + 1) * it.bet);
	Basic.print("%\n", total);
}

score :: (values: []int) -> int {
	three_of_a_kind := false;
	pairs := 0;

	for values {
		if it == {
			case 5; return 6;
			case 4; return 5;
			case 3;
				three_of_a_kind = true;
			case 2;
				pairs += 1;
		} 
	}

	if three_of_a_kind && pairs == 1 return 4;
	if three_of_a_kind  return 3;
	if pairs == 2  return 2;
	if pairs == 1  return 1;
	return 0;
}

Hand :: struct {
	score: int;
	cards: [5]Card;
	bet: int;
}

Card :: enum {
	One;
	Two;
	Three;
	Four;
	Five;
	Six;
	Seven;
	Eight;
	Nine;
	Ten;
	Jack;
	Queen;
	King;
	Ace;
}

to_card :: (ch: u8) -> Card {
	if ch == {
		case #char "A"; return .Ace;
		case #char "K"; return .King;
		case #char "Q"; return .Queen;
		case #char "J"; return .Jack;
		case #char "T"; return .Ten;
	}
	return cast(Card) ch - 1 - #char "0";
}

