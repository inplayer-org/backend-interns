package bookstore

func Cost(books []int) int {
	var one, two, three, four, five float64
	var popust int
	var total float64
	var totalBooks []float64
	var min float64
	for _, value := range books {
		if value == 1 {
			one++
		} else if value == 2 {
			two++
		} else if value == 3 {
			three++
		} else if value == 4 {
			four++
		} else if value == 5 {
			five++
		}
	}
	for one > 0 || two > 0 || three > 0 || four > 0 || five > 0 {
		popust = 0

		if one > 0 {
			popust++
		}
		one--

		if two > 0 {
			popust++
		}
		two--

		if three > 0 {
			popust++
		}
		three--

		if four > 0 {
			popust++
		}
		four--

		if five > 0 {
			popust++
		}
		five--

		if popust == 1 {
			total += 800
		}
		if popust == 2 {
			total += 2 * 800 * 0.95
		}
		if popust == 3 {
			total += 3 * 800 * 0.9
		}
		if popust == 4 {
			total += 4 * 800 * 0.8
		}
		if popust == 5 {
			if one == two && two == three && (one+two+three)/3 < (four+five)/2 ||
				one == two && two == four && (one+two+four)/3 < (three+five)/2 ||
				one == two && two == five && (one+two+five)/3 < (three+four)/2 ||
				two == three && three == four && (one+three+four)/3 < (two+five)/2 ||
				two == three && three == five && (two+three+five)/3 < (one+four)/2 ||
				three == four && four == five && (three+four+five)/3 < (one+two)/2 {
				total += 5 * 800 * 0.75
				continue
			}

			if one != two || one != three || one != four || one != five {
				total += 4 * 800 * 0.8
				min = one
				totalBooks = append(totalBooks, one, two, three, four, five)
				for _, value := range totalBooks {
					if value < min {
						min = value
					}
				}
				if one == min {
					one++
				} else if two == min {
					two++
				} else if three == min {
					three++
				} else if four == min {
					four++
				} else if five == min {
					five++
				}
			} else {
				total += 5 * 800 * 0.75
			}
		}
	}
	return int(total)
}
