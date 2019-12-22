package RedisPool

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	v, err := Get("t") // Get->抓1個值 其他的應該要不同method
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(v))
}

func TestB(t *testing.T) {
	v, err := HGet("Hyan", "A")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(v))
}

//Benchmark
func BenchmarkA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println(i)
	}
}
