// https://www.codewars.com/kata/585d7d5adb20cf33cb000235/train/go
package unique_num

func FindUniq(arr []float32) float32 {
	uniqueMap := make(map[float32]int)

	for _, ele := range arr {
		uniqueMap[ele] += 1
	}

	for _, ele := range arr {
		if uniqueMap[ele] == 1 {
			return ele
		}
	}

	return float32(0)
}
