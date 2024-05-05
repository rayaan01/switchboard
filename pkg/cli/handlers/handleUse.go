package handlers

import (
	"encoding/json"
	"os"
)

func HandleUse(accessKey string) ([]byte, error) {
	fileName := "keys.json"
	existingKeys, err := os.ReadFile(fileName)

	if err != nil {
		keyStructure := map[string]bool{accessKey: true}
		marshalledKey, err := json.Marshal(keyStructure)
		if err != nil {
			return nil, err
		}
		os.WriteFile(fileName, marshalledKey, 0644)
		return []byte("OK"), nil
	}

	var unmarshalledKeys map[string]bool
	err = json.Unmarshal(existingKeys, &unmarshalledKeys)
	if err != nil {
		return nil, err
	}
	for k := range unmarshalledKeys {
		unmarshalledKeys[k] = false
	}
	unmarshalledKeys[accessKey] = true
	marshalledKeys, err := json.Marshal(unmarshalledKeys)
	if err != nil {
		return nil, err
	}
	os.WriteFile(fileName, marshalledKeys, 0644)
	return []byte("OK"), nil
}
