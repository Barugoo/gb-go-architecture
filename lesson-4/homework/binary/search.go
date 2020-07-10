package binary

func Search(startIndex int, endIndex int, compare func(index int) int) int {
	arrayCount := endIndex - startIndex
	half := arrayCount / 2

	newIndex := startIndex + half
	result := compare(newIndex)
	if result == 0 {
		return half
	}
	if arrayCount == 1 {
		return -1
	}

	if result == -1 {
		r := Search(newIndex, endIndex, compare)
		if r == -1 {
			return -1
		}
		return half + r
	} else {
		return Search(startIndex, newIndex, compare)
	}
}
