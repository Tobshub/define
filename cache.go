package main

import (
	"encoding/json"
	"os"
	"path"
)

var CACHE_FILE = func() string {
	dir, _ := os.UserCacheDir()
	return path.Join(dir, "define", "cache.json")
}()

func init() {
	if !fileExists(CACHE_FILE) {
		os.MkdirAll(path.Dir(CACHE_FILE), 0755)
		os.WriteFile(CACHE_FILE, []byte("{}"), 0644)
	}
}

var cache map[string]DictRes

func loadCache() {
	if cache == nil {
		cacheContent, _ := os.ReadFile(CACHE_FILE)
		json.Unmarshal(cacheContent, &cache)
	}
}

func SaveInCache(word string, dict *DictRes) {
	loadCache()
	cache[word] = *dict
	data, _ := json.Marshal(cache)
	os.WriteFile(CACHE_FILE, data, 0644)
}

func GetFromCache(word string) *DictRes {
	loadCache()
	dict, ok := cache[word]
	if !ok {
		return nil
	}
	return &dict
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
