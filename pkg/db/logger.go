package db

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"time"
)

func logger(index int, key string, value string, duration float64, writer *csv.Writer) {
	timestamp := time.Now().UTC().Format(time.RFC3339)
	formattedIndex := fmt.Sprintf("%d", index)
	executionTime := fmt.Sprintf("%.6f", duration)
	writer.Write([]string{formattedIndex, timestamp, key, value, executionTime})
}

func generateRandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
