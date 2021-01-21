// https://www.codewars.com/kata/5ce9c1000bab0b001134f5af/train/go
package quarter_of_the_year

func QuarterOf(month int) (result int) {
	switch month {
	case 1, 2, 3:
		result = 1
	case 4, 5, 6:
		result = 2
	case 7, 8, 9:
		result = 3
	case 10, 11, 12:
		result = 4
	}
	return
}
