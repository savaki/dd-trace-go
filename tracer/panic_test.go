package tracer

import (
	"math/rand"
	"sync"
	"testing"
)

const n = 1000000

func TestPanicEncodeTraces(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			payload := getTestTrace(rand.Intn(3), rand.Intn(100))
			encoder := NewMsgpackEncoder()
			encoder.EncodeTraces(payload)
		}()
	}
	wg.Wait()

	println("Test EncodeTraces done.")
}

func TestPanicSendTraces(t *testing.T) {
	_, transport := getTestTracer()

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			payload := getTestTrace(rand.Intn(3), rand.Intn(100))
			transport.SendTraces(payload)
		}()
	}
	wg.Wait()

	println("Test SendTraces done.")
}
