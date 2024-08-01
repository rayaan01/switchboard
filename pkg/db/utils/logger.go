package utils

import (
	"encoding/csv"
	"fmt"
	"time"
)

func Logger(duration float64, writer *csv.Writer) {
	writer.Write([]string{time.Now().Format(time.RFC3339), fmt.Sprintf("%.6f", duration)})
}
