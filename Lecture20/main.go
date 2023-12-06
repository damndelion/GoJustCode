package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type City struct {
	Name       string
	Population int
}

type CityData struct {
	CitiesMap map[string]City
	Mutex     sync.RWMutex
}

func updateData(appData *CityData) {
	appData.Mutex.Lock()
	defer appData.Mutex.Unlock()

	updatedCitiesMap := make(map[string]City)

	for key, city := range appData.CitiesMap {
		city.Population += 100
		updatedCitiesMap[key] = city
	}

	appData.CitiesMap = updatedCitiesMap

	fmt.Println("Updated!")
}

func main() {
	cityData := &CityData{
		CitiesMap: map[string]City{
			"Astana": {Name: "Astana", Population: 1200000},
			"Almaty": {Name: "Almaty", Population: 200000},
			"London": {Name: "London", Population: 5000000},
			"Tokyo":  {Name: "Tokyo", Population: 7000000},
		},
		Mutex: sync.RWMutex{},
	}

	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		for {
			select {
			case <-ticker.C:
				updateData(cityData)
			}
		}
	}()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		cityData.Mutex.RLock()
		defer cityData.Mutex.RUnlock()

		c.JSON(http.StatusOK, cityData.CitiesMap)
	})

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
