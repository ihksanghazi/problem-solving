package challenges

func CompareScore(score1 [3]int32, score2 [3]int32) []int32 {
	a := 0
	b := 0

	for i := 0; i < len(score1); i++ {
		if score1[i] > score2[i] {
			a++
		} else if score1[i] < score2[i] {
			b++
		}
	}

	return []int32{int32(a), int32(b)}
}
