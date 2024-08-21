package db

type StorageEngine interface {
	set(key string, value string) ([]byte, error)
	get(key string) ([]byte, error)
	del(key string) ([]byte, error)
	get_range(low string, high string) ([]byte, int, error)
}

var StoreMapper = map[string]StorageEngine{}
