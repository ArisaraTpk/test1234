package test

func FindOddNumber(data []int) int {
	c := Count(data)

	return c.SelectOdd()
}

type CountKey map[int]int

func Count(data []int) CountKey {
	count := CountKey{}
	for _, val := range data {
		count[val] = count[val] + 1
	}
	return count
}

func (c CountKey) SelectOdd() int {
	selectOdd := -1
	for key, val := range c {
		if (val % 2) == 1 {
			selectOdd = key
			break
		}
	}
	return selectOdd
}
