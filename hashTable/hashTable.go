package hashTable

import "fmt"

type HashTable struct {
	data [][]any
}

func (hashTable *HashTable) hash(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash = (hash + (int(key[i]) * i)) % len(key)
	}
	return hash
}

func (hashTable *HashTable) Set(key string, value int) {
	address := hashTable.hash(key)
	fmt.Printf("address is %v\n", address)
	temp := []any{}
	temp = append(temp, key)
	temp = append(temp, value)
	arr := hashTable.data[address]
	if arr == nil {

		hashTable.data[address] = []any{}
		//hashTable.data[address] = append(hashTable.data[address], temp)
	} //else {
	//fmt.Printf("address exists %v\n", address)
	//hashTable.data[address] = temp
	//}
	hashTable.data[address] = append(hashTable.data[address], temp)
}

func (hashTable HashTable) Get(key string) any {
	address := hashTable.hash(key)
	currentBucket := hashTable.data[address]

	for i := 0; i < len(currentBucket); i++ {

		internalArray, ok := currentBucket[i].([]any)
		if !ok {
			return nil
		}
		if internalArray[0] == key {
			return internalArray[1]
		}
	}
	return nil

}

func New(size int) *HashTable {
	return &HashTable{
		data: make([][]any, size),
	}
}

func (hashTable *HashTable) GetData() [][]any {
	return hashTable.data
}
