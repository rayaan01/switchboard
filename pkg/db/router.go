package db

import (
	"fmt"
	"io"
	"strings"
	"switchboard/pkg/common"
	"switchboard/pkg/prometheus"
	"time"

	"github.com/google/uuid"
)

func router(accessKey string, args []string) ([]byte, error) {
	usageMessage := common.GetUsageMessage()
	cmdType := strings.ToLower(args[0])

	switch cmdType {
	case "set":
		prometheus.SetCounter.Inc()
		if len(args) != 3 {
			return usageMessage, nil
		}
		key := args[1]
		val := args[2]
		engine, ok := StoreMapper[accessKey]
		if !ok {
			return []byte("(invalid access key)"), nil
		}
		start := time.Now()
		response, err := engine.set(key, val)
		if err != nil {
			return nil, err
		}
		duration := time.Since(start)
		prometheus.SetHistogram.WithLabelValues(key).Observe(duration.Seconds())
		fmt.Println("Time taken: ", duration.Seconds())
		return response, nil

	case "get":
		prometheus.GetCounter.Inc()
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
		prometheus.DelCounter.Inc()
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

	case "visualize-hash-table":
		engine, ok := StoreMapper[accessKey]
		if !ok {
			return []byte("(invalid access key)"), nil
		}
		engine.visualizeHashTable()
		return []byte("OK"), nil

	case "visualize-avl-tree":
		engine, ok := StoreMapper[accessKey]
		if !ok {
			return []byte("(invalid access key)"), nil
		}
		store := engine.getStore()
		engine.visualizeAVLTree(store)
		return []byte("OK"), nil

	default:
		return usageMessage, nil
	}
}
