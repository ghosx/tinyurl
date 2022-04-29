package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

func BenchmarkWithFmt(b *testing.B) {
	fmt.Println("test_withfmt")
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%b\n", rand.Int63())
	}
}

func BenchmarkWithStrvonv(b *testing.B) {
	fmt.Println("test_strconv")
	for i := 0; i < b.N; i++ {
		strconv.FormatInt(rand.Int63(), 2)
	}
}
