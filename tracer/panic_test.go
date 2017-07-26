package tracer

import (
	"math/rand"
	"sync"
	"testing"
)

const n = 10000

func TestPanicEncodeTraces(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			payload := getTestTrace(rand.Intn(3), rand.Intn(100))
			encoder := newMsgpackEncoder()
			encoder.EncodeTraces(payload)
		}()
	}
	wg.Wait()

	println("Test EncodeTraces done.")
}

func TestPanicSendTraces(t *testing.T) {
	transport := newHTTPTransport(defaultHostname, defaultPort)
	transport.traceURL = "http://localhost:8126/v0.0/traces"

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
