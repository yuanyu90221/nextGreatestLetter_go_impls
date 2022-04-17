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
	// nextLarge := 'a' + (target+1-'a')%26
	// for upperBound > lowerBound {
	// 	mid := (upperBound + lowerBound) / 2
	// 	if letters[mid] < nextLarge {
	// 		lowerBound = mid + 1
	// 	}
	// 	if letters[mid] > nextLarge {
	// 		upperBound = mid - 1
	// 	}
	// 	if letters[mid] == nextLarge {
	// 		return letters[mid]
	// 	}
	// }
	// mid := (upperBound + lowerBound) / 2
	// // fmt.Printf("target %c, mid: %c, nextLarge: %c\n", target, letters[mid], nextLarge)
	// if mid == len(letters)-1 && letters[mid] <= target {
	// 	return letters[0]
	// }
	// if letters[mid] >= nextLarge {
	// 	return letters[mid]
	// }
	// if mid >= 1 && letters[mid-1] > nextLarge {
	// 	return letters[mid-1]
	// } else if mid < len(letters)-1 && letters[mid+1] >= nextLarge {
	// 	return letters[mid+1]
	// }
	// return letters[mid]
}
