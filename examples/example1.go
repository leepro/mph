package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"time"

	"math/rand"

	"github.com/leepro/mph"
)

func SHA256(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	return h.Sum(nil)
}

func genKey() []byte {
	var temp [32]byte
	for i := 0; i < 32; i++ {
		temp[i] = byte(rand.Intn(254) + 1)
	}
	return SHA256(temp[:])
}

func main() {
	b := mph.Builder()

	println("start...")
	for i := 0; i < 100000; i++ {
		k := genKey()
		v := fmt.Sprintf("%d|%d", rand.Intn(9999999), rand.Intn(9999999))
		b.Add(k, []byte(v))
	}

	println("build...")
	st := time.Now()
	h, _ := b.Build()
	fmt.Printf("done... %d %f\n", len(h.HashParameters()), time.Since(st).Seconds())

	w, _ := os.Create("data.idx")

	h.Write(w)
	println("start...")

}
