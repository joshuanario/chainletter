package chainletter

// https://gist.github.com/Davidblkx/e12ab0bb2aff7fd8072632b396538560
func ComputeLevenshteinDistance(a []rune, b []rune) int {
	aLen := len(a)
	bLen := len(b)

	if aLen == 0 {
		return bLen
	}

	if bLen == 0 {
		return aLen
	}

	matrix := make([][]int, aLen+1)

	for i := range matrix {
		matrix[i] = make([]int, bLen+1)
	}

	for i := 0; i <= aLen; i++ {
		matrix[i][0] = i + 1
	}

	for j := 0; j <= bLen; j++ {
		matrix[0][j] = j + 1
	}

	for i := 1; i <= aLen; i++ {
		for j := 1; j <= bLen; j++ {
			cost := 0
			if b[j-1] == a[i-1] {
				cost = 0
			} else {
				cost = 1
			}
			matrix[i][j] = min(
				min(matrix[i-1][j]+1, matrix[i][j-1]+1),
				matrix[i-1][j-1]+cost)
		}
	}

	return matrix[aLen][bLen] - 1
}

func ComputeLevenshteinRatio(a []rune, b []rune) float64 {
	aLen := len(a)
	bLen := len(b)
	upperBound := float64(max(aLen, bLen))
	if upperBound == 0.0 {
		return 1.0
	}
	distance := float64(ComputeLevenshteinDistance(a, b))
	return (upperBound - distance) / upperBound
}
