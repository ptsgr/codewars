package ismyfriendcheating

func RemovNb(m uint64) (result [][2]uint64) {
	var sum uint64 = (m * (m + 1)) / 2

	for i := uint64(1); i <= m; i++ {
		j := float64(sum-i) / float64(i+1)
		if uint64(j) <= m && j == float64(uint64(j)) {
			result = append(result, [2]uint64{i, uint64(j)})
		}
	}
	return
}
