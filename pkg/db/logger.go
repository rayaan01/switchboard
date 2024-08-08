package db

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"time"
)

func logger(index int, duration float64, key string, writer *csv.Writer) {
	timestamp := time.Now().UTC().Format(time.RFC3339)
	formattedIndex := fmt.Sprintf("%d", index)
	executionTime := fmt.Sprintf("%.6f", duration)
	var csvData []string
	if key == "" {
		csvData = []string{formattedIndex, timestamp, executionTime}
	} else {
		csvData = []string{formattedIndex, timestamp, key, executionTime}
	}
	writer.Write(csvData)
}

func logger_tps(tps int, writer *csv.Writer) {
	formattedTps := fmt.Sprintf("%d", tps)
	writer.Write([]string{formattedTps})
}

func generateRandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
