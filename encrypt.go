package main

import "math/rand"

func stringToNumber(str string, seed int64, rng int) []int32 {
	encrypted := make([]int32, 0, len(str))
	rand.Seed(seed)

	bytes := []rune(str)

	cnt := int32(0)
	for i := 0; i < len(bytes) && cnt >= 0; cnt++ {
		r := rand.Int31n(int32(rng))
		k := bytes[i]
		_ = k

		if bytes[i] == r {
			encrypted = append(encrypted, cnt)
			i++
		}

	}

	return encrypted
}
