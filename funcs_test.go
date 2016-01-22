package bench_channel

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
)

var iterations int
var err error

func init() {
	stubChan <- false
	stubBool = true
	rand.Seed(time.Now().UTC().UnixNano())
	iterations = 5000000
	readPercStr := os.Getenv("READ_PERCENTAGE")
	if readPercStr == "" {
		panic("READ_PERCENTAGE env variable not found")
	}
	tempReadPercentage, err := strconv.ParseInt(readPercStr, 10, 64)
	if err != nil {
		log.Println("READ_PERCENTAGE env variable not found")
	}
	readPercentage = int(tempReadPercentage)
	log.Println("Read percentage is: ", readPercentage)

}
func BenchmarkChan(b *testing.B) {
	// b.N = iterations
	for i := 0; i < b.N; i++ {
		go WithChannel()
	}
}

func BenchmarkLock(b *testing.B) {
	// b.N = iterations
	for i := 0; i < b.N; i++ {
		go WithLock()
	}
}
