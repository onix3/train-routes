package src

import (
	"encoding/json"
	"strings"
	"time"
)

// Загрузить кэш запросов
func LoadCache() {
	if j := A.Preferences().StringWithFallback("cache_requests", ""); j != "" {
		// загрузить весь кэш, отобрать только актуальный (за сегодня) и пересохранить
		existsRoutes := map[string][]route{}
		err := json.Unmarshal([]byte(j),&existsRoutes)
		for k,v := range existsRoutes {
			if strings.Contains(k,time.Now().Format("2006-01-02")) {
				routesCache[k] = v
			}
		}
		IsErr(err)
		saveCache()
	}
}

// Сохранить кэш запросов
func saveCache() {
	data,err := json.Marshal(routesCache)
	IsErr(err)
	A.Preferences().SetString("cache_requests", string(data))
}