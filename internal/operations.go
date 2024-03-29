package internal

func handleGet(key string) ([]byte, error) {
	value, ok := Store[key]
	if !ok {
		return []byte("(nil)"), nil
	}
	return []byte(value), nil
}

func handleSet(key string, value string) ([]byte, error) {
	Store[key] = value
	return []byte("OK"), nil
}

func handleDel(key string) ([]byte, error) {
	delete(Store, key)
	return []byte("OK"), nil
}
