package maximum

import "errors"

var ErrNoMaximum = errors.New("no maximum element")

func MaxPair(numbers []float64) (float64, float64, error) {
	if len(numbers) <= 1 {
		return 0, 0, ErrNoMaximum
	}

	max1, max2 := startPair(numbers[0], numbers[1])
	for i := 2; i < len(numbers); i++ {
		if numbers[i] > numbers[max2] {
			tmp := max2
			max2 = i
			max1 = tmp
		} else {
			switch {
			case numbers[i] == numbers[max2]:
				continue
			case max1 == -1:
				max1 = i
			case numbers[i] > numbers[max1]:
				max1 = i
			}
		}
	}

	if max1 == -1 {
		return 0, 0, ErrNoMaximum
	}

	return numbers[max1], numbers[max2], nil
}

func startPair(a, b float64) (int, int) {
	switch {
	case a == b:
		return -1, 0
	case a > b:
		return 1, 0
	default:
		return 0, 1
	}
}
