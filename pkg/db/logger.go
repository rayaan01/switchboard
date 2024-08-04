package db

import (
	"encoding/csv"
	"fmt"
	"time"
)

func logger(index int, key string, value string, duration float64, writer *csv.Writer) {
	timestamp := time.Now().UTC().Format(time.RFC3339)
	formattedIndex := fmt.Sprintf("%d", index)
	executionTime := fmt.Sprintf("%.6f", duration)
	writer.Write([]string{formattedIndex, timestamp, key, value, executionTime})
}
