package maximum

import "errors"

var ErrNoMaximum = errors.New("no maximum element")

func FindMaximum(numbers []float64) (float64, float64, error) {
	//init start max indexes
	max1, max2 := maxPair(numbers[0], numbers[1])

	for i := 2; i < len(numbers); i++ {
		if numbers[i] > numbers[max2] {
			tmp := max2
			max2 = i
			max1 = tmp
		} else {
			if max1 == -1 {
				max1 = i
			} else {
				if numbers[i] > numbers[max1] && numbers[i] != numbers[max2] {
					max1 = i
				}
			}
		}
	}

	if numbers[max1] == numbers[max2] {
		return 0, 0, ErrNoMaximum
	}

	return numbers[max1], numbers[max2], nil
}

func maxPair(a, b float64) (int, int) {
	switch {
	case a == b:
		return -1, 0
	case a > b:
		return 1, 0
	default:
		return 0, 1
	}
}
