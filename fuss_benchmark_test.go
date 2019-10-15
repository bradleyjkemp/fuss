package fuss

import (
	"math/rand"
	"net/http"
	"testing"
)

func BenchmarkFussHttpRequest(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(1))
	data := make([]byte, 40960)
	r.Read(data)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		req := &http.Request{}
		Seed(data).Fuzz(req)
	}
}
