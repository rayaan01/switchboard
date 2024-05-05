package db

import (
	"fmt"
	"io"
	"strings"

	"github.com/google/uuid"
)

func router(args []string) ([]byte, error) {
	errorMessage := []byte("Available commands:\n1. set [key] [value]\n2. get [key]\n3. del [key]\n4. exit")
	if len(args) == 0 {
		return errorMessage, nil
	}
	cmd := strings.ToLower(args[0])
	switch cmd {
	case "set":
		// if len(args) != 3 {
		// 	return errorMessage, nil
		// }
		// key := args[1]
		// val := args[2]
		// response, err := handleSet(key, val)
		// if err != nil {
		// 	return nil, err
		// }
		// return response, nil
		return []byte("OK"), nil

	case "get":
		// if len(args) != 2 {
		// 	return errorMessage, nil
		// }
		// key := args[1]
		// response, err := handleGet(key)
		// if err != nil {
		// 	return nil, err
		// }
		// return response, nil
		return []byte("OK"), nil

	case "del":
		// if len(args) != 2 {
		// 	return errorMessage, nil
		// }
		// key := args[1]
		// response, err := handleDel(key)
		// if err != nil {
		// 	return nil, err
		// }
		// return response, nil
		return []byte("OK"), nil

	case "create-access-key":
		if len(args) != 2 {
			return errorMessage, nil
		}
		engineType := args[1]
		if engineType != "HashTable" && engineType != "AVLTree" {
			fmt.Println("len of response", len(errorMessage))
			return errorMessage, nil
		}

		accessKey := uuid.NewString()

		if engineType == "HashTable" {
			StoreMapper[accessKey] = &HashTable{}
		} else {
			StoreMapper[accessKey] = &AVLTree{}
		}

		response := fmt.Sprintf("Your access key is: %s. Please keep it safe as it's your gateway to the database. Run the command `use %s` to set it as the default key for this session.", accessKey, accessKey)
		return []byte(response), nil

	case "exit":
		return nil, io.EOF

	default:
		return errorMessage, nil
	}
}
