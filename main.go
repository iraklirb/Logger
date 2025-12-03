package main

import (
	"crypto/rand"
	"logger/logger"
	"strconv"
	"sync"
)

func main() {
	log := logger.GetInstance(logger.Debug)
	log.Debug("Starting program")

	var wg sync.WaitGroup

	for i := 1; i <= 1; i++ {
		wg.Add(1)

		go worker(i, &wg, log)
	}

	wg.Wait()

	log.Debug("all workers finished job")
}

func worker(id int, wg *sync.WaitGroup, log *logger.Logger) {
	defer wg.Done()

	for i := 1; i <= 1; i++ {
		log.Debug("worker " + strconv.Itoa(id) + ": " + randomSecureString(50))
	}

}

func randomSecureString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	bytes := make([]byte, n)
	rand.Read(bytes)

	for i := range bytes {
		bytes[i] = letters[bytes[i]%byte(len(letters))]
	}
	return string(bytes)
}
