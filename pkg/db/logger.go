package db

import (
	"encoding/csv"
	"fmt"
)

func logger(index int, duration float64, writer *csv.Writer) {
	// timestamp := time.Now().UTC().Format(time.RFC3339)
	formattedIndex := fmt.Sprintf("%d", index)
	executionTime := fmt.Sprintf("%.6f", duration)
	writer.Write([]string{formattedIndex, executionTime})
}
