package services

import (
	"encoding/json"
	"log"
	"os"
	"server/models"
)

var stores map[string]models.Store

func InitStoreData(filePath string) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to load store data: %v", err)
	}

	var storeList []models.Store
	if err := json.Unmarshal(data, &storeList); err != nil {
		log.Fatalf("Failed to unmarshal store data: %v", err)
	}

	stores = make(map[string]models.Store)
	for _, store := range storeList {
		stores[store.StoreID] = store
	}
}
