package server

import (
	"strconv"
	"sync"
	"time"
)

var (
	TTLmap   = make(map[string]map[string]time.Time)
	ttlmutex sync.RWMutex
)

func TTLWatcher() {

	go func() {
		for {

			time.Sleep(time.Second * 1)
			ttlmutex.Lock()
			mut.Lock()

			for secretKey, keys := range TTLmap {
				for key, expiery := range keys {
					if time.Now().After(expiery) {
						delete(globalStore[secretKey], key)
						delete(TTLmap[secretKey], key)
					}
				}
			}
			ttlmutex.Unlock()
			mut.Unlock()

		}

	}()

}

func SetTTL(user *client, key string, expierySeconds string) bool {

	ttlmutex.Lock()
	defer ttlmutex.Unlock()

	secretKey := user.secretKey
	expSeconds, err := strconv.Atoi(expierySeconds)
	if err != nil {
		return false
	}
	exp := time.Now().Add(time.Second * time.Duration(expSeconds))

	_, exists := TTLmap[secretKey]

	if exists {
		TTLmap[secretKey][key] = exp
		return true
	} else {
		TTLmap[secretKey] = make(map[string]time.Time)
		TTLmap[secretKey][key] = exp
		return true
	}

}
