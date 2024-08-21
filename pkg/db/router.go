package db

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
	"switchboard/pkg/common"
	"time"

	"github.com/google/uuid"
)

var keys = []string{}

func router(accessKey string, args []string) ([]byte, error) {
	usageMessage := common.GetUsageMessage()
	cmdType := strings.ToLower(args[0])

	switch cmdType {
	case "set":
		if len(args) != 3 {
			return usageMessage, nil
		}
		key := args[1]
		val := args[2]
		engine, ok := StoreMapper[accessKey]
		if !ok {
			return []byte("(invalid access key)"), nil
		}
		response, err := engine.set(key, val)
		if err != nil {
			return nil, err
		}
		return response, nil

	case "get":
		if len(args) != 2 {
			return usageMessage, nil
		}
		key := args[1]
		engine, ok := StoreMapper[accessKey]
		if !ok {
			return []byte("(invalid access key)"), nil
		}
		response, err := engine.get(key)
		if err != nil {
			return nil, err
		}
		return response, nil

	case "del":
		if len(args) != 2 {
			return usageMessage, nil
		}
		key := args[1]
		engine, ok := StoreMapper[accessKey]
		if !ok {
			return []byte("(invalid access key)"), nil
		}
		response, err := engine.del(key)
		if err != nil {
			return nil, err
		}
		return response, nil

	case "get-range":
		if len(args) != 3 {
			return usageMessage, nil
		}
		low := args[1]
		high := args[2]
		engine, ok := StoreMapper[accessKey]
		if !ok {
			return []byte("(invalid access key)"), nil
		}
		response, _, err := engine.get_range(low, high)

		if err != nil {
			return nil, err
		}
		return response, nil

	case "create-access-key":
		if len(args) != 2 {
			return usageMessage, nil
		}
		engineType := args[1]
		if engineType != "HashTable" && engineType != "AVLTree" {
			return usageMessage, nil
		}

		accessKey := uuid.NewString()

		if engineType == "HashTable" {
			StoreMapper[accessKey] = &HashTable{store: map[string]string{}}
		} else {
			StoreMapper[accessKey] = &AVLTree{store: nil}
		}

		response := fmt.Sprintf("Your access key is: %s. Please keep it safe as it's your gateway to the database. Run the command `use %s` to set it as the default key for this session.", accessKey, accessKey)
		return []byte(response), nil

	case "exit":
		return nil, io.EOF

	case "benchmark-set":
		filePath := "benchmark_set.csv"
		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		defer file.Close()
		if err != nil {
			fmt.Printf("Error opening %s file: %s", filePath, err)
		}
		metricsWriter := csv.NewWriter(file)
		defer metricsWriter.Flush()
		engine, ok := StoreMapper[accessKey]
		if !ok {
			return []byte("(invalid access key)"), nil
		}
		for i := 0; i < 1000; i++ {
			// For benchmarking custom random keys - Example: 1MB keys
			key := generateRandomString(36, false, false)

			// key := uuid.NewString()
			// For benchmarking get and del, keep track of keys inserted
			keys = append(keys, key)
			value := uuid.NewString()
			start := time.Now()
			_, err := engine.set(key, value)
			if err != nil {
				return nil, err
			}
			duration := time.Since(start).Seconds() * 1e9
			logger(i+1, duration, key, metricsWriter)
		}
		return []byte("Done"), nil

	case "benchmark-set-from-keys":
		engine, ok := StoreMapper[accessKey]
		if !ok {
			return []byte("(invalid access key)"), nil
		}
		for i := 0; i < len(keys); i++ {
			key := keys[i]
			value := uuid.NewString()
			_, err := engine.set(key, value)
			if err != nil {
				return nil, err
			}
		}
		return []byte("Done"), nil

	case "benchmark-get":
		engine, ok := StoreMapper[accessKey]
		if !ok {
			return []byte("(invalid access key)"), nil
		}

		file, err := os.Open("benchmark_get.csv")
		defer file.Close()
		if err != nil {
			fmt.Printf("Error opening file: %s", err)
		}
		metricsReader := csv.NewReader(file)
		records, err := metricsReader.ReadAll()
		if err != nil {
			fmt.Println("Error reading CSV file:", err)
			return nil, err
		}

		file_get, err := os.OpenFile("benchmark_get.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Printf("Error opening %s file: %s", "benchmark-get.csv", err)
		}
		defer file_get.Close()
		metricsWriter := csv.NewWriter(file_get)
		defer metricsWriter.Flush()

		for i := 0; i < len(records); i++ {
			key := records[i][2]
			start := time.Now()
			_, err := engine.get(key)
			if err != nil {
				return nil, err
			}
			duration := time.Since(start).Seconds() * 1e9
			logger(i+1, duration, "", metricsWriter)
		}

		return []byte("Done"), nil

	case "benchmark-get-range":
		filePath := "benchmark_get_range.csv"
		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		defer file.Close()
		if err != nil {
			fmt.Printf("Error opening %s file: %s", filePath, err)
		}
		metricsWriter := csv.NewWriter(file)
		metricsWriter.Write([]string{"query", "keys_returned", "duration"})
		defer metricsWriter.Flush()

		engine, ok := StoreMapper[accessKey]
		if !ok {
			return []byte("(invalid access key)"), nil
		}

		ranges := [][]string{{"a", "k"}, {"d", "v"}, {"p", "y"}, {"g", "m"}, {"w", "x"}}

		for i := 0; i < len(ranges); i++ {
			start := time.Now()
			_, keys_returned, err := engine.get_range(ranges[i][0], ranges[i][1])
			if err != nil {
				return nil, err
			}
			duration := time.Since(start).Seconds() * 1e9
			rangeQuery := strings.Join(ranges[i], "-")
			logger_range_get(rangeQuery, keys_returned, duration, metricsWriter)
		}

		return []byte("Done"), nil

	case "benchmark-del":
		engine, ok := StoreMapper[accessKey]
		if !ok {
			return []byte("(invalid access key)"), nil
		}

		file, err := os.OpenFile("benchmark_del.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Printf("Error opening %s file: %s", "benchmark-del.csv", err)
		}
		defer file.Close()
		metricsWriter := csv.NewWriter(file)
		defer metricsWriter.Flush()

		for i := 0; i < len(keys); i++ {
			start := time.Now()
			_, err := engine.del(keys[i])
			if err != nil {
				return nil, err
			}
			duration := time.Since(start).Seconds() * 1e9
			logger(i+1, duration, "", metricsWriter)
		}

		return []byte("Done"), nil

	case "benchmark-tps-set":
		file_get, err := os.OpenFile("benchmark_tps_set.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Printf("Error opening %s file: %s", "metrics_tps_set.csv", err)
		}
		defer file_get.Close()
		metricsWriter := csv.NewWriter(file_get)
		defer metricsWriter.Flush()

		engine, ok := StoreMapper[accessKey]
		if !ok {
			return []byte("(invalid access key)"), nil
		}

		operations, err := measure_tps_set(engine)
		if err != nil {
			fmt.Printf("Error measuring throughput: %s", err)
		}
		logger_tps(operations, metricsWriter)

		return []byte("Done"), nil

	case "benchmark-tps-get":
		file_get, err := os.OpenFile("benchmark_tps_get.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Printf("Error opening %s file: %s", "metrics_tps_get.csv", err)
		}
		defer file_get.Close()
		metricsWriter := csv.NewWriter(file_get)
		defer metricsWriter.Flush()

		engine, ok := StoreMapper[accessKey]
		if !ok {
			return []byte("(invalid access key)"), nil
		}

		operations, err := measure_tps_get(engine, keys)
		if err != nil {
			fmt.Printf("Error measuring throughput: %s", err)
		}

		logger_tps(operations, metricsWriter)

		return []byte("Done"), nil

	case "benchmark-tps-del":
		file_get, err := os.OpenFile("benchmark_tps_del.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Printf("Error opening %s file: %s", "metrics_tps_del.csv", err)
		}
		defer file_get.Close()
		metricsWriter := csv.NewWriter(file_get)
		defer metricsWriter.Flush()

		engine, ok := StoreMapper[accessKey]
		if !ok {
			return []byte("(invalid access key)"), nil
		}

		operations, err := measure_tps_del(engine, keys)
		if err != nil {
			fmt.Printf("Error measuring throughput: %s", err)
		}

		logger_tps(operations, metricsWriter)

		return []byte("Done"), nil

	case "log-key-value":
		filePath := "log_key_value.csv"
		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		defer file.Close()
		if err != nil {
			fmt.Printf("Error opening %s file: %s", filePath, err)
		}

		metricsWriter := csv.NewWriter(file)
		defer metricsWriter.Flush()
		metricsWriter.Write([]string{"key", "value"})

		for i := 0; i < 1000; i++ {
			key := uuid.NewString()
			value := uuid.NewString()
			log_key_value(key, value, metricsWriter)
		}

		return []byte("Done"), nil

	case "set-key-value-from-log":
		filePath := "log_key_value.csv"
		file, err := os.Open(filePath)
		defer file.Close()
		if err != nil {
			fmt.Printf("Error opening %s file: %s", filePath, err)
		}

		reader := csv.NewReader(file)
		records, err := reader.ReadAll()
		if err != nil {
			fmt.Printf("Error opening %s file: %s", filePath, err)
		}

		engine, ok := StoreMapper[accessKey]
		if !ok {
			return []byte("(invalid access key)"), nil
		}

		for i, record := range records {
			if i == 0 {
				continue
			}
			key := record[0]
			value := record[1]
			_, err := engine.set(key, value)
			if err != nil {
				return nil, err
			}
		}

		return []byte("Done"), nil

	case "verify-implementation":
		filePath := "log_key_value.csv"
		file, err := os.Open(filePath)
		defer file.Close()
		if err != nil {
			fmt.Printf("Error opening %s file: %s", filePath, err)
		}

		reader := csv.NewReader(file)
		records, err := reader.ReadAll()
		if err != nil {
			fmt.Printf("Error opening %s file: %s", filePath, err)
		}

		writeFilePath := "verify_implementation.csv"
		writeFile, err := os.OpenFile(writeFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		defer file.Close()
		if err != nil {
			fmt.Printf("Error opening %s file: %s", filePath, err)
		}

		metricsWriter := csv.NewWriter(writeFile)
		metricsWriter.Write([]string{"key", "value"})
		defer metricsWriter.Flush()

		engine, ok := StoreMapper[accessKey]
		if !ok {
			return []byte("(invalid access key)"), nil
		}

		for i, record := range records {
			if i == 0 {
				continue
			}
			key := record[0]
			value, err := engine.get(key)
			if err != nil {
				return nil, err
			}
			log_key_value(key, string(value), metricsWriter)
		}

		return []byte("Done"), nil

	default:
		return usageMessage, nil
	}
}
