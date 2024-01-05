package main

type coords struct {
	x int
	y int
}

func isPathCrossing(path string) bool {
	points := []rune(path)
	var x, y int
	previous := map[coords]struct{}{
		{0, 0}: {},
	}

	for _, point := range points {
		switch point {
		case 'N':
			y += 1
		case 'S':
			y -= 1
		case 'E':
			x += 1
		case 'W':
			x -= 1
		}

		c := coords{x, y}

		if _, ok := previous[c]; !ok {
			previous[c] = struct{}{}
		} else {
			return true
		}
	}

	return false
}
