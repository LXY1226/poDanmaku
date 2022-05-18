package bili

import (
	"log"
	"testing"
	"time"
)

var testBucket = NewTicketBucket(16, 1*time.Second, 2*time.Second)

func TestBucket(t *testing.T) {
	log.SetFlags(log.Ltime | log.Lmicroseconds)
	for i := 0; i < 100; i++ {
		go func(i int) {
			testBucket.Use()
			log.Print(i)
		}(i)
	}
	select {}
}
