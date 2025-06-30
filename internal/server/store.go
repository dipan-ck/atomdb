package server

import "sync"

var (
	globalStore = make(map[string]map[string]string)
	mut         sync.RWMutex
)

func GetKey(user *client, key string) (string, bool) {

	mut.RLock()
	defer mut.RUnlock()

	secretKey := user.secretKey

	_, exists := globalStore[secretKey]

	if exists {
		val, ok := globalStore[secretKey][key]
		RecentlyUsed(user.LRU, key)
		return val, ok
	} else {
		return "", false
	}

}

func SetKey(user *client, key string, value string) bool {

	mut.Lock()
	defer mut.Unlock()

	secretKey := user.secretKey
	valMap, exists := globalStore[secretKey]

	if !exists {
		globalStore[secretKey] = make(map[string]string)
		valMap = globalStore[secretKey]
	}

	if _, keyExists := valMap[key]; keyExists {
		valMap[key] = value
		RecentlyUsed(user.LRU, key)
	} else {
		valMap[key] = value
		AddNode(user.LRU, key, secretKey)
	}

	return true

}
