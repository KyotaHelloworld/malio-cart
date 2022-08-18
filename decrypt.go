package main

import "math/rand"

func numToString(num []int32, seed int64, rng int) string {
	decrypted := make([]rune, 0, len(num))
	rand.Seed(seed)
	cnt := 0

	for i := 0; i < len(num) && cnt >= 0; cnt++ {
		r := rand.Intn(rng)
		if num[i] == int32(cnt) {
			decrypted = append(decrypted, rune(r))
			i++
		}
	}

	return string(decrypted)
}
