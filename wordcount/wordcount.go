package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type KeyValue struct {
	Key   string
	Value int
}

var MappedData []KeyValue
var ReducedData = make(map[string]int)

// This is mapper for map function
func mapper() {
	data := read()
	defer data.Close()

	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		cleanedWord := strings.Trim(word, ",.?!;:\"'")
		cleanedWord = strings.ToLower(cleanedWord)

		MappedData = append(MappedData, KeyValue{cleanedWord, 1})
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error scanning file: %v", err)
	}

	sort.Slice(MappedData, func(i, j int) bool {
		return MappedData[i].Key < MappedData[j].Key
	})

	for _, kv := range MappedData {
		fmt.Println(kv.Key, kv.Value)
	}

	reducer()
}

func reducer() {
	fmt.Println("...........................................................................")

	if len(MappedData) == 0 {
		return
	}

	currentKey := MappedData[0].Key
	count := 1

	for i := 1; i < len(MappedData); i++ {
		kv := MappedData[i]

		if kv.Key == currentKey {
			count++
		} else {
			ReducedData[currentKey] = count

			currentKey = kv.Key
			count = 1
		}
	}

	ReducedData[currentKey] = count

	for key, value := range ReducedData {
		fmt.Println(key, value)
	}
}

func read() *os.File {
	data, err := os.Open("file1.txt")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	return data
}

func main() {
	mapper()
}
