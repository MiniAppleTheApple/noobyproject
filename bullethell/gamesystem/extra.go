package sys

import (
    "math/rand"
    "time"
)

func Random(num int32) int32{
    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
	return int32(r1.Intn(int(num + 1)))
}