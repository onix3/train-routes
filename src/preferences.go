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

// Загрузить из preferences города, маршруты между которыми запрашивались в последний раз
func LoadLastCities() (c1,c2 string) {
	s := A.Preferences().StringWithFallback("last_cities", "")
	split := strings.Split(s,"•")
	if s == "" || len(split) != 2 {
		return "Брест", "Минск"
	}
	return split[0], split[1]
}

// Сохранить в preferences города, маршруты между которыми запрашивались в последний раз
func saveLastCities(c1,c2 string) {
	A.Preferences().SetString("last_cities", c1 + "•" + c2)
}