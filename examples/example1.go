package main

import (
	"fmt"
	"os"
	"time"

	"math/rand"

	"github.com/leepro/mph"
)

func main() {
	b := mph.Builder()

	println("start...")
	for i := 0; i < 1000000; i++ {
		k := fmt.Sprintf("%d", i)
		v := fmt.Sprintf("%d", rand.Intn(10))
		b.Add([]byte(k), []byte(v))
	}

	time.Sleep(2 * time.Second)

	println("build...")
	st := time.Now()
	h, _ := b.Build()
	println("done...", len(h.HashParameters()), time.Since(st).Seconds())

	w, _ := os.Create("data.idx")

	h.Write(w)
	println("start...")

}
