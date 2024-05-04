package db

type StorageEngine interface {
	set(key string, value string) ([]byte, error)
	get(key string) ([]byte, error)
	del(key string) ([]byte, error)
}

// Map access keys to store instances
var StoreMapper = map[string]StorageEngine{}
