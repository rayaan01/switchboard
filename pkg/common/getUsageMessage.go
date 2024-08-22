package common

func GetUsageMessage() []byte {
	msg := "Available commands:\n1. set [key] [value]\n2. get [key]\n3. del [key]\n4. get-range [lower_bound] [upper_bound]\n5. create-access-key {HashTable|AVLTree}\n6. exit"
	return []byte(msg)
}
