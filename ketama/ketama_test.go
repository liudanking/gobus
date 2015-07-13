package ketama

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func Benchmark_Hash(b *testing.B) {
	ring := NewRing(Base)
	ring.AddNode("node1", 1)
	ring.AddNode("node2", 1)
	ring.AddNode("node3", 1)
	ring.AddNode("node4", 1)
	ring.AddNode("node5", 1)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ring.Hash(strconv.Itoa(i))
	}
}

func TestHash(t *testing.T) {
	ring := NewRing(Base)
	ring.AddNode("node1", 1)
	ring.AddNode("node2", 1)

	ring.Bake()

	var (
		count1 = 0
		count2 = 0
	)
	for i := 0; i < 10000; i++ {
		node := ring.Hash(fmt.Sprintf("v:%d", i))
		switch node {
		case "node1":
			count1++
		case "node2":
			count2++
		default:
			fmt.Println("invalid node:", node)
		}
	}

	fmt.Printf("count1: %d, count2: %d", count1, count2)

	if math.Abs(float64(count1-count2)) > 100.00 {
		t.Error("Hash not average")
	}

}
