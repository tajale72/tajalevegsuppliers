package add

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// AddNumbersFromFile reads numbers from a file, trims spaces, and adds them
func AddNumbersFromFile(filePath string) (float64, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return 0, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	var total float64

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Trim spaces and parse the number
		numberStr := strings.TrimSpace(line)
		if numberStr == "" {
			continue // Skip empty lines
		}

		number, err := strconv.ParseFloat(numberStr, 64)
		if err != nil {
			log.Printf("Skipping invalid number '%s': %v", numberStr, err)
			continue
		}

		// Add to total
		total += number
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("failed to read file: %v", err)
	}

	return total, nil
}
