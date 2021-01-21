// https://www.codewars.com/kata/5ce9c1000bab0b001134f5af/train/go
package quarter_of_the_year

func QuarterOf(month int) (result int) {
	quarters := []int{1, 2, 3, 4}

	switch month {
	case 1, 2, 3:
		result = quarters[0]
	case 4, 5, 6:
		result = quarters[1]
	case 7, 8, 9:
		result = quarters[2]
	case 10, 11, 12:
		result = quarters[3]
	}
	return
}
