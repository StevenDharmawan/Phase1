package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

func main() {
	mutex := &sync.Mutex{}
	mapping := make(map[string]int)
	wg := &sync.WaitGroup{}
	f, err := os.Open("sample.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		wg.Add(1)
		go addToMap(mapping, scanner.Text(), wg, mutex)
	}
	wg.Wait()
	for key, value := range mapping {
		fmt.Println(key, ":", value)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func addToMap(mapping map[string]int, key string, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	mutex.Lock()
	defer mutex.Unlock()
	mapping[key]++
}
