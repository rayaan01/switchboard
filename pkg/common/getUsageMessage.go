package common

func GetUsageMessage() []byte {
	msg := "Available commands:\n1. set [key] [value]\n2. get [key]\n3. del [key]\n4. create-access-key {HashTable|AVLTree}\n5. exit"
	return []byte(msg)
}
