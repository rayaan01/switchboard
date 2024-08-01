package db

import (
	"encoding/csv"
	"fmt"
	"time"
)

func logger(duration float64, writer *csv.Writer) {
	timestamp := time.Now().UTC().Format(time.RFC3339)
	executionTime := fmt.Sprintf("%.6f", duration)
	writer.Write([]string{timestamp, executionTime})
}
