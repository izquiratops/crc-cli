package main

import (
	"encoding/json"
	"fmt"
	"hash/crc32"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: json-crc32 <json-file>")
		os.Exit(1)
	}

	filename := os.Args[1]

	// Read the JSON file
	jsonData, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Validate JSON
	var js json.RawMessage
	if err := json.Unmarshal(jsonData, &js); err != nil {
		fmt.Printf("Error: Invalid JSON in file: %v\n", err)
		os.Exit(1)
	}

	// Calculate CRC-32
	crc := crc32.ChecksumIEEE(jsonData)

	fmt.Printf("CRC-32 of %s: %08x\n", filename, crc)
}
