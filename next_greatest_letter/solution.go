package next_greatest_letter

func nextGreatestLetter(letters []byte, target byte) byte {
	lowerBound := 0
	lastIdx := len(letters) - 1
	upperBound := lastIdx
	for mid := (lowerBound + upperBound) / 2; lowerBound <= upperBound; mid = (lowerBound + upperBound) / 2 {
		if target >= letters[mid] {
			if mid == lastIdx {
				return letters[0]
			}
			lowerBound = mid + 1
		} else if target < letters[mid] {
			if mid == lastIdx {
				return letters[mid]
			}
			upperBound = mid - 1
		}
	}
	if lowerBound >= lastIdx {
		lowerBound = 0
	}
	return letters[lowerBound]
}
